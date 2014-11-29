package datacenters

type GetDatacenterDeploymentCapabilitiesRes struct {
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
