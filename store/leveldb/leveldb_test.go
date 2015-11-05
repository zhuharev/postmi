package leveldb

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhuharev/postmi"
	"testing"
)

func TestOpenDB(t *testing.T) {
	Convey("Test open database", t, func() {
		db := new(LevelDBStore)
		e := db.Connect("db.test")
		Convey("The error should be nil", func() {
			So(e, ShouldBeNil)
		})
		db.db.Close()
	})
}

func TestInsertArticle(t *testing.T) {
	Convey("Test insert a article", t, func() {
		db := new(LevelDBStore)
		e := db.Connect("db.test")
		defer db.db.Close()
		Convey("The error should be nil", func() {
			So(e, ShouldBeNil)
		})

		p := new(postmi.Post)
		p.Title = "ololo"
		e = db.Save(p)
		So(e, ShouldBeNil)
		So(p.Id, ShouldNotEqual, 0)
	})
}

func TestGetArticle(t *testing.T) {
	Convey("test get article", t, func() {
		db := new(LevelDBStore)
		e := db.Connect("db.test")
		Convey("The error should be nil", func() {
			So(e, ShouldBeNil)
		})
		defer db.db.Close()

		p, e := db.Get(1)
		So(e, ShouldBeNil)
		So(p.Id, ShouldEqual, 1)
		So(p.Title, ShouldEqual, "ololo")
	})
}
