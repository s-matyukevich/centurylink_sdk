package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetDatacenterDeploymentCapabilitiesRes struct {
	models.ResModelBase
	SupportsPremiumStorage     bool
	SupportsSharedLoadBalancer bool
	DeployableNetworks         []DeployableNetworks
	Templates                  []Templates
}

type DeployableNetworks struct {
	Name      string
	NetworkId string
	Type      string
	AccountID string
}

type Templates struct {
	Name               string
	Description        string
	StorageSizeGB      int
	Capabilities       []string
	ReservedDrivePaths []string
	DrivePathLength    int
}
