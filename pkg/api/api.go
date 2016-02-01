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

func New(host, accessKey, secretKey string) *API {
	return &API{host, accessKey, secretKey}
}

func (api *API) get(route string, args url.Values) (body []byte, err error) {
	client := http.Client{}
	url := fmt.Sprintf("%v%v?%s", api.host, route, args.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	awsauth.SignS3(req, awsauth.Credentials{AccessKeyID: api.accessKey, SecretAccessKey: api.secretKey, Expiration: time.Now().Add(1 * time.Minute)})
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
