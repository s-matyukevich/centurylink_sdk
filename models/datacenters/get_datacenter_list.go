package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type Datacenter struct {
	models.ResModelBase
	Id    string
	Name  string
	Links []models.Link
}
