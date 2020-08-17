#ifndef CALLER_H
#define CALLER_H

typedef int watch;
typedef int ret;
typedef  void (*callback)(watch*);

extern void foo (callback,watch*, ret*);
void caller (void);
#endif