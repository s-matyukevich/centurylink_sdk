package centurylink_sdk

import (
	"github.com/s-matyukevich/centurylink_sdk/base"
	"github.com/s-matyukevich/centurylink_sdk/models"
	gc "gopkg.in/check.v1"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ConnectionSuite struct{}

var _ = gc.Suite(&ConnectionSuite{})

type TestModelReq struct {
	Username string
	Password string
}

func (s *ConnectionSuite) TestPrepareRequest(c *gc.C) {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	cn := newConnectionRaw("alias", "token", logger)
	req, err := cn.prepareRequest("GET", "someUrl/{accountAlias}/someId", nil)
	c.Check(err, gc.IsNil)
	c.Check(req.URL.Path, gc.Equals, "/someUrl/alias/someId")
	c.Check(req.Header["Content-Type"][0], gc.Equals, "application/json")
	c.Check(req.Header["Authorization"][0], gc.Equals, "Bearer token")
	c.Check(req.Body, gc.Equals, nil)

	req, err = cn.prepareRequest("GET", "someUrl", &TestModelReq{Username: "Username", Password: "Password"})
	c.Check(err, gc.IsNil)
	body := make([]byte, 100)
	l, err := req.Body.Read(body)
	body = body[0:l]
	c.Check(err, gc.IsNil)
	c.Check((string)(body), gc.Equals, `{"Username":"Username","Password":"Password"}`)
}

type TestModelRes struct {
	Connection base.Connection
	Username   string
	Password   string
	Links      []models.Link
}

var _ models.LinkModel = (*TestModelRes)(nil)

func (r *TestModelRes) GetLinks() []models.Link {
	return r.Links
}

func (r *TestModelRes) GetConnection() base.Connection {
	return r.Connection
}

func (r *TestModelRes) SetConnection(connection base.Connection) {
	r.Connection = connection
}
func (s *ConnectionSuite) TestProcessResponse(c *gc.C) {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	cn := newConnectionRaw("", "", logger)
	res := &http.Response{
		StatusCode: 500,
	}
	err := cn.processResponse(res, nil)
	c.Check(err, gc.ErrorMatches, "Error occured while sending request to API. Status code: 500")

	res = &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"Username":"Username","Password":"Password"}`)),
	}
	modelRes := &TestModelRes{}
	err = cn.processResponse(res, modelRes)
	c.Check(err, gc.IsNil)
	c.Check(modelRes.Username, gc.Equals, "Username")
	c.Check(modelRes.Password, gc.Equals, "Password")
	c.Check(modelRes.Connection, gc.Equals, cn)

	res = &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`[{"Username":"Username0","Password":"Password0"},{"Username":"Username1","Password":"Password1"}]`)),
	}
	var modelResArray []*TestModelRes
	err = cn.processResponse(res, &modelResArray)
	c.Check(err, gc.IsNil)
	c.Check(modelResArray[0].Username, gc.Equals, "Username0")
	c.Check(modelResArray[0].Password, gc.Equals, "Password0")
	c.Check(modelResArray[0].Connection, gc.Equals, cn)
	c.Check(modelResArray[1].Username, gc.Equals, "Username1")
	c.Check(modelResArray[1].Password, gc.Equals, "Password1")
	c.Check(modelResArray[1].Connection, gc.Equals, cn)
}
