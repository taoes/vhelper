package utils

import "strconv"

func AddPrefix(x int, prefix string) string {
	if x < 10 {
		return prefix + strconv.Itoa(x)
	} else {
		return strconv.Itoa(x)
	}
}
