package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
)

type GetDatacenterListRes struct {
	Connection base.Connection
	Id         string
	Name       string
	Links      []models.Link
}

var _ models.LinkModel = (*GetDatacenterListRes)(nil)

func (r *GetDatacenterListRes) GetLinks() []models.Link {
	return r.Links
}

func (r *GetDatacenterListRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *GetDatacenterListRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *GetDatacenterListRes) Self() (res *GetDatacenterListRes, err error) {
	err = models.ResolveLink(r, "self", res)
	return
}
