//go:build windows

package system

import (
	"syscall"
	"unsafe"
)

type osVersionInfo struct {
	size        uint32
	major       uint32
	minor       uint32
	build       uint32
	platformID  uint32
	servicePack [128]uint16
}

func windowsVersion() (uint32, uint32) {
	procedure := syscall.NewLazyDLL("ntdll.dll").NewProc("RtlGetVersion")
	if procedure.Find() != nil {
		return 0, 0
	}
	var info osVersionInfo
	info.size = uint32(unsafe.Sizeof(info))
	procedure.Call(uintptr(unsafe.Pointer(&info)))
	return info.major, info.build
}
