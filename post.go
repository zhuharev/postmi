package postmi

import (
	"encoding/json"
	"time"
)

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
}

func NewPostFromJSON(data []byte) (*Post, error) {
	p := new(Post)
	e := json.Unmarshal(data, p)
	return p, e
}

func (p *Post) JSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Post) MustJSON() []byte {
	bts, _ := p.JSON()
	return bts
}
