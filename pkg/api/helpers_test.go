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
