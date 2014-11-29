package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
)

type GetGroupRes struct {
	Connection   base.Connection
	Id           string
	Name         string
	Description  string
	Type         string
	Status       string
	ServersCount int
	Limits       GroupLimits
	Groups       []GetGroupRes
	Links        []models.Link
}

type GroupLimits struct {
	Cpu       int
	MemoryGB  int
	StorageGB int
}

var _ models.LinkModel = (*GetGroupRes)(nil)

func (r *GetGroupRes) GetLinks() []models.Link {
	return r.Links
}

func (r *GetGroupRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *GetGroupRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *GetGroupRes) Self() (res *GetGroupRes, err error) {
	err = models.ResolveLink(r, "self", res)
	return
}

func (r *GetGroupRes) Billing() (res *GetGroupBilingDetailsRes, err error) {
	err = models.ResolveLink(r, "billing", res)
	return
}

func (r *GetGroupRes) Statistics() (res *GetGroupMonitoringStatisticsRes, err error) {
	err = models.ResolveLink(r, "statistics", res)
	return
}

func (r *GetGroupRes) Server() (res *servers.GetServerRes, err error) {
	err = models.ResolveLink(r, "server", res)
	return
}
