package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/TriM-Organization/bedrock-chunk-diff/timeline"
	"go.etcd.io/bbolt"
)

var savedTimelineDB = NewSimpleManager[timeline.TimelineDatabase]()

func main() {}

//export FreeMemory
func FreeMemory(address unsafe.Pointer) {
	C.free(address)
}

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
func DO3(path *C.char) C.longlong {
	tldb, err := timeline.Open(C.GoString(path), false, false)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	err = tldb.CloseTimelineDB()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return C.longlong(savedTimelineDB.AddObject(tldb))
}
