package main

/*
#include <stdio.h>
#include "caller.h"
extern void GoCallee(callback, watch*, ret*);
void foo (callback c, watch* w, ret* r){
	printf("calling the go function in Cgo \n");
	fflush(stdout);
	GoCallee(c,w,r);
}
*/
import "C"
