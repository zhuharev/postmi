package leveldb

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/zhuharev/postmi"
)

var (
	AUTOINCREMENT_KEY = []byte("_autoincrement")
)

type LevelDBStore struct {
	db                   *leveldb.DB
	currentAutoIncrement int64
}

func (ls *LevelDBStore) Connect(setting string) error {
	db, e := leveldb.OpenFile(setting, nil)
	if e != nil {
		return e
	}
	ls.db = db
	val, e := db.Get(AUTOINCREMENT_KEY, nil)
	if e != nil {
		if e == leveldb.ErrNotFound {
			val = []byte("0")
		} else {
			return e
		}
	}
	ls.currentAutoIncrement = com.StrTo(string(val)).MustInt64()
	return nil
}

func idKey(id int64) []byte {
	return []byte("p" + fmt.Sprint(id))
}

func (ls *LevelDBStore) getAutoIncrement() (int64, error) {
	if ls.currentAutoIncrement != 0 {
		return ls.currentAutoIncrement, nil
	}
	val, e := ls.db.Get(AUTOINCREMENT_KEY, nil)
	if e != nil {
		if e == leveldb.ErrNotFound {
			e = ls.setAutoIncrement(0)
			if e != nil {
				return 0, e
			}
			return 0, nil
		}
		return 0, e
	}
	ls.currentAutoIncrement = com.StrTo(string(val)).MustInt64()
	return ls.currentAutoIncrement, nil
}

func (ls *LevelDBStore) setAutoIncrement(val int64) error {
	e := ls.db.Put(AUTOINCREMENT_KEY, []byte(fmt.Sprint(val)), nil)
	if e != nil {
		return e
	}
	ls.currentAutoIncrement = val
	return nil
}

func (ls *LevelDBStore) Inc() error {
	cur, e := ls.getAutoIncrement()
	if e != nil {
		return e
	}
	cur++
	return ls.setAutoIncrement(cur)
}

func (ls *LevelDBStore) Save(p *postmi.Post) error {
	if p.Id == 0 {
		e := ls.Inc()
		if e != nil {
			return e
		}
		p.Id = ls.currentAutoIncrement
	}

	e := ls.db.Put(idKey(p.Id), p.MustJSON(), nil)
	if e != nil {
		return e
	}

	return nil
}

func (ls *LevelDBStore) Get(id int64) (*postmi.Post, error) {
	bts, e := ls.db.Get(idKey(id), nil)
	if e != nil {
		return nil, e
	}
	return postmi.NewPostFromJSON(bts)
}

func (ls *LevelDBStore) Delete(id int64) error {
	return ls.db.Delete(idKey(id), nil)
}

func (ls *LevelDBStore) GetSlice(limit int64, offset int64) (posts []*postmi.Post, e error) {
	iter := ls.db.NewIterator(util.BytesPrefix([]byte("p")), nil)
	iter.Last()
	var p *postmi.Post
	p, e = postmi.NewPostFromJSON(iter.Value())
	if e != nil {
		return
	}
	posts = append(posts, p)
	for iter.Prev() {
		offset--
		if offset > 0 {
			continue
		}
		p, e = postmi.NewPostFromJSON(iter.Value())
		if e != nil {
			break
		}
		posts = append(posts, p)
		if int64(len(posts)) == limit {
			break
		}
	}
	iter.Release()
	if e != nil {
		return
	}
	e = iter.Error()
	return
}
