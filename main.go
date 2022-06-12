package main

import (
	"fmt"
	"github.com/aladinbecovic/zimpler_test/customers"
	"github.com/aladinbecovic/zimpler_test/scrape"
)

func main() {
	tc := customers.New()

	scrape.ExtractTopCustomerTableData(tc)

	fmt.Println("hello")

}
