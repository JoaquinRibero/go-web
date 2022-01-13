package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data []byte
	Err  error
}

const (
	FileType Type = "filestorage"
	TestType Type = "teststorage"
)

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock(mock *Mock) {
	fs.Mock = nil
}

func (fs *FileStore) Read(data interface{}) error {

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}

	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) { //controlo que el error no sea porque no existe el archivo.
			return err
		}
		file = []byte("[]") // devuelvo un array vacío si no existe el archivo.
	}
	return json.Unmarshal(file, &data)
}

func (fs *FileStore) Write(data interface{}) error {

	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return nil
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	fl, err := os.OpenFile(fs.FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = fl.Write(file)
	if err != nil {
		return err
	}
	return nil
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName, nil}
	}
	return nil
}
