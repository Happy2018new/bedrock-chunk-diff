package main

import "C"
import (
	"fmt"

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
