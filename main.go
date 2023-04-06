package main

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"
import (
	"bytes"
	"time"
)

func main() {
	var buf bytes.Buffer
    
	display := C.XOpenDisplay(nil)
	if(display == nil) {
		panic("Failed to open display.")
	}
    defer C.XCloseDisplay(display)

	for range time.Tick(1 * time.Second){
		fg(&buf, "282828")
		bg(&buf, "e78a4e")
		time_24h(&buf)

		if (C.XStoreName(display, C.XDefaultRootWindow(display), C.CString(buf.String())) < 0) {
			panic("Failed to allocate memory")
		} else {
			C.XFlush(display)
		}
		buf.Reset()
	}
}

