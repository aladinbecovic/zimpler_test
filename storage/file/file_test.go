package file_test

import (
	"github.com/aladinbecovic/zimpler_test/customers"
	"github.com/aladinbecovic/zimpler_test/storage/file"
	"testing"
)

func TestFileStorage_SaveData(t *testing.T) {
	tc := customers.New()
	*tc = append(*tc, &customers.Customer{Name: "Test", FavouriteSnack: "Chips", TotalSnacks: 12})
	*tc = append(*tc, &customers.Customer{Name: "Tes2t", FavouriteSnack: "Chipss", TotalSnacks: 134})

	f, err := file.New()
	defer f.Close()
	if err != nil {
		t.Error(err)
	}

	f.SaveData(tc)
}

func TestFileStorage_LoadData(t *testing.T) {
	tc := customers.New()

	f, err := file.New()
	if err != nil {
		t.Error(err)
	}

	if err := f.LoadData(tc); err != nil {
		t.Error(err)
	}
}
