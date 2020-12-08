package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize get the file size
func GetSize(file multipart.File) (int, error) {
	content, err := ioutil.ReadAll(file)
	return len(content), err
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMKDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MKDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir create a directory
func MKDir(src string) error {
	if err := os.Mkdir(src, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err:%v", err)
	}

	src := dir + "\\" + fileName
	perm := CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermisson denied src: %s", src)
	}
	err = IsNotExistMKDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMKDir src:%s,err:%v ", src, err)
	}
	file, err := Open(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, fmt.Errorf("open file:%s err:%v", fileName, err)
	}
	return file, nil
}
