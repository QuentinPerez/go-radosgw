package radosAPI

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/QuentinPerez/go-encodeUrl"
)

// UsageConfig usage request
type UsageConfig struct {
	UID         string     `url:"uid,ifStringIsNotEmpty"`     // The user for which the information is requested. If not specified will apply to all users
	Start       *time.Time `url:"start,ifTimeIsNotNilCeph"`   // Date and (optional) time that specifies the start time of the requested data
	End         *time.Time `url:"end,ifTimeIsNotNilCeph"`     // Date and (optional) time that specifies the end time of the requested data (non-inclusive)
	ShowEntries bool       `url:"show-entries,ifBoolIsFalse"` // Specifies whether data entries should be returned.
	ShowSummary bool       `url:"show-summary,ifBoolIsFalse"` // Specifies whether data summary should be returned
	RemoveAll   bool       `url:"remove-all,ifBoolIsTrue"`    // Required when uid is not specified, in order to acknowledge multi user data removal.
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
	var (
		ret    = &Usage{}
		values = url.Values{}
		errs   []error
	)

	values.Add("format", "json")
	if conf != nil {
		values, errs = encurl.Translate(conf)
		if len(errs) > 0 {
			return nil, errs[0]
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
	var (
		values = url.Values{}
		errs   []error
	)

	values.Add("format", "json")
	if conf != nil {
		values, errs = encurl.Translate(conf)
		if len(errs) > 0 {
			return errs[0]
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
