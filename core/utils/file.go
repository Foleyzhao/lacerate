package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

// CreateFile 创建文件
func CreateFile(dir string, name string) (string, error) {
	src := path.Join(dir, name)
	_, err := os.Stat(src)
	if os.IsExist(err) {
		return src, nil
	}
	if err := os.MkdirAll(dir, 0777); err != nil {
		if os.IsPermission(err) {
			panic("insufficient permissions")
		}
		return "", err
	}
	_, err = os.Create(src)
	if err != nil {
		return "", err
	}
	return src, nil
}

// MkDir 创建路径
func MkDir(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// CopyFile 复制文件
func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)
	desFile, err := os.Create(des)
	if err != nil {
		return 0, err
	}
	defer func(desFile *os.File) {
		_ = desFile.Close()
	}(desFile)
	return io.Copy(desFile, srcFile)
}

// CopyDir 复制路径
func CopyDir(source string, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errors.New("source is not a directory")
	}
	_, err = os.Open(dest)
	if os.IsExist(err) {
		err = os.RemoveAll(dest)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(source)
	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				panic(err)
			}
		} else {
			_, err = CopyFile(sfp, dfp)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

// WriteFile 写文件
func WriteFile(file string, text string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		_ = fmt.Errorf("open file error: %s", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	w := bufio.NewWriter(f)
	_, err = w.Write([]byte(text))
	if err != nil {
		return err
	}
	return w.Flush()
}
