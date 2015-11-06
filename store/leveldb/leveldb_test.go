package leveldb

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhuharev/postmi"
	"os"
	"testing"
)

func init() {
	os.RemoveAll("./db.test")
}

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

func TestGetSliceArticle(t *testing.T) {
	Convey("test get slice of articles", t, func() {
		db := new(LevelDBStore)
		e := db.Connect("db.test")
		Convey("The error should be nil", func() {
			So(e, ShouldBeNil)
		})
		defer db.db.Close()

		p := new(postmi.Post)
		p.Title = "ololo"
		e = db.Save(p)

		p = new(postmi.Post)
		p.Title = "alolo"
		e = db.Save(p)

		posts, e := db.GetSlice(10, 0)
		So(e, ShouldBeNil)
		So(len(posts), ShouldEqual, 3)
		So(posts[0].Id, ShouldEqual, 3)
		So(posts[2].Id, ShouldEqual, 1)
	})
}
