package utils

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
)

// Storage 文件存储
type Storage struct {
	storagePath string // 文件存储路径
	name        string // 文件名
}

// NewStorage 新建文件存储
func NewStorage(storagePath, fileName string) (*Storage, error) {
	if _, err := os.Stat(storagePath); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(storagePath, os.ModePerm)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &Storage{storagePath: storagePath, name: fileName}, nil
}

// Get 解析文件存储
func (s *Storage) Get(value interface{}) error {
	var filepath = path.Join(s.storagePath, s.name)
	return storageRead(filepath, value)
}

// Store 缓存文件存储
func (s *Storage) Store(value interface{}) error {
	var filepath = path.Join(s.storagePath, s.name)
	return storageWrite(filepath, value)
}

// Del 删除文件存储
func (s *Storage) Del() error {
	var filepath = path.Join(s.storagePath, s.name)
	return os.Remove(filepath)
}

// 读文件存储
func storageRead(storagePath string, value interface{}) error {
	f, err := os.OpenFile(storagePath, os.O_RDWR, 0666)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return err
	}
	return json.NewDecoder(bufio.NewReader(f)).Decode(&value)
}

// 写文件存储
func storageWrite(storagePath string, value interface{}) error {
	content, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, content, os.ModePerm)
}
