// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudhsmiface provides an interface for the Amazon CloudHSM.
package cloudhsmiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
)

// CloudHSMAPI is the interface type for cloudhsm.CloudHSM.
type CloudHSMAPI interface {
	AddTagsToResourceRequest(*cloudhsm.AddTagsToResourceInput) (*request.Request, *cloudhsm.AddTagsToResourceOutput)

	AddTagsToResource(*cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)

	CreateHapgRequest(*cloudhsm.CreateHapgInput) (*request.Request, *cloudhsm.CreateHapgOutput)

	CreateHapg(*cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)

	CreateHsmRequest(*cloudhsm.CreateHsmInput) (*request.Request, *cloudhsm.CreateHsmOutput)

	CreateHsm(*cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)

	CreateLunaClientRequest(*cloudhsm.CreateLunaClientInput) (*request.Request, *cloudhsm.CreateLunaClientOutput)

	CreateLunaClient(*cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)

	DeleteHapgRequest(*cloudhsm.DeleteHapgInput) (*request.Request, *cloudhsm.DeleteHapgOutput)

	DeleteHapg(*cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)

	DeleteHsmRequest(*cloudhsm.DeleteHsmInput) (*request.Request, *cloudhsm.DeleteHsmOutput)

	DeleteHsm(*cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)

	DeleteLunaClientRequest(*cloudhsm.DeleteLunaClientInput) (*request.Request, *cloudhsm.DeleteLunaClientOutput)

	DeleteLunaClient(*cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)

	DescribeHapgRequest(*cloudhsm.DescribeHapgInput) (*request.Request, *cloudhsm.DescribeHapgOutput)

	DescribeHapg(*cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)

	DescribeHsmRequest(*cloudhsm.DescribeHsmInput) (*request.Request, *cloudhsm.DescribeHsmOutput)

	DescribeHsm(*cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)

	DescribeLunaClientRequest(*cloudhsm.DescribeLunaClientInput) (*request.Request, *cloudhsm.DescribeLunaClientOutput)

	DescribeLunaClient(*cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)

	GetConfigRequest(*cloudhsm.GetConfigInput) (*request.Request, *cloudhsm.GetConfigOutput)

	GetConfig(*cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)

	ListAvailableZonesRequest(*cloudhsm.ListAvailableZonesInput) (*request.Request, *cloudhsm.ListAvailableZonesOutput)

	ListAvailableZones(*cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)

	ListHapgsRequest(*cloudhsm.ListHapgsInput) (*request.Request, *cloudhsm.ListHapgsOutput)

	ListHapgs(*cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)

	ListHsmsRequest(*cloudhsm.ListHsmsInput) (*request.Request, *cloudhsm.ListHsmsOutput)

	ListHsms(*cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)

	ListLunaClientsRequest(*cloudhsm.ListLunaClientsInput) (*request.Request, *cloudhsm.ListLunaClientsOutput)

	ListLunaClients(*cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)

	ListTagsForResourceRequest(*cloudhsm.ListTagsForResourceInput) (*request.Request, *cloudhsm.ListTagsForResourceOutput)

	ListTagsForResource(*cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)

	ModifyHapgRequest(*cloudhsm.ModifyHapgInput) (*request.Request, *cloudhsm.ModifyHapgOutput)

	ModifyHapg(*cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)

	ModifyHsmRequest(*cloudhsm.ModifyHsmInput) (*request.Request, *cloudhsm.ModifyHsmOutput)

	ModifyHsm(*cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)

	ModifyLunaClientRequest(*cloudhsm.ModifyLunaClientInput) (*request.Request, *cloudhsm.ModifyLunaClientOutput)

	ModifyLunaClient(*cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)

	RemoveTagsFromResourceRequest(*cloudhsm.RemoveTagsFromResourceInput) (*request.Request, *cloudhsm.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
}

var _ CloudHSMAPI = (*cloudhsm.CloudHSM)(nil)
