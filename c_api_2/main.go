package main

import "C"
import (
	"runtime"
	"time"
)

func main() {}

//export DO
func DO() {
	s := 2
	s++

	go func() {
		timer := time.NewTicker(time.Second / 20)
		for {
			<-timer.C
			runtime.GC()
		}
	}()
}
