package servers

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type ServerRes struct {
	models.ResModelBase
	Server       string
	IsQueued     bool
	Links        []models.Link
	ErrorMessage string
}
