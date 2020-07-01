// +build windows

package host

import (
	"os"
	"syscall"
	"unsafe"
)

func uname() string {
	return ""
}

func network() string {
	return ""
}

func resolvConf() string {
	return ""
}

func etcHosts() string {
	return ""
}

func mountedPaths() (disks []string) {
	for _, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Stat(string(c) + ":\\")
		if err == nil {

			disks = append(disks, c)
		}
	}
}

type df struct {
	total int64
	avail int64
}

func makeDf(path string) *df {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	df := &df{}

	c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&df.total)),
		uintptr(unsafe.Pointer(&df.avail)))

	return df
}

func (d *df) total() uint64 {
	return uint64(d.total)
}

func (d *df) available() uint64 {
	return uint64(d.avail)
}
