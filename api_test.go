package centurylink_sdk

import (
	"encoding/json"
	gc "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/s-matyukevich/centurylink_sdk/models"
)

func Test(t *testing.T) { gc.TestingT(t) }

type ApiSuite struct {
	client   *client
	serveMux *http.ServeMux
}

var _ = gc.Suite(&ApiSuite{})

func (s *ApiSuite) SetUpSuite(c *gc.C) {
	s.serveMux = http.NewServeMux()
	s.serveMux.HandleFunc("/authentication/login", func(w http.ResponseWriter, req *http.Request) {
		loginRes := models.LoginRes{AccountAlias: "test", BearerToken: "token"}
		js, _ := json.Marshal(loginRes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	server := httptest.NewServer(s.serveMux)
	BaseUrl = server.URL + "/"

	s.client = NewClient()
	err := s.client.Connect("someUser", "somePassword")
	c.Check(err, gc.IsNil)
	c.Check(s.client.connection.accountAlias, gc.Equals, "test")
	c.Check(s.client.connection.bearerToken, gc.Equals, "token")
}

func (s *ApiSuite) setResponse(url string, resBody string) {
	s.serveMux.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(resBody))
	})
}

func (s *ApiSuite) TestGetGroup(c *gc.C) {
	s.setResponse("/groups/test/123", `
{
  "id": "wa1-0001",
  "name": "Web Applications",
  "description": "public facing web apps",
  "type": "default",
  "status": "active",
  "serversCount": 2,
  "limits": {
    "cpu": 80,
    "memoryGB": 160,
    "storageGB": 4096
  },
  "groups": [
    {
      "id": "wa1-0002",
      "name": "Training Environment",
      "description": "Temporary servers",
      "type": "default",
      "status": "active",
      "serversCount": 0,
      "limits": {
        "cpu": 80,
        "memoryGB": 160,
        "storageGB": 4096
      },
      "groups": [],
      "links": [
        {
          "rel": "self",
          "href": "/v2/groups/acct/wa1-0002"
        },
        {
          "rel": "delete",
          "href": "/v2/groups/acct/wa1-0002"
        },
        {
          "rel": "billing",
          "href": "/v2/groups/acct/wa1-0002/billing"
        },
        {
          "rel": "archiveGroupAction",
          "href": "/v2/groups/acct/wa1-0002/archive"
        },
        {
          "rel": "statistics",
          "href": "/v2/groups/acct/wa1-0002/statistics"
        },
        {
          "rel": "scheduledActivities",
          "href": "/v2/groups/acct/wa1-0002/scheduledActivities"
        }
      ]
    }
  ],
  "links": [
    {
      "rel": "self",
      "href": "/v2/groups/acct/wa1-0001"
    },
    {
      "rel": "delete",
      "href": "/v2/groups/acct/wa1-0001"
    },
    {
      "rel": "parentGroup",
      "href": "/v2/groups/acct/wa1-0000",
      "id": "wa1-3728"
    },
    {
      "rel": "billing",
      "href": "/v2/groups/acct/wa1-0001/billing"
    },
    {
      "rel": "archiveGroupAction",
      "href": "/v2/groups/acct/wa1-0001/archive"
    },
    {
      "rel": "statistics",
      "href": "/v2/groups/acct/wa1-0001/statistics"
    },
    {
      "rel": "scheduledActivities",
      "href": "/v2/groups/acct/wa1-0001/scheduledActivities"
    },
    {
      "rel": "server",
      "href": "/v2/servers/acct/wa1acctpre7101",
      "id": "WA1ACCTPRE7101"
    },
    {
      "rel": "server",
      "href": "/v2/servers/btdi/wa1acctpre7202",
      "id": "WA1ACCTPRE7202"
    }
  ]
}`)

	res, err := s.client.GetGroup("123")
	c.Check(err, gc.IsNil)
	c.Check(res.Id, gc.Equals, "wa1-0001")
	c.Check(res.Name, gc.Equals, "Web Applications")
	c.Check(res.Description, gc.Equals, "public facing web apps")

}
