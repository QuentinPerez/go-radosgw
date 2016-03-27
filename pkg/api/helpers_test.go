package radosAPI

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAPI(t *testing.T) {
	Convey("Testing New API", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		So(api, ShouldNotBeNil)
	})
}

func TestNewAPIWithPrefix(t *testing.T) {
	Convey("Testing New API with prefix", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		So(api, ShouldNotBeNil)
	})
}

func TestCreateUser(t *testing.T) {
	Convey("Testing Create user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			UID:         "UnitTest",
			DisplayName: "Unit Test",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.DisplayName, ShouldEqual, "Unit Test")
	})
}

func TestCreateUserWithoutUID(t *testing.T) {
	Convey("Testing Create user without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			DisplayName: "Unit Test",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestCreateUserWithoutName(t *testing.T) {
	Convey("Testing Create user without name", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"), "admin")

		user, err := api.CreateUser(UserConfig{
			UID: "UnitTest",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestGetUser(t *testing.T) {
	Convey("Testing Get user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.GetUser("UnitTest")
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
	})
}

func TestUpdateUser(t *testing.T) {
	Convey("Testing Update user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.UpdateUser(UserConfig{
			UID:   "UnitTest",
			Email: "UnitTest@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Email, ShouldEqual, "UnitTest@test.com")
	})
}

func TestUpdateUserWithoutUID(t *testing.T) {
	Convey("Testing Update user without UID", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		user, err := api.UpdateUser(UserConfig{
			Email: "UnitTest@test.com",
		})
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}

func TestRemoveUser(t *testing.T) {
	Convey("Testing Remove user", t, func() {
		api := New(os.Getenv("RADOSGW_API"), os.Getenv("RADOSGW_ACCESS"), os.Getenv("RADOSGW_SECRET"))

		err := api.RemoveUser(UserConfig{
			UID:       "UnitTest",
			PurgeData: true,
		})
		So(err, ShouldBeNil)
		user, err := api.GetUser("UnitTest")
		So(err, ShouldNotBeNil)
		So(user, ShouldBeNil)
	})
}
