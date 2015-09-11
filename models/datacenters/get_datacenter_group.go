package datacenters

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/groups"
)

type GetDatacenterRes struct {
	// Injected field
	Connection	base.Connection

	// Short value representing the data center code
	Id		string		`json: "id"`

	// Full, friendly name of the data center
	Name		string		`json: "name"`

	// Collection of entity links that point to resources related to this data center
	Links		[]models.Link	`json: "links"`
}

func (r *GetDatacenterRes) GetLinks() []models.Link {
	return r.Links
}

func (r *GetDatacenterRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *GetDatacenterRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *GetDatacenterRes) RootGroup() (res *groups.GetGroupRes, err error) {
	res = &groups.GetGroupRes{}
	err = models.ResolveLink(r, "group", "GET", res)
	return
}
