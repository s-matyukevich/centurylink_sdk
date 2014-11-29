package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
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
	err = models.ResolveLink(r, "self", res)
	return
}

func (r *GetDatacenterGroupRes) RootGroup() (res *GetDatacenterGroupRes, err error) {
	err = models.ResolveLink(r, "group", res)
	return
}
