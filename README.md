
## Overview

```
radosgw-admin is a RADOS gateway user administration utility.
It allows operations on user/bucket/quota/capability.
```

:warning: This library has been tested with **Jewel** but it should work with **Hammer** too.

|Operation    | Implemented        | Tests               |
|-------------|--------------------|---------------------|
|  User       | :white_check_mark: |:white_check_mark:   |
|  Bucket     | :x:                |:x:                  |
|  Capability | :x:                |:x:                  |
|  Quota      | :x:                |:x:                  |

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
