package radosAPI

import (
	"encoding/json"
	"errors"
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
func (api *API) GetUsage(conf UsageConfig) (*Usage, error) {
	var (
		ret    = &Usage{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("GET", "/usage", values, true)
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
func (api *API) DeleteUsage(conf UsageConfig) error {
	var (
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("DELETE", "/usage", values, true)
	return err
}

// GetUser gets user information. If no user is specified returns the list of all users along with suspension information
//
// ** If no user is specified returns the list ... ** Don't works for me
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
	body, _, err := api.call("GET", "/user", values, true)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// UserConfig user request
type UserConfig struct {
	UID         string `url:"uid,ifStringIsNotEmpty"`          // The user ID to be created
	DisplayName string `url:"display-name,ifStringIsNotEmpty"` // The display name of the user to be created
	Email       string `url:"email,ifStringIsNotEmpty"`        // The email address associated with the user
	KeyType     string `url:"key-type,ifStringIsNotEmpty"`     // Key type to be generated, options are: swift, s3 (default)
	AccessKey   string `url:"access-key,ifStringIsNotEmpty"`   // Specify access key
	SecretKey   string `url:"secret-key,ifStringIsNotEmpty"`   // Specify secret key
	UserCaps    string `url:"user-caps,ifStringIsNotEmpty"`    // User capabilities
	MaxBuckets  *int   `url:"max-buckets,itoaIfNotNil"`        // Specify the maximum number of buckets the user can own
	GenerateKey bool   `url:"generate-key,ifBoolIsTrue"`       // Generate a new key pair and add to the existing keyring
	Suspended   bool   `url:"suspended,ifBoolIsTrue"`          // Specify whether the user should be suspended
	PurgeData   bool   `url:"suspended,ifBoolIsTrue"`          // When specified the buckets and objects belonging to the user will also be removed
}

// CreateUser creates a new user. By Default, a S3 key pair will be created automatically and returned in the response.
// If only one of access-key or secret-key is provided, the omitted key will be automatically generated.
// By default, a generated key is added to the keyring without replacing an existing key pair.
// If access-key is specified and refers to an existing key owned by the user then it will be modified
//
// !! caps: users=write !!
//
// @UID
// @DisplayName
// @Email
// @KeyType
// @AccessKey
// @SecretKey
// @UserCaps
// @GenerateKey
// @MaxBuckets
// @Suspended
//
func (api *API) CreateUser(conf UserConfig) (*User, error) {
	if conf.UID == "" {
		return nil, errors.New("UID field is required")
	}
	if conf.DisplayName == "" {
		return nil, errors.New("DisplayName field is required")
	}

	var (
		ret    = &User{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("PUT", "/user", values, true)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// UpdateUser modifies a user
//
// !! caps: users=write !!
//
// @UID
// @DisplayName
// @Email
// @KeyType
// @AccessKey
// @SecretKey
// @UserCaps
// @GenerateKey
// @MaxBuckets
// @Suspended
//
func (api *API) UpdateUser(conf UserConfig) (*User, error) {
	if conf.UID == "" {
		return nil, errors.New("UID field is required")
	}

	var (
		ret    = &User{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("POST", "/user", values, true)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// RemoveUser removes an existing user.
//
// !! caps: users=write !!
//
// @UID
// @PurgeData
//
func (api *API) RemoveUser(conf UserConfig) error {
	if conf.UID == "" {
		return errors.New("UID field is required")
	}
	var (
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("DELETE", "/user", values, true)
	return err
}

// SubUserConfig subuser request
type SubUserConfig struct {
	UID            string `url:"uid,ifStringIsNotEmpty"`        // The user ID under which a subuser is to be created
	SubUser        string `url:"subuser,ifStringIsNotEmpty"`    // Specify the subuser ID to be created
	KeyType        string `url:"key-type,ifStringIsNotEmpty"`   // Key type to be generated, options are: swift (default), s3
	Access         string `url:"access,ifStringIsNotEmpty"`     // Set access permissions for sub-user, should be one of read, write, readwrite, full
	Secret         string `url:"secret,ifStringIsNotEmpty"`     // Specify secret key
	SecretKey      string `url:"secret-key,ifStringIsNotEmpty"` // Specify secret key
	GenerateSecret bool   `url:"generate-secret,ifBoolIsTrue"`  // Generate the secret key
	PurgeKeys      bool   `url:"purge-keys,ifBoolIsTrue"`       // Remove keys belonging to the subuser
}

// CreateSubUser creates a new subuser (primarily useful for clients using the Swift API).
// Note that either gen-subuser or subuser is required for a valid request.
// Note that in general for a subuser to be useful, it must be granted permissions by specifying access.
// As with user creation if subuser is specified without secret, then a secret key will be automatically generated.
//
// !! caps:	users=write !!
//
// @UID
// @SubUser
// @KeyType
// @Access
// @SecretKey
// @GenerateSecret
//
func (api *API) CreateSubUser(conf SubUserConfig) (*SubUsers, error) {
	if conf.UID == "" {
		return nil, errors.New("UID field is required")
	}

	var (
		ret    = &SubUsers{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("PUT", "/user", values, true, "subuser")
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// UpdateSubUser modifies an existing subuser
//
// !! caps:	users=write !!
//
// @UID
// @SubUser
// @KeyType
// @Access
// @Secret
// @GenerateSecret
//
func (api *API) UpdateSubUser(conf SubUserConfig) (*SubUsers, error) {
	if conf.UID == "" {
		return nil, errors.New("UID field is required")
	}
	if conf.SubUser == "" {
		return nil, errors.New("SubUser field is required")
	}

	var (
		ret    = &SubUsers{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("POST", "/user", values, true, "subuser")
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// RemoveSubUser remove an existing subuser
//
// !! caps:	users=write !!
//
// @UID
// @SubUser
// @PurgeKeys
//
func (api *API) RemoveSubUser(conf SubUserConfig) error {
	if conf.UID == "" {
		return errors.New("UID field is required")
	}
	if conf.SubUser == "" {
		return errors.New("SubUser field is required")
	}
	var (
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("DELETE", "/user", values, true, "subuser")
	return err
}

// KeyConfig key request
type KeyConfig struct {
	UID            string `url:"uid,ifStringIsNotEmpty"`        // The user ID to receive the new key
	SubUser        string `url:"subuser,ifStringIsNotEmpty"`    // The subuser ID to receive the new key
	KeyType        string `url:"key-type,ifStringIsNotEmpty"`   // Key type to be generated, options are: swift, s3 (default)
	AccessKey      string `url:"access-key,ifStringIsNotEmpty"` // Specify the access key
	SecretKey      string `url:"secret-key,ifStringIsNotEmpty"` // Specify secret key
	GenerateSecret bool   `url:"generate-secret,ifBoolIsTrue"`  // Generate a new key pair and add to the existing keyring
}

// CreateKey creates a new key. If a subuser is specified then by default created keys will be swift type.
// If only one of access-key or secret-key is provided the committed key will be automatically generated,
// that is if only secret-key is specified then access-key will be automatically generated.
// By default, a generated key is added to the keyring without replacing an existing key pair.
// If access-key is specified and refers to an existing key owned by the user then it will be modified.
// The response is a container listing all keys of the same type as the key created.
// Note that when creating a swift key, specifying the option access-key will have no effect.
// Additionally, only one swift key may be held by each user or subuser.
//
// !! caps:	users=write !!
//
// @UID
// @SubUser
// @KeyType
// @AccessKey
// @SecretKey
// @GenerateSecret
//
func (api *API) CreateKey(conf KeyConfig) (*KeysDefinition, error) {
	if conf.UID == "" {
		return nil, errors.New("UID field is required")
	}

	var (
		ret    = &KeysDefinition{}
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("PUT", "/user", values, true, "key")
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// RemoveKey removes an existing key
//
// !! caps:	users=write !!
//
// @UID
// @SubUser
// @KeyType
// @AccessKey
//
func (api *API) RemoveKey(conf KeyConfig) error {
	if conf.AccessKey == "" {
		return errors.New("AccessKey field is required")
	}

	var (
		values = url.Values{}
		errs   []error
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("DELETE", "/user", values, true, "key")
	return err
}

// BucketConfig bucket request
type BucketConfig struct {
	Bucket       string `url:"bucket,ifStringIsNotEmpty"`  // The bucket to return info on
	UID          string `url:"uid,ifStringIsNotEmpty"`     // The user to retrieve bucket information for
	Stats        bool   `url:"stats,ifBoolIsTrue"`         // Return bucket statistics
	CheckObjects bool   `url:"check-objects,ifBoolIsTrue"` // Check multipart object accounting
	Fix          bool   `url:"fix,ifBoolIsTrue"`           // Also fix the bucket index when checking
	PurgeObjects bool   `url:"purge-objects,ifBoolIsTrue"` // Remove a buckets objects before deletion
}

// GetBucket gets information about a subset of the existing buckets.
// If uid is specified without bucket then all buckets beloning to the user will be returned.
// If bucket alone is specified, information for that particular bucket will be retrieved
//
// !! caps:	buckets=read !!
//
//@Bucket
//@UID
//@Stats
//
func (api *API) GetBucket(conf BucketConfig) (Buckets, error) {
	var (
		ret     = Buckets{}
		values  = url.Values{}
		errs    []error
		variant interface{}
	)

	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	values.Add("format", "json")
	body, _, err := api.call("GET", "/bucket", values, true)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &variant); err != nil {
		return nil, err
	}
	if tab, ok := variant.([]interface{}); ok {
		add := bucket{}
		for _, v := range tab {
			if name, ok := v.(string); ok {
				if add.Name != "" {
					ret = append(ret, add)
					add = bucket{}
				}
				add.Name = name
			} else {
				js, errMarshal := json.Marshal(v)
				if errMarshal != nil {
					return nil, errMarshal
				}
				if add.Stats != nil {
					ret = append(ret, add)
					add = bucket{}
				}
				add.Stats = new(stats)
				err = json.Unmarshal(js, add.Stats)
				if err != nil {
					return nil, err
				}
			}
		}
		ret = append(ret, add)
	} else {
		add := bucket{}
		add.Stats = new(stats)
		err = json.Unmarshal(body, add.Stats)
		if err != nil {
			return nil, err
		}
		ret = append(ret, add)
	}
	return ret, nil
}

// RemoveBucket removes an existing bucket.
//
// !! caps:	buckets=write !!
//
//@Bucket
//@PurgeObjects
//
func (api *API) RemoveBucket(conf BucketConfig) error {
	var (
		values = url.Values{}
		errs   []error
	)

	if conf.Bucket == "" {
		return errors.New("Bucket field is required")
	}
	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("DELETE", "/bucket", values, true)
	return err
}

// UnlinkBucket unlinks a bucket from a specified user.
// Primarily useful for changing bucket ownership.
//
// !! caps:	buckets=write !!
//
//@Bucket
//@UID
//
func (api *API) UnlinkBucket(conf BucketConfig) error {
	var (
		values = url.Values{}
		errs   []error
	)

	if conf.Bucket == "" {
		return errors.New("Bucket field is required")
	}
	if conf.UID == "" {
		return errors.New("UID field is required")
	}
	values, errs = encurl.Translate(conf)
	if len(errs) > 0 {
		return errs[0]
	}
	values.Add("format", "json")
	_, _, err := api.call("POST", "/bucket", values, true)
	return err
}
