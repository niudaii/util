package util

import "strings"

// RemoveDuplicate 去重
func RemoveDuplicate(slice []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func HasSuffixStr(list []string, to string) (string, bool) {
	for _, item := range list {
		if strings.HasSuffix(to, item) {
			return item, true
		}
	}
	return "", false
}

func HasInt(list []int, to int) bool {
	for _, item := range list {
		if to == item {
			return true
		}
	}
	return false
}

func HasString(list []string, to string) bool {
	for _, item := range list {
		if to == item {
			return true
		}
	}
	return false
}
