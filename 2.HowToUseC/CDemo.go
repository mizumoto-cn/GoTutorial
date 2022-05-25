package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

// You need CString to form a C format char[] string
// still C functions shall be written in comments
// which is called "regenerate cgo definitions"
// also remember to free the memory manually
