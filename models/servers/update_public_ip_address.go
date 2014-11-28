package servers

type UpdatePublicIpAddressReq struct {
	Ports              []PortDef
	SourceRestrictions []SourceRestrictionDef
}

type PortDef struct {
	Protocol string
	Port     int
	PortTo   int
}

type SourceRestrictionDef struct {
	Cidr string
}
