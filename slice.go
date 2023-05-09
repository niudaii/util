package util

import "strings"

// RemoveDuplicate 去重
func RemoveDuplicate(list []string) []string {
	var set []string
	hashSet := make(map[string]struct{})
	for _, v := range list {
		hashSet[v] = struct{}{}
	}
	for k := range hashSet {
		// 去除空字符串
		if k == "" {
			continue
		}
		set = append(set, k)
	}
	return set
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
