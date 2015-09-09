package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/groups"
)

type GetDatacenterGroupRes struct {
	Connection base.Connection
	Id         string
	Name       string
	Links      []models.Link
}

var _ models.LinkModel = (*GetDatacenterGroupRes)(nil)

func (r *GetDatacenterGroupRes) GetLinks() []models.Link {
	return r.Links
}

func (r *GetDatacenterGroupRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *GetDatacenterGroupRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *GetDatacenterGroupRes) Self() (res *GetDatacenterGroupRes, err error) {
	res = &GetDatacenterGroupRes{}
	err = models.ResolveLink(r, "self", "GET", res)
	return
}

func (r *GetDatacenterGroupRes) RootGroup() (res *groups.GetGroupRes, err error) {
	res = &groups.GetGroupRes{}
	err = models.ResolveLink(r, "group", "GET", res)
	return
}
