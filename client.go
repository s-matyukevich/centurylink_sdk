package centurylink_sdk

import (
	"fmt"

	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/groups"
	"log"
	"os"
)

type client struct {
	connection *connection
	logger     *log.Logger
}

func NewClient() *client {
	return &client{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (cl *client) Connect(username string, password string) (err error) {
	cl.connection, err = newConnection(username, password, cl.logger)
	return
}

func (cl *client) SetLogger(logger *log.Logger) (err error) {
	if logger == nil {
		err = fmt.Errorf("Logger must not be nil.")
		return
	}
	cl.logger = logger
	if cl.connection != nil {
		cl.connection.logger = logger
	}
	return
}

func (cl *client) DeleteAntiAfinityPolicy(policyId string) (res *models.GetGroupRes, err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), nil, nil)
	return
}

func (cl *client) GetGroup(groupId string) (res *groups.GetGroupRes, err error) {
	res = &groups.GetGroupRes{ResModelBase: models.NewBaseModel(cl.connection)}
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s", groupId), nil, res)
	return
}

func (cl *client) GetGroupBiling(groupId string) (res *groups.GroupBilingDetailsRes, err error) {
	res = &groups.GroupBilingDetailsRes{ResModelBase: models.NewBaseModel(cl.connection)}
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s/biling", groupId), nil, res)
	return
}

func (cl *client) executeRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	cl.logger.Printf("Sending request to API endpoint: %q, parameters: %#v", url, reqModel)
	if cl.connection == nil {
		err = fmt.Errorf("The client is not initialized. You should call Connect method first.")
		return
	}
	return cl.connection.ExecuteRequest(verb, url, reqModel, resModel)
}
