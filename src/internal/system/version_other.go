//go:build !windows

package system

func windowsVersion() (uint32, uint32) {
	return 0, 0
}
