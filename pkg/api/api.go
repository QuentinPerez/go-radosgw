package radosAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/smartystreets/go-aws-auth"
)

type API struct {
	host      string
	accessKey string
	secretKey string
}

// New returns client for Ceph RADOS Gateway
func New(host, accessKey, secretKey string) *API {
	return &API{host, accessKey, secretKey}
}

func (api *API) get(route string, args url.Values) (body []byte, statusCode int, err error) {
	client := http.Client{}
	url := fmt.Sprintf("%v%v?%s", api.host, route, args.Encode())
	fmt.Println("URL:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	awsauth.SignS3(req, awsauth.Credentials{
		AccessKeyID:     api.accessKey,
		SecretAccessKey: api.secretKey,
		Expiration:      time.Now().Add(1 * time.Minute)},
	)
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, 0, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	statusCode = resp.StatusCode
	return
}

func (api *API) delete(route string, args url.Values) (body []byte, statusCode int, err error) {
	client := http.Client{}
	url := fmt.Sprintf("%v%v?%s", api.host, route, args.Encode())
	fmt.Println("URL:", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}
	awsauth.SignS3(req, awsauth.Credentials{
		AccessKeyID:     api.accessKey,
		SecretAccessKey: api.secretKey,
		Expiration:      time.Now().Add(1 * time.Minute)},
	)
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, 0, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	statusCode = resp.StatusCode
	return
}
