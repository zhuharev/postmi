package main

import (
	"github.com/zhuharev/postmi"
	_ "github.com/zhuharev/postmi/store/leveldb"
)

func main() {
	srv, e := postmi.New("cnf")
	if e != nil {
		panic(e)
	}

	srv.Run()
}
