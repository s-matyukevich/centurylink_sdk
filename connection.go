package centurylink_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/s-matyukevich/centurylink_sdk/models/authentication"
)

//this made a variable instead of a constant for testing purpoises
var BaseUrl = "https://api.tier3.com/v2/"

type connection struct {
	bearerToken  string
	accountAlias string
	logger       *log.Logger
}

func newConnection(username string, password string, logger *log.Logger) (cn *connection, err error) {
	cn = &connection{
		logger: logger,
	}
	cn.logger.Printf("Creating new connection. Username: %s", username)
	loginReq := &authentication.LoginReq{username, password}
	loginRes := &authentication.LoginRes{}
	err = cn.ExecuteRequest("POST", "authentication/login", loginReq, loginRes)
	if err != nil {
		return
	}
	cn.bearerToken = loginRes.BearerToken
	cn.accountAlias = loginRes.AccountAlias
	cn.logger.Printf("Updateing connection. Bearer: %s, Alias: %s", cn.bearerToken, cn.accountAlias)
	return
}

func newConnectionRaw(accountAlias string, bearerToken string, logger *log.Logger) (cn *connection) {
	return &connection{
		bearerToken:  bearerToken,
		accountAlias: accountAlias,
		logger:       logger,
	}
}

func (cn *connection) ExecuteRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	req, err := cn.prepareRequest(verb, url, reqModel)
	if err != nil {
		return
	}
	reqDump, _ := httputil.DumpRequest(req, true)
	cn.logger.Printf("Sending request: %s", reqDump)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	resDump, _ := httputil.DumpResponse(res, true)
	cn.logger.Printf("Response received: %s", resDump)
	err = cn.processResponse(res, resModel)
	return
}

func (cn *connection) prepareRequest(verb string, url string, reqModel interface{}) (req *http.Request, err error) {
	var inputData io.Reader
	if reqModel != nil {
		b, err := json.Marshal(reqModel)
		if err != nil {
			return nil, err
		}
		inputData = bytes.NewReader(b)
		cn.logger.Printf("Input model converted to json: %s", b)
	}
	url = BaseUrl + strings.Replace(url, "{accountAlias}", cn.accountAlias, 1)
	req, err = http.NewRequest(verb, url, inputData)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return
	}
	if cn.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+cn.bearerToken)
	}
	return
}

func (cn *connection) processResponse(res *http.Response, resModel interface{}) (err error) {
	switch res.StatusCode {
	case 200, 201, 202:
	default:
		err = fmt.Errorf("Error occured while sending request to API. Status code: %d", res.StatusCode)
		return
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(resModel)
	return err
}
