package account

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
)

type AntiAffinityPolicyRes struct {
	Connection base.Connection
	Id         string
	Name       string
	Location   string
	Links      []models.Link
}

var _ models.LinkModel = (*AntiAffinityPolicyRes)(nil)

func (r *AntiAffinityPolicyRes) GetLinks() []models.Link {
	return r.Links
}

func (r *AntiAffinityPolicyRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *AntiAffinityPolicyRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *AntiAffinityPolicyRes) Self() (res *AntiAffinityPolicyRes, err error) {
	err = models.ResolveLink(r, "self", res)
	return
}

func (r *AntiAffinityPolicyRes) Server() (res *servers.GetServerRes, err error) {
	err = models.ResolveLink(r, "server", res)
	return
}
