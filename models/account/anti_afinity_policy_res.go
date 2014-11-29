package account

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
)

type AntiAfinityPolicyRes struct {
	Connection base.Connection
	Id         string
	Name       string
	Location   string
	Links      []models.Link
}

var _ models.LinkModel = (*AntiAfinityPolicyRes)(nil)

func (r *AntiAfinityPolicyRes) GetLinks() []models.Link {
	return r.Links
}

func (r *AntiAfinityPolicyRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *AntiAfinityPolicyRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *AntiAfinityPolicyRes) Self() (res *AntiAfinityPolicyRes, err error) {
	err = models.ResolveLink(r, "self", res)
	return
}

func (r *AntiAfinityPolicyRes) Server() (res *servers.GetServerRes, err error) {
	err = models.ResolveLink(r, "server", res)
	return
}
