package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetGroupRes struct {
	models.ResModelBase
	Id          string
	Name        string
	Description string
	Type        string
	Status      string
	ServerCount int
	Limits      GroupLimits
	Groups      []GetGroupRes
	Links       []models.Link
}

type GroupLimits struct {
	Cpu       int
	MemoryGB  int
	StorageGB int
}
