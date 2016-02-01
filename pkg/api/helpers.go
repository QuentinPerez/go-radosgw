package radosAPI

import (
	"encoding/json"
	"net/url"
)

func (api *API) GetUsage() (*Usage, error) {
	ret := &Usage{}
	values := url.Values{}

	values.Add("format", "json")
	body, err := api.get("/admin/usage", values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
