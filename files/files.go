package files

/*

References:
https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
https://www.tutorialspoint.com/how-to-check-if-a-file-exists-in-golang

type FileInfo interface {   // this struct is returned in os.Stat()
	Name() string             // base name of the file
	Size() int64              // length in bytes for regular files; system-dependent for others
	Mode() FileMode           // file mode bits
	ModTime() time.Time       // modification time
	IsDir() bool              // abbreviation for Mode().IsDir()
	Sys() any                 // underlying data source (can return nil)
}

*/

import (
  "io/ioutil"
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

func Exists(src string) bool { // Check if a file or directory exists and return true or false
	_, err := os.Stat(src)
	if os.IsNotExist(err) { // Check if err shows that file doesn't exist
		return false
	}
	return true
}

func IsFile(src string) (bool, error) { // Check if especified path is a file
	info, err := os.Stat(src)
	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}

func IsDir(src string) (bool, error) { // Check if especified path is a directory
	info, err := os.Stat(src)
	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

func GetContent(src string) (string, error) {
	byte_content, err := ioutil.ReadFile(src)
	if err != nil {
		return "", err
	}

	return string(byte_content), nil
}

func WriteContent(filename string, text string) error {
	err := os.WriteFile(filename, []byte(text), 0644)

	if err != nil {
		return err
	}

	return nil
}

func Move(src string, dst string) error {
	err := os.Rename(src, dst)
	if err != nil {
		return err
	}

	return nil
}

func Copy(src string, dst string) error { // Copy file or directory (recursive)
	check, err := IsFile(src)
	if err != nil {
		return err
	}

	if check == true { // Enter here if especified source is a file
		file_bytes, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(dst, file_bytes, 0644)
		if err != nil {
			return err
		}

	} else if check == false { // Enter here if especified source is a directory
		src_info, err := os.Stat(src)
		if err != nil {
			return err
		}

		err = os.MkdirAll(dst, src_info.Mode())
		if err != nil {
			return err
		}

		directory, _ := os.Open(src)
		objects, err := directory.Readdir(-1)

		for _, obj := range objects { // Iterate over files and dirs
			srcfilepointer := src + "/" + obj.Name()
			dstfilepointer := dst + "/" + obj.Name()

			if obj.IsDir() {
				err = Copy(srcfilepointer, dstfilepointer)
				if err != nil {
					return err
				}
			} else {
				file_bytes, err := ioutil.ReadFile(srcfilepointer)
				if err != nil {
					return err
				}
				err = ioutil.WriteFile(dstfilepointer, file_bytes, 0644)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
