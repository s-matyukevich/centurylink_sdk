package account

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type AntiAfinityPolicyRes struct {
	models.ResModelBase
	Id       string
	Name     string
	Location string
	Links    []models.Link
}
