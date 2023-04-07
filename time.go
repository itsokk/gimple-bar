package main

import (
	"bytes"
	"fmt"
	"time"
)

func time_24h(buf *bytes.Buffer) {
	fmt.Fprint(buf, time.Now().Format(" 15:04 "))
}

func time_12h(buf *bytes.Buffer) {
	fmt.Fprint(buf, time.Now().Format(" 15:04 AM "))
}
