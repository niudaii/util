package util

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
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

// WriteFile 写入文件
func WriteFile(filename string, data string) (err error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(data)
	_ = writer.Flush()
	return nil
}

// SaveMarshal json 格式保存
func SaveMarshal(filename string, results interface{}) (err error) {
	var data []byte
	data, err = json.Marshal(results)
	if err != nil {
		return
	}
	err = WriteFile(filename, string(data))
	return
}

// GetAllFile 获取指定目录下所有文件
func GetAllFile(dirPath string) (results []string, err error) {
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			results = append(results, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
