package models

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	gc "gopkg.in/check.v1"
	"testing"
)

type LinkResolverSuite struct{}

var _ = gc.Suite(&LinkResolverSuite{})

func Test(t *testing.T) { gc.TestingT(t) }

type TestModelRes struct {
	Connection base.Connection
	Username   string
	Password   string
	Links      []Link
}

var _ LinkModel = (*TestModelRes)(nil)

func (r *TestModelRes) GetLinks() []Link {
	return r.Links
}

func (r *TestModelRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *TestModelRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}

func (s *LinkResolverSuite) TestGetLink(c *gc.C) {
	model := &TestModelRes{
		Links: []Link{Link{Rel: "self"}, Link{Rel: "parent"}},
	}
	res, err := getLink(model, "self")
	c.Check(err, gc.IsNil)
	c.Check(res.Rel, gc.Equals, "self")

	res, err = getLink(model, "someLink")
	c.Check(err, gc.ErrorMatches, "There is no link with rel someLink in model \\*models.TestModelRes")
}

type fakeConnection struct {
	C    *gc.C
	verb string
}

func (cn *fakeConnection) ExecuteRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	cn.C.Check(verb, gc.Equals, cn.verb)
	return
}

func (s *LinkResolverSuite) TestResolveLink(c *gc.C) {
	model := &TestModelRes{
		Links: []Link{Link{Rel: "self"}, Link{Rel: "parent"}},
	}
	err := ResolveLink(model, "self", "GET", nil)
	c.Check(err, gc.ErrorMatches, "Model connection is not initialized.")

	model.Connection = &fakeConnection{C: c, verb: "GET"}
	err = ResolveLink(model, "self", "GET", nil)
	c.Check(err, gc.IsNil)

	model.Connection = &fakeConnection{C: c, verb: "POST"}
	model.Links[0].Verbs = []string{"POST"}
	err = ResolveLink(model, "self", "POST", nil)
	c.Check(err, gc.IsNil)
}
