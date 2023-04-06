package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func temp(buf *bytes.Buffer, file string) {
	if data, err := os.ReadFile(file); err != nil {
		fmt.Fprint(buf, "N/A")
		return
	} else {
		if v, err := strconv.Atoi(strings.Replace(string(data), "\n", "", 1)); err != nil {
			fmt.Fprint(buf, "N/A")
			return
		} else {
			fmt.Fprintf(buf, "%d°C", v/1000)
			return
		}
	}
}
