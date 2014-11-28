package centurylink_sdk

import (
	"encoding/json"
	"fmt"
	gc "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"testing"
	"unicode"

	"github.com/s-matyukevich/centurylink_sdk/models/authentication"
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
		loginRes := authentication.LoginRes{AccountAlias: "ALIAS", BearerToken: "token"}
		js, _ := json.Marshal(loginRes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	server := httptest.NewServer(s.serveMux)
	BaseUrl = server.URL + "/"

	s.client = NewClient()
	err := s.client.Connect("someUser", "somePassword")
	c.Check(err, gc.IsNil)
	c.Check(s.client.connection.accountAlias, gc.Equals, "ALIAS")
	c.Check(s.client.connection.bearerToken, gc.Equals, "token")
}

func (s *ApiSuite) setResponse(url string, resBody []byte) {
	s.serveMux.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write(resBody)
	})
}

type sorter struct {
	array []interface{}
}

func (s *sorter) Len() int {
	return len(s.array)
}

func (s *sorter) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

func (s *sorter) Less(i, j int) bool {
	return fmt.Sprintf("%v", s.array[i]) < fmt.Sprintf("%v", s.array[j])
}

func (s *ApiSuite) DeepCompareObjects(prefix string, obj1 interface{}, obj2 interface{}) error {
	switch obj1.(type) {
	case string, float64, bool:
		if obj1 != obj2 {
			return fmt.Errorf("Mistmatch in property %s. Values: %v %v", prefix, obj1, obj2)
		}
		return nil
	case []interface{}:
		array1 := obj1.([]interface{})
		array2 := obj2.([]interface{})
		if len(array1) == 0 && len(array2) == 0 {
			return nil
		}
		if len(array1) != len(array2) {
			return fmt.Errorf("Different array length for property %s - %b %b. Values %v %v", prefix, len(array1), len(array2), obj1, obj2)
		}
		sorter1 := &sorter{array: array1}
		sorter2 := &sorter{array: array2}
		sort.Sort(sorter1)
		sort.Sort(sorter2)
		for i := 0; i < len(array1); i++ {
			res := s.DeepCompareObjects(prefix+"["+strconv.Itoa(i)+"]", array1[i], array2[i])
			if res != nil {
				return res
			}
		}
	case map[string]interface{}:
		map1 := obj1.(map[string]interface{})
		map2 := obj2.(map[string]interface{})
		for key, value := range map1 {
			//all property names in modesl starts with uppercase, but in returned json they can be lowercase
			//so we make a conversion here
			r := []rune(key)
			r[0] = unicode.ToUpper(r[0])
			key2 := string(r)
			res := s.DeepCompareObjects(prefix+"."+key, value, map2[key2])
			if res != nil {
				return res
			}
		}
	}
	return nil
}

func (s *ApiSuite) Check(c *gc.C, js []byte, obj interface{}) {
	var obj1 interface{}
	err := json.Unmarshal(js, &obj1)
	c.Check(err, gc.IsNil)

	newJson, err := json.Marshal(&obj)
	c.Check(err, gc.IsNil)
	var obj2 interface{}
	err = json.Unmarshal(newJson, &obj2)
	c.Check(err, gc.IsNil)
	//check two maps dor deep equality and get user friendly error message
	err = s.DeepCompareObjects("", obj1, obj2)
	c.Check(err, gc.IsNil)
}

func (s *ApiSuite) TestGetDatacenterDeploymentCapabilities(c *gc.C) {
	resJson := []byte(`{
  "supportsPremiumStorage":true,
  "supportsSharedLoadBalancer":true,
  "deployableNetworks":[
    {
      "name":"My Network",
      "networkId":"a933432bd8894e84b6c4fb123e48cb8b",
      "type":"private",
      "accountID":"ACCT"
    }
  ],
  "templates":[
    {
      "name":"CENTOS-6-64-TEMPLATE",
      "description":"CentOS 6 | 64-bit",
      "storageSizeGB":17,
      "capabilities":[
        "cpuAutoscale"
      ],
      "reservedDrivePaths":[
        "bin",
        "boot",
        "build",
        "cdrom",
        "compat",
        "dist",
        "dev",
        "entropy",
        "etc",
        "home",
        "initrd.img",
        "lib",
        "lib64",
        "libexec",
        "lost+found",
        "media",
        "mnt",
        "opt",
        "proc",
        "root",
        "sbin",
        "selinux",
        "srv",
        "sys",
        "tmp",
        "usr",
        "var",
        "vmlinuz"
      ]
    },
    {
      "name":"WA1ACCTCUST01",
      "description":"My Custom Template",
      "storageSizeGB":16,
      "capabilities":[
        "cpuAutoscale"
      ],
      "reservedDrivePaths":[
        "bin",
        "boot",
        "build",
        "cdrom",
        "compat",
        "dist",
        "dev",
        "entropy",
        "etc",
        "home",
        "initrd.img",
        "lib",
        "lib64",
        "libexec",
        "lost+found",
        "media",
        "mnt",
        "opt",
        "proc",
        "root",
        "sbin",
        "selinux",
        "srv",
        "sys",
        "tmp",
        "usr",
        "var",
        "vmlinuz"
      ]
    },
    {
      "name":"RHEL-6-64-TEMPLATE",
      "description":"RedHat Enterprise Linux 6 | 64-bit",
      "storageSizeGB":17,
      "capabilities":[
        "cpuAutoscale"
      ],
      "reservedDrivePaths":[
        "bin",
        "boot",
        "build",
        "cdrom",
        "compat",
        "dist",
        "dev",
        "entropy",
        "etc",
        "home",
        "initrd.img",
        "lib",
        "lib64",
        "libexec",
        "lost+found",
        "media",
        "mnt",
        "opt",
        "proc",
        "root",
        "sbin",
        "selinux",
        "srv",
        "sys",
        "tmp",
        "usr",
        "var",
        "vmlinuz"
      ]
    },
    {
      "name":"UBUNTU-14-64-TEMPLATE",
      "description":"Ubuntu 14 | 64-bit",
      "storageSizeGB":17,
      "capabilities":[
        "cpuAutoscale"
      ],
      "reservedDrivePaths":[
        "bin",
        "boot",
        "build",
        "cdrom",
        "compat",
        "dist",
        "dev",
        "entropy",
        "etc",
        "home",
        "initrd.img",
        "lib",
        "lib64",
        "libexec",
        "lost+found",
        "media",
        "mnt",
        "opt",
        "proc",
        "root",
        "sbin",
        "selinux",
        "srv",
        "sys",
        "tmp",
        "usr",
        "var",
        "vmlinuz"
      ]
    },
    {
      "name":"WIN2008R2ENT-64",
      "description":"Windows 2008 R2 Enterprise | 64-bit",
      "storageSizeGB":60,
      "capabilities":[],
      "reservedDrivePaths":[
        "a",
        "b",
        "c",
        "d"
      ],
      "drivePathLength":1
    },
    {
      "name":"WIN2008R2STD-64",
      "description":"Windows 2008 R2 Standard | 64-bit",
      "storageSizeGB":60,
      "capabilities":[],
      "reservedDrivePaths":[
        "a",
        "b",
        "c",
        "d"
      ],
      "drivePathLength":1
    },
    {
      "name":"WIN2012R2DTC-64",
      "description":"Windows 2012 R2 Datacenter Edition | 64-bit",
      "storageSizeGB":60,
      "capabilities":[
        "cpuAutoscale"
      ],
      "reservedDrivePaths":[
        "a",
        "b",
        "c",
        "d"
      ],
      "drivePathLength":1
    }
  ]
}`)

	s.setResponse("/datacenters/ALIAS/UC1/deploymentCapabilities", resJson)

	res, err := s.client.GetDatacenterDeploymentCapabilities("UC1")
	c.Check(err, gc.IsNil)
	s.Check(c, resJson, res)
}

func (s *ApiSuite) TestGetDatacenterGroup(c *gc.C) {
	resJson := []byte(`{
    "id": "DC1",
    "name": "DC FRIENDLY NAME",
    "links": [
        {
           "rel": "self",
           "href": "/v2/datacenters/ALIAS/DC1"
        },
        {
           "rel": "group",
           "href": "/v2/groups/ALIAS/GROUP123",
           "id": "groupid",
           "name": "DC1 Hardware"
        },
        {
           "rel": "billing",
           "href": "/v2/groups/ALIAS/GROUP123/billing"
        },
        {
           "rel": "archiveGroupAction",
           "href": "/v2/groups/ALIAS/GROUP123/archive"
        },
        {
           "rel": "statistics",
           "href": "/v2/groups/ALIAS/GROUP123/statistics"
        },
        {
           "rel": "scheduledActivities",
           "href": "/v2/groups/ALIAS/GROUP123/scheduledActivities"
        }]
}`)

	s.setResponse("/datacenters/ALIAS/UC1", resJson)

	res, err := s.client.GetDatacenterGroup("UC1", true)
	c.Check(err, gc.IsNil)
	s.Check(c, resJson, res)
}

func (s *ApiSuite) TestGetDatacenterList(c *gc.C) {
	resJson := []byte(`[
    {
    "id": "DC1",
    "name": "DC FRIENDLY NAME",
    "links": [
        {
        "rel": "self",
        "href": "/v2/datacenters/ALIAS/DC1"
        }]
    },
    {
    "id": "DC2",
    "name": "DC2 FRIENDLY NAME",
    "links": [
        {
        "rel": "self",
        "href": "/v2/datacenters/ALIAS/DC2"
        }]
    }
]`)

	s.setResponse("/datacenters/ALIAS", resJson)

	res, err := s.client.GetDatacenterList()
	c.Check(err, gc.IsNil)
	s.Check(c, resJson, res)
}

func (s *ApiSuite) TestGetGroup(c *gc.C) {
	resJson := []byte(`{
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
	s.setResponse("/groups/ALIAS/wa1-5030", resJson)

	res, err := s.client.GetGroup("wa1-5030")
	c.Check(err, gc.IsNil)
	s.Check(c, resJson, res)
}
