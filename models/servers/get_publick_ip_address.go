package servers

type GetPublicIpAddressRes struct {
	InternalIPAddress  string
	Ports              []PortDef
	SourceRestrictions []SourceRestrictionDef
}
