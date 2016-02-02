package radosAPI

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type UsageConfig struct {
	UID         string
	Start       *time.Time
	End         *time.Time
	ShowEntries bool
	ShowSummary bool
}

func (api *API) GetUsage(conf *UsageConfig) (*Usage, error) {
	ret := &Usage{}
	values := url.Values{}

	values.Add("format", "json")
	if conf != nil {
		if conf.UID != "" {
			values.Add("uid", conf.UID)
		}
		if !conf.ShowEntries {
			values.Add("show-entries", "False")
		}
		if !conf.ShowSummary {
			values.Add("show-summary", "False")
		}
		if conf.Start != nil {
			timeStamp := fmt.Sprintf("%v-%d-%v %v:%v:%v",
				conf.Start.Year(), conf.Start.Month(), conf.Start.Day(),
				conf.Start.Hour(), conf.Start.Minute(), conf.Start.Second())
			values.Add("start", timeStamp)
		}
		if conf.End != nil {
			timeStamp := fmt.Sprintf("%v-%d-%v %v:%v:%v",
				conf.End.Year(), conf.End.Month(), conf.End.Day(),
				conf.End.Hour(), conf.End.Minute(), conf.End.Second())
			values.Add("end", timeStamp)
		}
	}
	body, err := api.get("/admin/usage", values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
