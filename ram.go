package main

import (
	"bytes"
	"fmt"
	"syscall"
)

func ram(buf *bytes.Buffer) {
	var si syscall.Sysinfo_t
	if err := syscall.Sysinfo(&si); err != nil {
		fmt.Fprint(buf, "N/A")
	}
	total_ram := float64(si.Totalram * uint64(si.Unit)) / (1024 * 1024 * 1024)
	used_ram := float64((si.Totalram - si.Freeram - si.Bufferram - si.Sharedram) * uint64(si.Unit)) / (1024 * 1024)
	if (used_ram < 1000) {
		fmt.Fprintf(buf, " %.1fM/%.1fG ", used_ram, total_ram)
		return
	}

	fmt.Fprintf(buf, " %.1fG/%.1fG ", (used_ram / 1024), total_ram)
}
