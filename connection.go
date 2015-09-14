package centurylink_sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"

	"github.com/s-matyukevich/centurylink_sdk/errors"
	"github.com/s-matyukevich/centurylink_sdk/models"
	"github.com/s-matyukevich/centurylink_sdk/models/authentication"
)

//this made a variable instead of a constant for testing purpoises
var BaseUrl = "https://api.ctl.io"

const (
	API_VERSION = "/v2/"
)

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
	err = cn.ExecuteRequest("POST", API_VERSION + "authentication/login", loginReq, loginRes)
	if err != nil {
		return
	}
	cn.bearerToken = loginRes.BearerToken
	cn.accountAlias = loginRes.AccountAlias
	cn.logger.Printf("Updating connection. Bearer: %s, Alias: %s", cn.bearerToken, cn.accountAlias)
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
		err := cn.DecodeResponse(res, resModel)
		if err != nil {
			cn.logger.Printf(err.Error())
		}
		// FIXME: this is not descriptive enough
		// FIXME: check if body contains JSON { "message": <descriptive text> }
		//        and return that
		return &errors.ApiError{
			StatusCode:  res.StatusCode,
			ApiResponse: resModel,
		}
	}
	err = cn.DecodeResponse(res, resModel)
	if err != nil {
		return
	}
	if linkModel, ok := resModel.(models.LinkModel); ok {
		linkModel.SetConnection(cn)
	}
	if reflect.TypeOf(resModel).Elem().Kind() == reflect.Slice {
		s := reflect.ValueOf(resModel).Elem()

		for i := 0; i < s.Len(); i++ {
			item := s.Index(i).Interface()
			if linkModel, ok := item.(models.LinkModel); ok {
				linkModel.SetConnection(cn)
			}
		}
	}
	return err
}

func (cn *connection) DecodeResponse(res *http.Response, resModel interface{}) (err error) {
	if resModel == nil {
		return
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(resModel)
	return
}
