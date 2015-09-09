package servers

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/queue"
)

type ServerRes struct {
	Connection   base.Connection
	Server       string
	IsQueued     bool
	Links        []models.Link
	ErrorMessage string
}

var _ models.LinkModel = (*ServerRes)(nil)

func (r *ServerRes) GetLinks() []models.Link {
	return r.Links
}

func (r *ServerRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *ServerRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (r *ServerRes) Status() (res *queue.GetStatusRes, err error) {
	err = models.ResolveLink(r, "status", "GET", res)
	return
}
