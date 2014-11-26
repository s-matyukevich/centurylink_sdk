package centurylink_sdk

import (
	gc "gopkg.in/check.v1"
	"log"
	"os"
)

type ClientSuite struct{}

var _ = gc.Suite(&ClientSuite{})

func (s *ClientSuite) TestNewClient(c *gc.C) {
	client := NewClient()
	c.Check(client, gc.NotNil)
	c.Check(client.logger, gc.NotNil)
}

func (s *ClientSuite) TestSetLogger(c *gc.C) {
	client := &client{}

	err := client.SetLogger(nil)
	c.Check(err, gc.ErrorMatches, "Logger must not be nil.")

	logger := log.New(os.Stdout, "", log.LstdFlags)
	err = client.SetLogger(logger)
	c.Check(err, gc.IsNil)
	c.Check(client.logger, gc.Equals, logger)

	c.Check(client.connection, gc.IsNil)
	client.connection = &connection{}
	c.Check(client.connection.logger, gc.IsNil)
	client.SetLogger(logger)
	c.Check(client.connection.logger, gc.Equals, logger)
}

func (s *ClientSuite) TestExecuteRequest(c *gc.C) {
	client := NewClient()
	err := client.executeRequest("", "", nil, nil)
	c.Check(err, gc.ErrorMatches, "The client is not initialized. You should call Connect method first.")
}
