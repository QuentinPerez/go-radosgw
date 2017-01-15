
## Overview

```
radosgw-admin is a RADOS gateway user administration utility.
It allows operations on user/bucket/quota/capability.
```

:warning: This library has been tested with **Jewel** but it should work with **Hammer** too.

|Operation    | Implemented        | Tests               |
|-------------|--------------------|---------------------|
|  User       | :white_check_mark: |:white_check_mark:   |
|  Bucket     | :white_check_mark: |:white_check_mark:   |
|  Capability | :white_check_mark: |:white_check_mark:   |
|  Quota      | :white_check_mark: |:white_check_mark:   |

## Setup

```
$> go get github.com/QuentinPerez/go-radosgw/pkg/api
```

## How it works

> Ensure you have cluster ceph with a radosgw.

```go
api := radosAPI.New("http://192.168.42.40", "1ZZWD0G5IDP57I0751HE", "3ydvK64eWuWwup0FKtznmf9FDVXhB8jleEFRTH0D")

// create a new user named JohnDoe
user, err := api.CreateUser(radosAPI.UserConfig{
    UID:         "JohnDoe",
    DisplayName: "John Doe",
})
// ...
// remove JohnDoe
err = api.RemoveUser(radosAPI.UserConfig{
    UID: "JohnDoe",
})
```

## API

```go
// New returns an API object to intertact with Admin RadosGW
func New(host, accessKey, secretKey string, adminPrefix ...string) (*API, error) {}

// GetUsage requests bandwidth usage information.
func (api *API) GetUsage(conf UsageConfig) (*Usage, error) {}

// DeleteUsage removes usage information. With no dates specified, removes all usage information
func (api *API) DeleteUsage(conf UsageConfig) error {}

// GetUser gets user information. If no user is specified returns the list of all users along with suspension information
func (api *API) GetUser(uid ...string) (*User, error) {}

// CreateUser creates a new user. By Default, a S3 key pair will be created automatically and returned in the response.
func (api *API) CreateUser(conf UserConfig) (*User, error) {}

// UpdateUser modifies a user
func (api *API) UpdateUser(conf UserConfig) (*User, error) {}

// RemoveUser removes an existing user.
func (api *API) RemoveUser(conf UserConfig) error {}

// CreateSubUser creates a new subuser (primarily useful for clients using the Swift API).
func (api *API) CreateSubUser(conf SubUserConfig) (*SubUsers, error) {}

// UpdateSubUser modifies an existing subuser
func (api *API) UpdateSubUser(conf SubUserConfig) (*SubUsers, error) {}

// RemoveSubUser remove an existing subuser
func (api *API) RemoveSubUser(conf SubUserConfig) error {}

// CreateKey creates a new key. If a subuser is specified then by default created keys will be swift type.
func (api *API) CreateKey(conf KeyConfig) (*KeysDefinition, error) {}

// RemoveKey removes an existing key
func (api *API) RemoveKey(conf KeyConfig) error {}

// GetBucket gets information about a subset of the existing buckets.
func (api *API) GetBucket(conf BucketConfig) (Buckets, error) {}

// RemoveBucket removes an existing bucket.
func (api *API) RemoveBucket(conf BucketConfig) error {}

// UnlinkBucket unlinks a bucket from a specified user.
func (api *API) UnlinkBucket(conf BucketConfig) error {}

// CheckBucket checks the index of an existing bucket.
func (api *API) CheckBucket(conf BucketConfig) (string, error) {}

// LinkBucket links a bucket to a specified user, unlinking the bucket from any previous user.
func (api *API) LinkBucket(conf BucketConfig) error {}

// RemoveObject removes an existing object.
func (api *API) RemoveObject(conf BucketConfig) error {}

// GetBucketPolicy reads the bucket policy
func (api *API) GetBucketPolicy(conf BucketConfig) (*Policy, error) {}

// GetObjectPolicy reads the object policy
func (api *API) GetObjectPolicy(conf BucketConfig) (*Policy, error) {}

// GetQuotas returns user's quotas
func (api *API) GetQuotas(conf QuotaConfig) (*Quotas, error) {}

// UpdateQuota updates user's quotas
func (api *API) UpdateQuota(conf QuotaConfig) error {}

// AddCapability returns user's quotas
func (api *API) AddCapability(conf CapConfig) ([]Capability, error) {}

// DelCapability returns user's quotas
func (api *API) DelCapability(conf CapConfig) ([]Capability, error) {}
```

## Changelog

### master (unreleased)

---

## Development

Feel free to contribute :smiley::beers:

## Links

- **Radowsgw-admin Documentaion**: http://docs.ceph.com/docs/jewel/radosgw/adminops/
- **Report bugs**: https://github.com/QuentinPerez/go-radosgw/issues

## License

[MIT](https://github.com/QuentinPerez/go-radosgw/blob/master/LICENSE)
