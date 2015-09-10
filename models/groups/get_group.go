package groups

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
)

type ChangeInfo struct {
	CreatedDate	string
	CreatedBy	string
	ModifiedDate	string
	ModifiedBy	string
}

type CustomFields struct {
	Id		string
	Name		string
	Value		string
	DisplayValue	string
}

type GetGroupRes struct {
	Connection   base.Connection
	Id           string
	Name         string
	Description  string
	LocationId   string
	Type         string
	Status       string
	ServersCount int
	Groups       []GetGroupRes
	Links        []models.Link
	ChangeInfo   ChangeInfo
	CustomFields []CustomFields
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
	err = models.ResolveLink(r, "self", "GET", res)
	return
}

func (r *GetGroupRes) Billing() (res *GetGroupBillingDetailsRes, err error) {
	err = models.ResolveLink(r, "billing", "GET", res)
	return
}

func (r *GetGroupRes) Statistics() (res *GetGroupMonitoringStatisticsRes, err error) {
	err = models.ResolveLink(r, "statistics", "GET", res)
	return
}

func (r *GetGroupRes) Server() (res *servers.GetServerRes, err error) {
	err = models.ResolveLink(r, "server", "GET", res)
	return
}
