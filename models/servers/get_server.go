package servers

import (
	"time"

	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetServerRes struct {
	// Injected field; to traverse links
	Connection  base.Connection

	// ID of the server
	Id		string		`json: "id"`

	// Name of the server
	Name        	string		`json: "name"`

	// User-defined description of this server
	Description	string		`json: "description"`

	// ID of the parent group
	GroupId     	string		`json: "groupId"`

	// Boolean indicating whether this is a custom template or running server
	IsTemplate  	bool		`json: "isTemplate"`

	// Data center that this server resides in
	LocationId  	string		`json: "locationId"`

	// Friendly name of the Operating System the server is running
	OsType      	string		`json: "osType"`

	// Describes whether the server is active or not
	Status      	string		`json: "status"`

	// Resource allocations, alert policies, snapshots, and more.
	Details     	DetailsDef	`json: "details"`

	// Whether a standard or premium server
	Type        	string		`json: "type"`

	// Whether it uses standard or premium storage
	StorageType	string		`json: "storageType"`

	// Describes "created" and "modified" details
	ChangeInfo	ChangeInfoDef	`json: "changeInfo"`

	// Collection of entity links that point to resources related to this server
	Links		[]models.Link	`json: "links"`
}

type DetailsDef struct {
	// Details about IP addresses associated with the server
	IpAddresses       	[]IpAddressDef		`json: "ipAddresses"`

	// Describe each alert policy applied to the server
	AlertPolicies     	[]AlertPolicyDef	`json: "alertPolicies"`

	// How many vCPUs are allocated to the server
	Cpu               	int			`json: "cpu"`

	// How many disks are attached to the server
	DiskCount         	int			`json: "diskCount"`

	// Fully qualified name of the server
	HostName          	string			`json: "hostName"`

	// Indicator of whether server has been placed in maintenance mode
	InMaintenanceMode	bool			`json: "inMaintenanceMode"`

	// How many MB of memory are allocated to the server
	MemoryMB          	int			`json: "memoryMB"`

	// Whether the server is running or not
	PowerState        	string			`json: "powerState"`

	// How many total GB of storage are allocated to the server
	StorageGB         	int			`json: "storageGB"`

	// The disks attached to the server
	Disks			[]DiskDef		`json: "disks"`

	// The partitions defined for the server
	Partitions		[]PartitionDef		`json: "partitions"`

	// Details about any snapshot associated with the server
	Snapshots         	[]SnapshotDef		`json: "snapshots"`

	// Details about any custom fields and their values
	CustomFields      	[]CustomFieldDef	`json: "customFields"`

	// Processor configuration description (for bare metal servers only)
	ProcessorDescription   	string			`json: "processorDescription"`
	// Storage configuration description (for bare metal servers only)
	StorageDescription   	string			`json: "storageDescription"`
}

type IpAddressDef struct {
	// If applicable, the public IP
	Public		string	`json: "public"`
	// Private IP address. If associated with a public IP address, then the "public" value is populated
	Internal	string	`json: "internal"`
}

type AlertPolicyDef struct {
	// Unique identifier of the policy
	Id	string		`json: "id"`
	// User-defined name of the alert policy
	Name	string		`json: "name"`
	// Collection of entity links that point to resources related to this policy
	Links	[]models.Link	`json: "links"`
}

type DiskDef struct {
	// Unique identifier of the disk
	Id		string		`json: "id"`
	// Size of the disk in GB
	SizeGB		string		`json: "sizeGB"`
	// List of partition paths on the disk
	PartitionPaths	[]string	`json: "partitionPaths"`
}

type PartitionDef struct {
	// Size of the partition in GB
	SizeGB	string	`json: "sizeGB"`
	// List of partition paths on the disk
	Path	string	`json: "path"`
}

type SnapshotDef struct {
	// Timestamp of the snapshot
	Name  string		`json: "name"`
	// Collection of entity links that point to resources related to this snapshot
	Links []models.Link	`json: "links"`
}

// Note: this struct is also used elsewhere
type CustomFieldDef struct {
	// Unique ID of the custom field
	Id           string	`json: "id"`
	// Friendly name of the custom field
	Name         string	`json: "name"`
	// Underlying value of the field
	Value        string	`json: "value"`
	// Shown value of the field
	DisplayValue string	`json: "displayValue"`
}

// Note: this struct is also used elsewhere
type ChangeInfoDef struct {
	// Date/time that the server was created
	CreatedDate  time.Time	`json: "createdDate"`
	// Who created the server
	CreatedBy    string	`json: "createdBy"`
	// Date/time that the server was last updated
	ModifiedDate time.Time	`json: "modifiedDate"`
	// Who modified the server last
	ModifiedBy   string	`json: "modifiedBy"`
}

func (r *GetServerRes) GetLinks() []models.Link {
	return r.Links
}

func (r *GetServerRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *GetServerRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *GetServerRes) Self() (res *GetServerRes, err error) {
	err = models.ResolveLink(r, "self", "GET", res)
	return
}
