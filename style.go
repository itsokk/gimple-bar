package main

import (
	"bytes"
	"fmt"
)

func fg(buf *bytes.Buffer, color string) int {
	if v, err := fmt.Fprintf(buf, "^c#%s^", color); err != nil {
		return 0
	} else {
		return v
	}
}

func bg(buf *bytes.Buffer, color string) int {
	if v, err := fmt.Fprintf(buf, "^b#%s^", color); err != nil {
		return 0
	} else {
		return v
	}
}

const gap_str string = "^f10^"

func gap(buf *bytes.Buffer) {
	fmt.Fprint(buf, gap_str)
}
