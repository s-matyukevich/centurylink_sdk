package servers

type AddPublicIpAddressReq struct {
	InternalIPAddress  string
	Ports              []PortDef
	SourceRestrictions []SourceRestrictionDef
}
