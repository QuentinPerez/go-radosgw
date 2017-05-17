package radosAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/smartystreets/go-aws-auth"
)

// API contains fields to communicate with the rados-gateway
type API struct {
	host      string
	accessKey string
	secretKey string
	prefix    string
}

// New returns client for Ceph RADOS Gateway
func New(host, accessKey, secretKey string, adminPrefix ...string) (*API, error) {
	prefix := "admin"
	if len(adminPrefix) > 0 {
		prefix = adminPrefix[0]
	}
	if host == "" || accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("host, accessKey, secretKey must be not nil")
	}
	return &API{host, accessKey, secretKey, prefix}, nil
}

func (api *API) makeRequest(verb, url string) (body []byte, statusCode int, err error) {
	var apiErr apiError
	client := http.Client{}

	// fmt.Printf("URL [%v]: %v\n", verb, url)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return
	}
	awsauth.SignS3(req, awsauth.Credentials{
		AccessKeyID:     api.accessKey,
		SecretAccessKey: api.secretKey,
		Expiration:      time.Now().Add(1 * time.Minute)},
	)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	statusCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if errMarshal := json.Unmarshal(body, &apiErr); errMarshal == nil && apiErr.Code != "" {
		err = errors.New(apiErr.Code)
	}
	return
}

func (api *API) call(verb, route string, args url.Values, usePrefix bool, sub ...string) (body []byte, statusCode int, err error) {
	subreq := ""
	if len(sub) > 0 {
		subreq = fmt.Sprintf("%s&", sub[0])
	}
	if usePrefix {
		route = fmt.Sprintf("/%s%s", api.prefix, route)
	}
	body, statusCode, err = api.makeRequest(verb, fmt.Sprintf("%v%v?%v%s", api.host, route, subreq, args.Encode()))
	if statusCode != 200 {
		err = fmt.Errorf("[%v]: %v", statusCode, err)
	}
	return
}
