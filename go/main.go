package main

import "fmt"

/*
#cgo CFLAGS: -I./../c
#cgo LDFLAGS: -L./../c -lexcaller
#include <stdio.h>
#include "caller.h"
*/
import "C"

//export GoCallee
func GoCallee(c C.callback, w *C.watch, r *C.ret) {
	fmt.Println("In go func")
	if *w== 100{
		*r = 555
	}
}


func main(){
	C.caller()
}