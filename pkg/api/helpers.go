package radosAPI

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// UsageConfig usage request
type UsageConfig struct {
	UID         string     // The user for which the information is requested. If not specified will apply to all users
	Start       *time.Time // Date and (optional) time that specifies the start time of the requested data
	End         *time.Time // Date and (optional) time that specifies the end time of the requested data (non-inclusive)
	ShowEntries bool       // Specifies whether data entries should be returned.
	ShowSummary bool       // Specifies whether data summary should be returned
	RemoveAll   bool       // Required when uid is not specified, in order to acknowledge multi user data removal.
}

// GetUsage requests bandwidth usage information.
//
// !! caps: usage=read !!
//
// @UID
// @Start
// @End
// @ShowEntries
// @ShowSummary
//
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
	body, _, err := api.get("/admin/usage", values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// DeleteUsage removes usage information. With no dates specified, removes all usage information
//
// !! caps: usage=write !!
//
// @UID
// @Start
// @End
// @RemoveAll
//
func (api *API) DeleteUsage(conf *UsageConfig) error {
	values := url.Values{}

	values.Add("format", "json")
	if conf != nil {
		if conf.UID != "" {
			values.Add("uid", conf.UID)
		}
		if conf.RemoveAll {
			values.Add("remove-all", "True")
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
	_, _, err := api.delete("/admin/usage", values)
	if err != nil {
		return err
	}
	return nil
}

// GetUser gets user information. If no user is specified returns the list of all users along with suspension information
//
// !! caps: users=read !!
//
// @uid
//
func (api *API) GetUser(uid ...string) (*User, error) {
	ret := &User{}
	values := url.Values{}

	values.Add("format", "json")
	if len(uid) != 0 {
		values.Add("uid", uid[0])
	}
	body, _, err := api.get("/admin/user", values)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
