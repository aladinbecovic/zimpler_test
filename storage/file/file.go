package file

import (
	"encoding/json"
	"github.com/aladinbecovic/zimpler_test/customers"
	"os"
)

type File struct {
	*os.File
}

func New() (*File, error) {
	f, err := os.OpenFile("fileStorage.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return &File{f}, nil
}

func (f *File) SaveData(tc *customers.TopCustomers) error {
	jsonData, err := json.MarshalIndent(tc, "", " ")
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) LoadData() error {

	return nil
}

func (f *File) Close() error {
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
