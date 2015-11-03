package leveldb

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOpenDB(t *testing.T) {
	Convey("Test open database", t, func() {
		db := new(LevelDBStore)
		e := db.Connect("db.test")
		Convey("The error should be nil", func() {
			So(e, ShouldBeNil)
		})
	})
}
