package file_test

import (
	"github.com/aladinbecovic/zimpler_test/customers"
	"github.com/aladinbecovic/zimpler_test/storage/file"
	"testing"
)

func TestFile_SaveData(t *testing.T) {
	tc := customers.New()
	*tc = append(*tc, &customers.Customer{Name: "Test", FavouriteSnack: "Chips", TotalSnacks: 12})
	*tc = append(*tc, &customers.Customer{Name: "Tes2t", FavouriteSnack: "Chipss", TotalSnacks: 134})

	f, err := file.New()
	if err != nil {
		t.Error(err)
	}

	if f == nil {
		t.Error("File object not created")
	}

	f.SaveData(tc)
}
