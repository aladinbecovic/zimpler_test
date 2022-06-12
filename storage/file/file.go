package file

import (
	"encoding/json"
	"github.com/aladinbecovic/zimpler_test/customers"
	"io/ioutil"
	"os"
)

type FileStorage struct {
	*os.File
}

func New() (*FileStorage, error) {
	f, err := os.OpenFile("fileStorage.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &FileStorage{f}, nil
}

func (f *FileStorage) SaveData(tc *customers.TopCustomers) error {
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

func (f *FileStorage) LoadData(tc *customers.TopCustomers) error {
	fRes, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return err
	}

	err = json.Unmarshal(fRes, &tc)
	if err != nil {
		return err
	}

	return nil
}
