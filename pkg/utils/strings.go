package utils

import (
	"sort"
)

func OrderString(s string) string {
	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}
