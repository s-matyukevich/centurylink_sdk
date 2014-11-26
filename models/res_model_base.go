package models

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
)

func NewBaseModel(cn base.Connection) ResModelBase {
	return ResModelBase{connection: cn}
}

type ResModelBase struct {
	connection base.Connection
}

func (m *ResModelBase) resolveLink(link Link) (err error) {
	return nil
}
