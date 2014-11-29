package centurylink_sdk

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/account"
	"github.com/s-matyukevich/centurylink_sdk/models/datacenters"
	"github.com/s-matyukevich/centurylink_sdk/models/groups"
	"github.com/s-matyukevich/centurylink_sdk/models/queue"
	"github.com/s-matyukevich/centurylink_sdk/models/servers"
	"log"
)

type client struct {
	connection *connection
	logger     *log.Logger
}

func NewClient() *client {
	return &client{
		logger: getDefaultLogger(),
	}
}

func NewClientInitialized(accountAlias string, bearerToken string) *client {
	logger := getDefaultLogger()
	return &client{
		logger:     logger,
		connection: newConnectionRaw(accountAlias, bearerToken, logger),
	}
}

func getDefaultLogger() *log.Logger {
	return log.New(ioutil.Discard, "", log.LstdFlags)
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

func (cl *client) GetDatacenterDeploymentCapabilities(datacenter string) (res *datacenters.GetDatacenterDeploymentCapabilitiesRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("datacenters/{accountAlias}/%s/deploymentCapabilities", datacenter), nil, res)
	return
}

func (cl *client) GetDatacenterGroup(datacenter string, groupLinks bool) (res *datacenters.GetDatacenterGroupRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("datacenters/{accountAlias}/%s?groupLinks=%t", datacenter, groupLinks), nil, &res)
	return
}

func (cl *client) GetDatacenterList() (res []datacenters.GetDatacenterListRes, err error) {
	err = cl.executeRequest("GET", "datacenters/{accountAlias}", nil, &res)
	return
}

func (cl *client) GetStatus(statusId string) (res *queue.GetStatusRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("operations/{acctAlias}/status/%s", statusId), nil, res)
	return
}

func (cl *client) DeleteAntiAfinityPolicy(policyId string) (err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), nil, nil)
	return
}

func (cl *client) UpdateAntiAfinityPolicy(policyId string, req *account.UpdateAntiAfinityPolicyReq) (res *account.AntiAfinityPolicyRes, err error) {
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), req, res)
	return
}

func (cl *client) CreateAntiAfinityPolicy(policyId string, req *account.CreateAntiAfinityPolicyReq) (res *account.AntiAfinityPolicyRes, err error) {
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), req, res)
	return
}

func (cl *client) GetAntiAfinityPolicy(policyId string) (res *account.AntiAfinityPolicyRes, err error) {
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), nil, res)
	return
}

func (cl *client) GetAntiAfinityPolicies() (res []*account.AntiAfinityPolicyRes, err error) {
	err = cl.executeRequest("PUT", "antiAffinityPolicies/{accountAlias}/%S", nil, res)
	return
}

func (cl *client) GetGroup(groupId string) (res *groups.GetGroupRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s", groupId), nil, res)
	return
}

func (cl *client) GetGroupBiling(groupId string) (res *groups.GetGroupBilingDetailsRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s/biling", groupId), nil, res)
	return
}

func (cl *client) GetGroupMonitoringStatistics(groupId string, start *time.Time, end *time.Time, sampleInterval *time.Duration, queryType *string) (res *groups.GetGroupMonitoringStatisticsRes, err error) {
	params := make(map[string]string)
	url := fmt.Sprintf("groups/{accountAlias}/%s/biling", groupId)
	if start != nil {
		params["start"] = start.Format("2006-01-02")
	}
	if end != nil {
		params["end"] = end.Format("2006-01-02")
	}
	if sampleInterval != nil {
		params["sampleInterval"] = fmt.Sprintf("%d:%d:%d", sampleInterval.Hours(), sampleInterval.Minutes(), sampleInterval.Seconds())
	}
	if queryType != nil {
		params["queryType"] = *queryType
	}
	if len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
	}

	err = cl.executeRequest("GET", url, nil, res)
	return
}

func (cl *client) GetServer(serverId string) (res *servers.GetServerRes, err error) {
	err = cl.executeRequest("GET", fmt.Sprintf("servers/{accountAlias}/%s", serverId), nil, res)
	return
}

func (cl *client) PauseServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/pause", req, res)
	return
}

func (cl *client) DeleteServer(serverId string) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("servers/{accountAlias}/%s", serverId), nil, res)
	return
}

func (cl *client) UpdatePublicIpAddress(serverId string, publicIp string, req *servers.UpdatePublicIpAddressReq) (res *models.Link, err error) {
	err = cl.executeRequest("PUT", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), req, res)
	return
}

func (cl *client) RemovePublicIpAddress(serverId string, publicIp string) (res *models.Link, err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), nil, res)
	return
}

func (cl *client) GetPublicIpAddress(serverId string, publicIp string) (res *servers.GetPublicIpAddressRes, err error) {
	err = cl.executeRequest("PUT", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), nil, res)
	return
}

func (cl *client) AddPublicIpAddress(serverId string, publicIp string, req *servers.AddPublicIpAddressReq) (res *models.Link, err error) {
	err = cl.executeRequest("POST", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses", serverId, publicIp), req, res)
	return
}

func (cl *client) ExecutePackage(req *servers.ExecutePackageReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/executePackage", req, res)
	return
}

func (cl *client) SetMaintenanceMode(req *servers.SetMaintenanceModeReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/setMaintenance", req, res)
	return
}

func (cl *client) StartMaintenanceMode(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/startMaintenance", req, res)
	return
}

func (cl *client) StopMaintenanceMode(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/stopMaintenance", req, res)
	return
}

func (cl *client) CreateServer(req *servers.CreateServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "servers/{accountAlias}", req, res)
	return
}

func (cl *client) CreateSnapshot(req *servers.CreateSnapshotReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/createSnapshot", req, res)
	return
}

func (cl *client) ShutDownServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/shutDown", req, res)
	return
}

func (cl *client) RebootServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/reboot", req, res)
	return
}

func (cl *client) ResetServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/reset", req, res)
	return
}

func (cl *client) PowerOnServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/powerOn", req, res)
	return
}

func (cl *client) PowerOffServer(req *servers.ServerReq) (res *servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/powerOff", req, res)
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
