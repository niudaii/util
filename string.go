package util

import "strings"

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Upper 字符串全部大写
func Upper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s)
}

// Lower 字符串全部小写
func Lower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s)
}
