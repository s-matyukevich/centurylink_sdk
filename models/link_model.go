package models

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
)

type LinkModel interface {
	GetConnection() base.Connection
	SetConnection(base.Connection)
	GetLinks() []Link
}
