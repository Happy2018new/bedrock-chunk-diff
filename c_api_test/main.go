package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"encoding/binary"
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

func asCbool(b bool) C.int {
	if b {
		return C.int(1)
	}
	return C.int(0)
}

func asGoBool(b C.int) bool {
	return (int(b) != 0)
}

func asCbytes(b []byte) *C.char {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(len(b)))
	result = append(result, b...)
	return (*C.char)(C.CBytes(result))
}

func asGoBytes(p *C.char) []byte {
	l := binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(p), 4))
	return C.GoBytes(unsafe.Pointer(p), C.int(4+l))[4:]
}

//export NewTimelineDB
func NewTimelineDB(path *C.char, noGrowSync C.int, noSync C.int) C.longlong {
	tldb, err := timeline.Open(C.GoString(path), asGoBool(noGrowSync), asGoBool(noSync))
	if err != nil {
		return -1
	}
	return C.longlong(savedTimelineDB.AddObject(tldb))
}
