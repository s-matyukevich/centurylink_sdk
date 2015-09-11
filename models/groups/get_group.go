package groups

import (
	"time"

	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
)

/* Group Entity Definition */
type GetGroupRes struct {
	// Injected field; to traverse links
	Connection	base.Connection

	// User-defined name of the group
	Name		string		`json: "name"`

	// User-defined description of this group
	Description	string		`json: "description"`

	// ID of the group being queried
	Id		string		`json: "id"`

	// Data center location identifier
	LocationId	string		`json: "locationId"`
	// Group type which could include system types like "archive"
	Type		string		`json: "type"`

	// Describes if group is online or not (e.g. "active")
	Status		string		`json: "status"`

	// Number of servers this group contains
	ServersCount	int		`json: "serversCount"`

	// Refers to this same entity type for each sub-group
	Groups		[]GetGroupRes	`json: "groups"`
	// Collection of entity links that point to resources related to this group
	Links		[]models.Link	`json: "links"`
	// Describes "created" and "modified" details
	ChangeInfo	ChangeInfo	`json: "changeInfo"`
	// Details about any custom fields and their values
	CustomFields	[]CustomFields	`json: "customFields"`
}

type ChangeInfo struct {
	// Date/time that the group was created (format: "2013-11-22T23:38:50Z", which
	// is compatible with RFC3339 format used by what the time package marshals).
	CreatedDate	time.Time	`json: "createdDate"`
	// Who created the group
	CreatedBy	string		`json: "createdBy"`
	// Date/time that the group was last updated (same format as @CreatedDate)
	ModifiedDate	time.Time	`json: "modifiedDate"`
	// Who modified the group last
	ModifiedBy	string		`json: "modifiedBy"`
}

type CustomFields struct {
	// Unique ID of the custom field
	Id		string	`json: "id"`
	// Friendly name of the custom field
	Name		string	`json: "name"`
	// Underlying value of the field
	Value		string	`json: "value"`
	// Shown value of the field
	DisplayValue	string	`json: "displayValue"`
}

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
