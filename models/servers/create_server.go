package servers

import (
	"time"
)

type CreateServerReq struct {
	Nane                 string
	Description          string
	GroupId              string
	SourceServerId       string
	IsManagedOS          bool
	PrimaryDns           string
	SecondaryDns         string
	IpAddress            string
	Password             string
	SourceServerPassword string
	Cpu                  int
	CpuAutoscalePolicyId string
	MemoryGB             int
	Type                 string
	StorageType          string
	AntiAffinityPolicyId string
	CustomFields         []CustomFieldDef
	AdditionalDisks      []AdditionalDiskDef
	Ttl                  time.Time
	Packages             []PackageDef
}

type CustomFieldDef struct {
	Id    string
	Value string
}

type AdditionalDiskDef struct {
	Path   string
	SizeGB int
	Type   string
}

type PackageDef struct {
	PackageId  string
	Parameters map[string]string
}
