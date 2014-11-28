package servers

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
	"time"
)

type GetServerRes struct {
	models.ResModelBase
	Id          string
	Name        string
	Description string
	GroupId     string
	IsTemplate  bool
	LocationId  string
	OsType      string
	Status      string
	Details     DetailsDef
	Type        string
	StorageType string
	ChangeInfo  ChangeInfoDef
	Links       []models.Link
}

type DetailsDef struct {
	ШpAddresses       []IpAddressDef
	AlertPolicies     []AlertPolicyDef
	Cpu               int
	DiskCount         int
	HostName          string
	InMaintenanceMode bool
	MemoryMB          int
	PowerState        string
	StorageGB         int
	Snapshots         []SnapshotsDef
	CustomFields      []CustomFieldDef
}

type IpAddressDef struct {
	Public   string
	Internal string
}

type AlertPolicyDef struct {
	Id    string
	Name  string
	Links []models.Link
}

type SnapshotsDef struct {
	Name  string
	Links []models.Link
}

type CustomFieldsDef struct {
	Id           string
	Name         string
	Value        string
	DisplayValue string
}

type ChangeInfoDef struct {
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedDate time.Time
	ModifiedBy   string
}
