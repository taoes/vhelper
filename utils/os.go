package utils

import (
	"runtime"
)

func GetOsType() string {
	goos := runtime.GOOS
	if goos == "linux" {
		return "LINUX"
	}

	if goos == "windows" {
		return "WINDOWS"
	}

	return "OTHER"
}
