package main

import (
	"github.com/aladinbecovic/zimpler_test/customers"
	"github.com/aladinbecovic/zimpler_test/scrape"
	"github.com/aladinbecovic/zimpler_test/storage/file"
	"log"
)

func main() {
	f, err := file.New()
	if err != nil {
		log.Fatal(err)
	}
	tc := customers.New()

	if err := scrape.ExtractTopCustomerTableData(tc); err != nil {
		log.Fatal(err)
	}
	tc.SortTopCustomers()

	if err := f.SaveData(tc); err != nil {
		log.Fatal(err)
	}

	f.Close()
}
