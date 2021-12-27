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
}

const (
	FileType Type = "filestorage"
)

func (fs *FileStore) Read(data interface{}) error {
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
		return &FileStore{fileName}
	}
	return nil
}
