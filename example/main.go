package main

import (
	"github.com/zhuharev/postmi"
)

func main() {
	srv, e := postmi.New("cnf")
	if e != nil {
		panic(e)
	}

	srv.Run()
}
