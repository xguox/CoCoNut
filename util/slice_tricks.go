package util

import (
	"sort"
)

func UniqSlice(strSlice []string) []string {
	sort.Strings(strSlice)
	j := 0
	for i := 1; i < len(strSlice); i++ {
		if strSlice[j] == strSlice[i] {
			continue
		}
		j++
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	result := strSlice[:j+1]
	return result
}

func SliceContains(sl []string, str string) bool {
	for _, a := range sl {
		if a == str {
			return true
		}
	}
	return false
}
