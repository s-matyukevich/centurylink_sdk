package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type DatacenterListRes struct {
	models.ResModelBase
	Id    string
	Name  string
	Links []models.Link
}
