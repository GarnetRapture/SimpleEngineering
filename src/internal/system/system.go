package system

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const windows11FirstBuild = 22000

type Info struct {
	OS             string
	Arch           string
	WindowsMajor   uint32
	WindowsBuild   uint32
	LinuxName      string
	LinuxID        string
	LinuxIDLike    []string
	WSL            bool
	MacVersion     string
	PackageManager string
}

func (info Info) IsWindows11() bool {
	return info.OS == "windows" && info.WindowsMajor == 10 && info.WindowsBuild >= windows11FirstBuild
}

func (info Info) IsWindows10() bool {
	return info.OS == "windows" && info.WindowsMajor == 10 && info.WindowsBuild < windows11FirstBuild
}

func Detect() Info {
	info := Info{OS: runtime.GOOS, Arch: runtime.GOARCH}
	switch info.OS {
	case "windows":
		info.WindowsMajor, info.WindowsBuild = windowsVersion()
		if Has("winget") {
			info.PackageManager = "winget"
		}
	case "linux":
		readOSRelease(&info)
		info.WSL = detectWSL()
		for _, manager := range []struct{ command, key string }{
			{"apt-get", "apt"}, {"dnf", "dnf"}, {"pacman", "pacman"}, {"zypper", "zypper"},
		} {
			if Has(manager.command) {
				info.PackageManager = manager.key
				break
			}
		}
	case "darwin":
		if output, err := exec.Command("sw_vers", "-productVersion").Output(); err == nil {
			info.MacVersion = strings.TrimSpace(string(output))
		}
		if Has("brew") {
			info.PackageManager = "brew"
		}
	}
	return info
}

func Has(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func HasMSVC() bool {
	programFiles := os.Getenv("ProgramFiles(x86)")
	if programFiles == "" {
		return false
	}
	vswhere := filepath.Join(programFiles, "Microsoft Visual Studio", "Installer", "vswhere.exe")
	if _, err := os.Stat(vswhere); err != nil {
		return false
	}
	output, err := exec.Command(vswhere, "-latest", "-products", "*", "-requires", "Microsoft.VisualStudio.Component.VC.Tools.x86.x64", "-property", "installationPath").Output()
	return err == nil && strings.TrimSpace(string(output)) != ""
}

func readOSRelease(info *Info) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value, found := strings.Cut(scanner.Text(), "=")
		if !found {
			continue
		}
		value = strings.Trim(value, `"'`)
		switch key {
		case "PRETTY_NAME":
			info.LinuxName = value
		case "ID":
			info.LinuxID = value
		case "ID_LIKE":
			info.LinuxIDLike = strings.Fields(value)
		}
	}
}

func detectWSL() bool {
	if os.Getenv("WSL_DISTRO_NAME") != "" {
		return true
	}
	version, err := os.ReadFile("/proc/version")
	return err == nil && strings.Contains(strings.ToLower(string(version)), "microsoft")
}
