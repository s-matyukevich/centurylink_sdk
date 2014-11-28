package queue

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetStatusRes struct {
	models.ResModelBase
	Status string
}
