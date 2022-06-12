package scrape

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aladinbecovic/zimpler_test/customers"
	"net/http"
	"strconv"
	"strings"
)

func GetHTMLDoc(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ExtractTopCustomerTableData(tc *customers.TopCustomers) error {
	doc, err := GetHTMLDoc("https://candystore.zimpler.net/#candystore-customers")
	if err != nil {
		return err
	}

	doc.Find("table.summary").Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		str, err := s.Html()
		if err != nil {
			fmt.Println(err)
			return
		}

		customer := new(customers.Customer)

		/* Extract Data from HTML */
		splitRes := strings.Split(str, "\"")
		nrSnacks, err := strconv.Atoi(splitRes[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		textRes := s.Text()
		textRes = strings.ReplaceAll(textRes, " ", "")
		textResSplit := strings.Split(textRes, "\n")
		fmt.Println(textResSplit)

		customer.Name = textResSplit[1]
		customer.FavouriteSnack = textResSplit[2]
		customer.TotalSnacks = nrSnacks

		*tc = append(*tc, customer)
	})

	return nil
}
