package servers

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetPublicIpAddressRes struct {
	models.ResModelBase
	InternalIPAddress  string
	Ports              []PortDef
	SourceRestrictions []SourceRestrictionDef
}
