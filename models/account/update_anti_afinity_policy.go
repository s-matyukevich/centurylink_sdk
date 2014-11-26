package account

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type UpdateAntiAfinityPolicyReq struct {
	Name string
}

type UpdateAntiAfinityPolicyRes struct {
	models.ResModelBase
	Id       string
	Name     string
	Location string
	Links    []models.Link
}
