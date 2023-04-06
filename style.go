package main

import (
	"bytes"
	"fmt"
)

func fg(buf *bytes.Buffer, color string) {
	fmt.Fprintf(buf, "^c#%s^", color)
}

func bg(buf *bytes.Buffer, color string) {
	fmt.Fprintf(buf, "^b#%s^", color)
}

const gap_str string = "^f10^"

func gap(buf *bytes.Buffer) {
	fmt.Fprint(buf, gap_str)
}
