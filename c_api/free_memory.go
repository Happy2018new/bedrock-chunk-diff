package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export FreeMemory
func FreeMemory(address unsafe.Pointer) {
	fmt.Println("CALL")
	C.free(address)
}
