#include <stdio.h>
#include <stdlib.h>
#include "caller.h"


void caller (void){
    printf("called the function in c \n");
    fflush(stdout);
    callback cb;
    watch* w;
    *w = 100;
    ret* r;
    foo(cb,w,r);
    printf("received return %d \n", *r);
}