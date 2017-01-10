package radosAPI

import (
	"bytes"
	"os"
	"testing"
	"time"

	"strings"

	minio "github.com/minio/minio-go"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	Url    = os.Getenv("RADOSGW_API")
	Access = os.Getenv("RADOSGW_ACCESS") // USER WITH ALL CAPABILITIES
	Secret = os.Getenv("RADOSGW_SECRET")
)

func createNewAPI() *API {
	api, err := New(Url, Access, Secret)
	if err != nil {
		panic(err)
	}
	return api
}

func TestAPI(t *testing.T) {
	Convey("Testing New API", t, func() {
		api := createNewAPI()

		So(api, ShouldNotBeNil)
	})

	Convey("Testing New API with prefix", t, func() {
		api, err := New(Url, Access, Secret, "adminEndpoint")
		if err != nil {
			panic(err)
		}
		So(api, ShouldNotBeNil)
	})

	Convey("Testing New API without arguments", t, func() {
		api, err := New("", "", "")
		So(api, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestUser(t *testing.T) {
	Convey("Testing Create user", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")
	})

	Convey("Testing Create user without UID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			DisplayName: "Unit Test",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})

	Convey("Testing Create user without name", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})

	Convey("Testing Get user", t, func() {
		api := createNewAPI()

		user, err := api.GetUser("UnitTest")
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
	})

	Convey("Testing Update user", t, func() {
		api := createNewAPI()

		user, err := api.UpdateUser(UserConfig{
			UID:   "UnitTest",
			Email: "UnitTest@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Email, ShouldEqual, "unittest@test.com")
	})

	Convey("Testing Update user without UID", t, func() {
		api := createNewAPI()

		user, err := api.UpdateUser(UserConfig{
			Email: "UnitTest@test.com",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})

	Convey("Testing Remove user", t, func() {
		api := createNewAPI()

		err := api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
		user, err := api.GetUser("UnitTest")
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})

	Convey("Testing Remove user without UID", t, func() {
		api := createNewAPI()

		err := api.RemoveUser(UserConfig{
			PurgeData: true,
		})
		So(err, ShouldNotBeNil)
	})
}

func TestUsage(t *testing.T) {
	Convey("Testing Get Usage with empty struct", t, func() {
		api := createNewAPI()

		usage, err := api.GetUsage(UsageConfig{})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldBeNil)
		So(usage.Summary, ShouldBeNil)
	})

	Convey("Testing Get Usage summary field", t, func() {
		api := createNewAPI()

		usage, err := api.GetUsage(UsageConfig{
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})

	Convey("Testing Get Usage entries field", t, func() {
		api := createNewAPI()

		usage, err := api.GetUsage(UsageConfig{
			ShowEntries: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldBeNil)
	})

	Convey("Testing Get Usage entries/summary field", t, func() {
		api := createNewAPI()

		usage, err := api.GetUsage(UsageConfig{
			ShowEntries: true,
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})

	Convey("Testing Get Usage entries/summary field with specified uid", t, func() {
		api := createNewAPI()

		usage, err := api.GetUsage(UsageConfig{
			UID:         "UnitTest",
			ShowEntries: true,
			ShowSummary: true,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})

	Convey("Testing Get Usage entries/summary field with specified uid, and start Time", t, func() {
		api := createNewAPI()

		now := time.Now().AddDate(0, 0, -1)
		usage, err := api.GetUsage(UsageConfig{
			UID:         "UnitTest",
			ShowEntries: true,
			ShowSummary: true,
			Start:       &now,
		})
		So(err, ShouldBeNil)
		So(usage, ShouldNotBeNil)
		So(usage.Entries, ShouldNotBeNil)
		So(usage.Summary, ShouldNotBeNil)
	})

	Convey("Testing Delete all usages", t, func() {
		api := createNewAPI()

		err := api.DeleteUsage(UsageConfig{
			UID:       "UnitTest",
			RemoveAll: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestKey(t *testing.T) {
	Convey("Testing Create Key", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing Create Key without UID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{})
		So(err, ShouldNotBeNil)
		So(keys, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing Remove Key", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveKey(KeyConfig{
			AccessKey: (*keys)[0].AccessKey,
		})

		So(err, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing Remove Key without access key", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		keys, err := api.CreateKey(KeyConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(keys, ShouldNotBeNil)

		err = api.RemoveKey(KeyConfig{})

		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestSubUser(t *testing.T) {
	Convey("Testing CreateSubUser", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing CreateSubUser without UID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			SubUser: "SubUnitTest",
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing UpdateSubUser", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing UpdateSubUser without UID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing UpdateSubUser without SubUID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		sub, err = api.UpdateSubUser(SubUserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(sub, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing RemoveSubUser", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
	Convey("Testing RemoveSubUser without UID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			SubUser: (*sub)[0].ID,
		})
		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing RemoveSubUser without SubUID", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")

		sub, err := api.CreateSubUser(SubUserConfig{
			UID:     "UnitTest",
			SubUser: "SubUnitTest",
		})
		So(err, ShouldBeNil)
		So(sub, ShouldNotBeNil)

		err = api.RemoveSubUser(SubUserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)

		err = api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
	})
}

func TestBucket(t *testing.T) {
	Convey("Testing Get Bucket with stats", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)

		buckets, err := api.GetBucket(BucketConfig{
			UID:   "UnitTest",
			Stats: true,
		})
		So(err, ShouldBeNil)
		So(buckets, ShouldNotBeNil)
	})

	Convey("Testing Get Bucket without stats", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)

		buckets, err := api.GetBucket(BucketConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(buckets, ShouldNotBeNil)
	})

	Convey("Testing Remove Bucket", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)

		api.RemoveBucket(BucketConfig{
			Bucket:       "unittestbucket",
			PurgeObjects: true,
		})

		So(err, ShouldBeNil)
	})

	Convey("Testing Remove Bucket without bucket name", t, func() {
		api := createNewAPI()

		err := api.RemoveBucket(BucketConfig{})

		So(err, ShouldNotBeNil)
	})

	Convey("Testing Unlink Bucket without UID", t, func() {
		api := createNewAPI()

		err := api.UnlinkBucket(BucketConfig{
			Bucket: "UnitTest",
		})

		So(err, ShouldNotBeNil)
	})

	Convey("Testing Unlink Bucket without Bucket Name", t, func() {
		api := createNewAPI()

		err := api.UnlinkBucket(BucketConfig{
			UID: "unittest",
		})

		So(err, ShouldNotBeNil)
	})

	Convey("Testing Unlink Bucket", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)
		err = api.UnlinkBucket(BucketConfig{
			Bucket: "unittestbucket",
			UID:    "UnitTest",
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing Unlink Invalid Bucket", t, func() {
		api := createNewAPI()

		err := api.UnlinkBucket(BucketConfig{
			Bucket: "unittestbucket",
			UID:    "UnitTest",
		})
		So(err, ShouldNotBeNil)
	})

	Convey("Testing Check index without Bucket name", t, func() {
		api := createNewAPI()

		_, err := api.CheckBucket(BucketConfig{})
		So(err, ShouldNotBeNil)
	})

	Convey("Testing Check index Bucket", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)
		index, err := api.CheckBucket(BucketConfig{
			Bucket: "unittestbucket",
		})
		So(err, ShouldBeNil)
		So(index, ShouldNotBeNil)
	})

	Convey("Testing Remove Object", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)
		b := bytes.NewBufferString("content")
		_, err = minioClient.PutObject("unittestbucket", "test.txt", b, "")
		So(err, ShouldBeNil)

		err = api.RemoveObject(BucketConfig{
			Bucket: "unittestbucket",
			Object: "test.txt",
		})
		So(err, ShouldBeNil)
	})

	Convey("Testing Remove Object without bucket", t, func() {
		api := createNewAPI()

		err := api.RemoveObject(BucketConfig{})
		So(err, ShouldNotBeNil)
	})

	Convey("Testing Remove Object without object", t, func() {
		api := createNewAPI()

		err := api.RemoveObject(BucketConfig{
			Bucket: "unittestbucket",
		})
		So(err, ShouldNotBeNil)
	})

	Convey("Testing Get Bucket Policy", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)

		policy, err := api.GetBucketPolicy(BucketConfig{})
		So(err, ShouldNotBeNil)
		policy, err = api.GetBucketPolicy(BucketConfig{
			Bucket: "unittestbucket",
		})
		So(err, ShouldBeNil)
		So(policy, ShouldNotBeNil)
	})

	Convey("Testing Get Object Policy", t, func() {
		api := createNewAPI()

		url := ""
		useSSL := false
		if strings.HasPrefix(Url, "http://") {
			url = Url[7:]
		} else if strings.HasPrefix(Url, "https://") {
			url = Url[8:]
			useSSL = true
		}
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		minioClient, err := minio.New(url, user.Keys[0].AccessKey, user.Keys[0].SecretKey, useSSL)
		So(err, ShouldBeNil)
		err = minioClient.MakeBucket("unittestbucket", "")
		So(err, ShouldBeNil)
		b := bytes.NewBufferString("content")
		_, err = minioClient.PutObject("unittestbucket", "test.txt", b, "")
		So(err, ShouldBeNil)

		policy, err := api.GetObjectPolicy(BucketConfig{})
		So(err, ShouldNotBeNil)

		policy, err = api.GetObjectPolicy(BucketConfig{
			Bucket: "unittestbucket",
		})
		So(err, ShouldNotBeNil)

		policy, err = api.GetObjectPolicy(BucketConfig{
			Bucket: "unittestbucket",
			Object: "test.txt",
		})
		So(err, ShouldBeNil)
		So(policy, ShouldNotBeNil)
	})
}

func TestQuota(t *testing.T) {
	Convey("Testing  Quota Without arguments", t, func() {
		api := createNewAPI()

		quotas, err := api.GetQuotas(QuotaConfig{})
		So(quotas, ShouldBeNil)
		So(err, ShouldNotBeNil)
		err = api.UpdateQuota(QuotaConfig{})
		So(err, ShouldNotBeNil)
		err = api.UpdateQuota(QuotaConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
	})

	Convey("Testing Get Quota", t, func() {
		api := createNewAPI()
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()
		quotas, err := api.GetQuotas(QuotaConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(quotas, ShouldNotBeNil)
	})

	Convey("Testing Update Quota", t, func() {
		api := createNewAPI()
		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()
		quotas, err := api.GetQuotas(QuotaConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(quotas, ShouldNotBeNil)
		err = api.UpdateQuota(QuotaConfig{
			UID:       "UnitTest",
			MaxSizeKB: "1000",
			QuotaType: "user",
		})
		So(err, ShouldBeNil)
		quotas, err = api.GetQuotas(QuotaConfig{
			UID: "UnitTest",
		})
		So(err, ShouldBeNil)
		So(quotas, ShouldNotBeNil)
		So(quotas.UserQuota.MaxSizeKb, ShouldEqual, 1000)
	})
}

func TestCapability(t *testing.T) {
	Convey("Testing AddCapability Without arguments", t, func() {
		api := createNewAPI()

		cap, err := api.AddCapability(CapConfig{})
		So(err, ShouldNotBeNil)
		So(cap, ShouldBeNil)
		cap, err = api.AddCapability(CapConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(cap, ShouldBeNil)
	})

	Convey("Testing AddCapability", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		cap, err := api.AddCapability(CapConfig{
			UID:      "UnitTest",
			UserCaps: "usage=*",
		})
		So(cap, ShouldNotBeNil)
		So(err, ShouldBeNil)
		found := false
		for _, c := range cap {
			if c.Type == "usage" && c.Perm == "*" {
				found = true
			}
		}
		So(found, ShouldEqual, true)
	})

	Convey("Testing DelCapability Without arguments", t, func() {
		api := createNewAPI()

		cap, err := api.DelCapability(CapConfig{})
		So(err, ShouldNotBeNil)
		So(cap, ShouldBeNil)
		cap, err = api.DelCapability(CapConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(cap, ShouldBeNil)
	})

	Convey("Testing DelCapability", t, func() {
		api := createNewAPI()

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		defer func() {
			err = api.RemoveUser(UserConfig{
				UID:       "UnitTest",
				PurgeData: true,
			})
			So(err, ShouldBeNil)
		}()

		cap, err := api.AddCapability(CapConfig{
			UID:      "UnitTest",
			UserCaps: "usage=*",
		})
		So(cap, ShouldNotBeNil)
		So(err, ShouldBeNil)
		found := false
		for _, c := range cap {
			if c.Type == "usage" && c.Perm == "*" {
				found = true
			}
		}

		So(found, ShouldEqual, true)
		cap, err = api.DelCapability(CapConfig{
			UID:      "UnitTest",
			UserCaps: "usage=*",
		})
		So(cap, ShouldNotBeNil)
		So(err, ShouldBeNil)
		found = false
		for _, c := range cap {
			if c.Type == "usage" && c.Perm == "*" {
				found = true
			}
		}
		So(found, ShouldEqual, false)
	})
}
