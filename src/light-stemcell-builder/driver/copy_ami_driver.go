package driver

import (
	"errors"
	"fmt"
	"io"
	"light-stemcell-builder/config"
	"light-stemcell-builder/resources"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/private/waiter"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// SDKCopyAmiDriver uses the AWS SDK to register an AMI from an existing snapshot in EC2
type SDKCopyAmiDriver struct {
	creds  config.Credentials
	logger *log.Logger
}

// NewCopyAmiDriver creates a SDKCopyAmiDriver for copying AMIs in EC2
func NewCopyAmiDriver(logDest io.Writer, creds config.Credentials) *SDKCopyAmiDriver {
	logger := log.New(logDest, "SDKCopyAmiDriver ", log.LstdFlags)
	return &SDKCopyAmiDriver{creds: creds, logger: logger}
}

// Create creates an AMI, copied from a source AMI, and optionally makes the AMI publically available
func (d *SDKCopyAmiDriver) Create(driverConfig resources.AmiDriverConfig) (resources.Ami, error) {
	srcRegion := d.creds.Region
	dstRegion := driverConfig.DestinationRegion

	awsConfig := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(d.creds.AccessKey, d.creds.SecretKey, "")).
		WithRegion(dstRegion).
		WithLogger(newDriverLogger(d.logger))

	ec2Client := ec2.New(session.New(), awsConfig)

	createStartTime := time.Now()
	defer func(startTime time.Time) {
		d.logger.Printf("completed Create() in %f minutes\n", time.Since(startTime).Minutes())
	}(createStartTime)

	d.logger.Printf("copying AMI from source AMI: %s\n", driverConfig.ExistingAmiID)
	input := &ec2.CopyImageInput{
		Description:   &driverConfig.Description,
		Name:          &driverConfig.Name,
		SourceImageId: &driverConfig.ExistingAmiID,
		SourceRegion:  &srcRegion,
		Encrypted:     &driverConfig.Encrypted,
	}
	if driverConfig.KmsKeyId != "" {
		input.KmsKeyId = &driverConfig.KmsKeyId
	}
	output, err := ec2Client.CopyImage(input)
	if err != nil {
		return resources.Ami{}, fmt.Errorf("copying AMI: %s", err)
	}

	amiIDptr := output.ImageId
	if amiIDptr == nil {
		return resources.Ami{}, errors.New("AMI id nil")
	}

	d.logger.Printf("waiting for AMI: %s to be available\n", *amiIDptr)
	err = d.waitUntilImageAvailable(&ec2.DescribeImagesInput{
		ImageIds: []*string{amiIDptr},
	}, ec2Client)
	if err != nil {
		return resources.Ami{}, fmt.Errorf("waiting for AMI %s to be available: %s", *amiIDptr, err)
	}

	if driverConfig.Accessibility == resources.PublicAmiAccessibility {
		d.logger.Printf("making AMI: %s public", *amiIDptr)
		ec2Client.ModifyImageAttribute(&ec2.ModifyImageAttributeInput{
			ImageId: amiIDptr,
			LaunchPermission: &ec2.LaunchPermissionModifications{
				Add: []*ec2.LaunchPermission{
					&ec2.LaunchPermission{
						Group: aws.String(publicGroup),
					},
				},
			},
		})
	}

	return resources.Ami{ID: *amiIDptr, Region: dstRegion}, nil
}

func (d *SDKCopyAmiDriver) waitUntilImageAvailable(input *ec2.DescribeImagesInput, c *ec2.EC2) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeImages",
		Delay:       15,
		MaxAttempts: 240,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Images[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Images[].State",
				Expected: "failed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
