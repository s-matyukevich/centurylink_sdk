package datacenters

type GetDatacenterDeploymentCapabilitiesRes struct {
	// Whether or not this data center provides support for servers with premium storage
	SupportsPremiumStorage		bool			`json: "supportsPremiumStorage"`

	// Whether or not this data center provides support for shared load balancer configuration
	SupportsSharedLoadBalancer	bool			`json: "supportsSharedLoadBalancer"`

	// Whether or not this data center provides support for  Whether or not this data center
	// provides support for provisioning bare metal servers
	SupportsBareMetalServers	bool			`json: "supportsBareMetalServers"`

	// Collection of networks that can be used for deploying servers
	DeployableNetworks		[]DeployableNetworks	`json: "deployableNetworks"`

	// Collection of available templates in the data center that can be used to create servers
	Templates			[]Templates		`json: "templates"`

	// Collection of available OS types that can be imported as virtual machines
	ImportableOSTypes		[]ImportableOsType	`json: "importableOSTypes"`

	// FIXME: the following appear in the output but are undocumented as of Sep 11/2015
	DataCenterEnabled		bool			`json: "dataCenterEnabled"`
	ImportVMEnabled			bool			`json: "importVMEnabled"`
}

type DeployableNetworks struct {
	// User-defined name of the network
	Name		string	`json: "name"`

	// Unique identifier of the network
	NetworkId	string	`json: "networkId"`

	// Network type, usually "private" for networks created by the user
	Type		string	`json: "type"`

	// Account alias for the account in which the network exists
	AccountID	string	`json: "accoundID"`
}

type Templates struct {
	// Underlying unique name for the template
	Name			string		`json: "name"`

	// FIXME: undocumented as of Sep 11/2015
	OsType			string		`json: "osType"`

	// Description of the template at it appears in the Control Portal UI
	Description		string		`json: "description"`

	// The amount of storage allocated for the primary OS root drive
	StorageSizeGB		int		`json: "storageSizeGB"`

	// List of capabilities supported by this specific OS template
	// (example: whether adding CPU or memory requires a reboot or not)
	Capabilities		[]string	`json: "capabilities"`

	// List of drive path names reserved by the OS that can't be used
	// to name user-defined drives. For example:
	// Linux:   bin, boot, build, cdrom, etc, home, initrd.img, lib, lib64, libexec, lost+found ...
	// Windows: a, b, c, d,
	ReservedDrivePaths	[]string	`json: "reservedDrivePaths"`

	// Length of the string for naming a drive path, if applicable. For examle:
	// Linux:   0
	// Windows: 1
	DrivePathLength		int		`json: "drivePathLength"`
}

// FIXME: there is no online description for this yet as of Sep 11/2015
type ImportableOsType struct {
	Type			string	`json: "type"`
	Description		string	`json: "description"`
	Id			int	`json: "id"`
	LabProductCode		string	`json: "labProductCode"`
	PremiumProductCode	string	`json: "premiumProductCode"`
}
