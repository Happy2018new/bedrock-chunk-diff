package main

import "C"
import (
	"fmt"

	"github.com/TriM-Organization/bedrock-chunk-diff/timeline"
	"go.etcd.io/bbolt"
)

func main() {}

//export DO
func DO() {
	db, err := bbolt.Open("ss", 0600, &bbolt.Options{
		FreelistType: bbolt.FreelistMapType,
		NoGrowSync:   false,
		NoSync:       false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//export DO2
func DO2() {
	db, err := timeline.Open("ss", false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.CloseTimelineDB()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//export DO3
func DO3(path *C.char) {
	tldb, err := timeline.Open(C.GoString(path), false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tldb.CloseTimelineDB()
	if err != nil {
		fmt.Println(err)
		return
	}
}
