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
	if display == nil {
		panic("Failed to open display.")
	}
	defer C.XCloseDisplay(display)

	for range time.Tick(1 * time.Second) {
		fg(&buf, "282828")
		bg(&buf, "e78a4e")
		temp(&buf, "/sys/class/thermal/thermal_zone2/temp")
		gap(&buf)
		bg(&buf, "a9b665")
		ram(&buf)
		gap(&buf)
		bg(&buf, "7daea3")
		time_24h(&buf)
		gap(&buf)

		if C.XStoreName(display, C.XDefaultRootWindow(display), C.CString(buf.String())) < 0 {
			panic("Failed to allocate memory")
		} else {
			C.XFlush(display)
		}
		buf.Reset()
	}
}
