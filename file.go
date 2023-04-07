package utils

import (
	"bufio"
	"os"
)

// PathExists 判断路径是否存在
func PathExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || err != nil || info == nil {
		return false
	}
	return !info.IsDir()
}

// ReadLines 按行读取文件
func ReadLines(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}

// ReadFile 读取文件 bytes
func ReadFile(filename string) (bytes []byte, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return data, nil
}
