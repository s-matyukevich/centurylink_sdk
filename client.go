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

type Client struct {
	connection *connection
	logger     *log.Logger
}

func NewClient() *Client {
	return &Client{
		logger: getDefaultLogger(),
	}
}

func NewClientInitialized(accountAlias string, bearerToken string) *Client {
	logger := getDefaultLogger()
	return &Client{
		logger:     logger,
		connection: newConnectionRaw(accountAlias, bearerToken, logger),
	}
}

func getDefaultLogger() *log.Logger {
	return log.New(ioutil.Discard, "", log.LstdFlags)
}

func (cl *Client) Connect(username string, password string) (err error) {
	cl.connection, err = newConnection(username, password, cl.logger)
	return
}

func (cl *Client) SetLogger(logger *log.Logger) (err error) {
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

func (cl *Client) GetDatacenterDeploymentCapabilities(datacenter string) (res *datacenters.GetDatacenterDeploymentCapabilitiesRes, err error) {
	res = &datacenters.GetDatacenterDeploymentCapabilitiesRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("datacenters/{accountAlias}/%s/deploymentCapabilities", datacenter), nil, res)
	return
}

func (cl *Client) GetDatacenterGroup(datacenter string, groupLinks bool) (res *datacenters.GetDatacenterGroupRes, err error) {
	res = &datacenters.GetDatacenterGroupRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("datacenters/{accountAlias}/%s?groupLinks=%t", datacenter, groupLinks), nil, &res)
	return
}

func (cl *Client) GetDatacenterList() (res []*datacenters.GetDatacenterListRes, err error) {
	err = cl.executeRequest("GET", "datacenters/{accountAlias}", nil, &res)
	return
}

func (cl *Client) GetStatus(statusId string) (res *queue.GetStatusRes, err error) {
	res = &queue.GetStatusRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("operations/{acctAlias}/status/%s", statusId), nil, res)
	return
}

func (cl *Client) DeleteAntiAfinityPolicy(policyId string) (err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), nil, nil)
	return
}

func (cl *Client) UpdateAntiAfinityPolicy(policyId string, req *account.UpdateAntiAfinityPolicyReq) (res *account.AntiAfinityPolicyRes, err error) {
	res = &account.AntiAfinityPolicyRes{}
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), req, res)
	return
}

func (cl *Client) CreateAntiAfinityPolicy(policyId string, req *account.CreateAntiAfinityPolicyReq) (res *account.AntiAfinityPolicyRes, err error) {
	res = &account.AntiAfinityPolicyRes{}
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), req, res)
	return
}

func (cl *Client) GetAntiAfinityPolicy(policyId string) (res *account.AntiAfinityPolicyRes, err error) {
	res = &account.AntiAfinityPolicyRes{}
	err = cl.executeRequest("PUT", fmt.Sprintf("antiAffinityPolicies/{accountAlias}/%S", policyId), nil, res)
	return
}

func (cl *Client) GetAntiAfinityPolicies() (res []*account.AntiAfinityPolicyRes, err error) {
	err = cl.executeRequest("PUT", "antiAffinityPolicies/{accountAlias}/%S", nil, &res)
	return
}

func (cl *Client) GetGroup(groupId string) (res *groups.GetGroupRes, err error) {
	res = &groups.GetGroupRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s", groupId), nil, res)
	return
}

func (cl *Client) GetGroupBillingDetails(groupId string) (res *groups.GetGroupBillingDetailsRes, err error) {
	res = &groups.GetGroupBillingDetailsRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("groups/{accountAlias}/%s/billing", groupId), nil, res)
	return
}

func (cl *Client) GetGroupMonitoringStatistics(groupId string, start *time.Time, end *time.Time, sampleInterval *time.Duration, queryType string) (res *groups.GetGroupMonitoringStatisticsRes, err error) {
	res = &groups.GetGroupMonitoringStatisticsRes{}
	params := make(map[string]string)
	url := fmt.Sprintf("groups/{accountAlias}/%s/statistics", groupId)
	if start != nil {
		params["start"] = start.Format("2006-01-02")
	}
	if end != nil {
		params["end"] = end.Format("2006-01-02")
	}
	if sampleInterval != nil {
		params["sampleInterval"] = fmt.Sprintf("%d:%d:%d", sampleInterval.Hours(), sampleInterval.Minutes(), sampleInterval.Seconds())
	}
	if queryType != "" {
		params["queryType"] = queryType
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

func (cl *Client) GetServer(serverId string) (res *servers.GetServerRes, err error) {
	res = &servers.GetServerRes{}
	err = cl.executeRequest("GET", fmt.Sprintf("servers/{accountAlias}/%s", serverId), nil, res)
	return
}

func (cl *Client) PauseServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/pause", req, &res)
	return
}

func (cl *Client) DeleteServer(serverId string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("DELETE", fmt.Sprintf("servers/{accountAlias}/%s", serverId), nil, &res)
	return
}

func (cl *Client) UpdatePublicIpAddress(serverId string, publicIp string, req *servers.UpdatePublicIpAddressReq) (res *models.Link, err error) {
	res = &models.Link{}
	err = cl.executeRequest("PUT", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), req, res)
	return
}

func (cl *Client) RemovePublicIpAddress(serverId string, publicIp string) (res *models.Link, err error) {
	res = &models.Link{}
	err = cl.executeRequest("DELETE", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), nil, res)
	return
}

func (cl *Client) GetPublicIpAddress(serverId string, publicIp string) (res *servers.GetPublicIpAddressRes, err error) {
	res = &servers.GetPublicIpAddressRes{}
	err = cl.executeRequest("PUT", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses/%s", serverId, publicIp), nil, res)
	return
}

func (cl *Client) AddPublicIpAddress(serverId string, publicIp string, req *servers.AddPublicIpAddressReq) (res *models.Link, err error) {
	res = &models.Link{}
	err = cl.executeRequest("POST", fmt.Sprintf("servers/{accountAlias}/%s}/publicIPAddresses", serverId, publicIp), req, res)
	return
}

func (cl *Client) ExecutePackage(req *servers.ExecutePackageReq) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/executePackage", req, &res)
	return
}

func (cl *Client) SetMaintenanceMode(req *servers.SetMaintenanceModeReq) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/setMaintenance", req, &res)
	return
}

func (cl *Client) StartMaintenanceMode(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/startMaintenance", req, &res)
	return
}

func (cl *Client) StopMaintenanceMode(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/stopMaintenance", req, &res)
	return
}

func (cl *Client) CreateServer(req *servers.CreateServerReq) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "servers/{accountAlias}", req, &res)
	return
}

func (cl *Client) CreateSnapshot(req *servers.CreateSnapshotReq) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/createSnapshot", req, &res)
	return
}

func (cl *Client) ShutDownServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/shutDown", req, &res)
	return
}

func (cl *Client) RebootServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/reboot", req, &res)
	return
}

func (cl *Client) ResetServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/reset", req, &res)
	return
}

func (cl *Client) PowerOnServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/powerOn", req, &res)
	return
}

func (cl *Client) PowerOffServer(req []string) (res []servers.ServerRes, err error) {
	err = cl.executeRequest("POST", "operations/{accountAlias}/servers/powerOff", req, &res)
	return
}

func (cl *Client) executeRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	cl.logger.Printf("Sending request to API endpoint: %q, parameters: %#v", url, reqModel)
	if cl.connection == nil {
		err = fmt.Errorf("The client is not initialized. You should call Connect method first.")
		return
	}
	return cl.connection.ExecuteRequest(verb, url, reqModel, resModel)
}
