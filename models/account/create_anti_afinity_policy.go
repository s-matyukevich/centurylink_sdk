package account

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type CreatenatiAfinityPolicyReq struct {
	Name     string
	Location string
}

type CreateAntiAfinityPolicyRes struct {
	models.ResModelBase
	Id       string
	Name     string
	Location string
	Links    []models.Link
}
