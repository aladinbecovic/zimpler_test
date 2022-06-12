package scrape

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
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

func GetCustomerTableData() error {
	doc, err := GetHTMLDoc("https://candystore.zimpler.net/#candystore-customers")
	if err != nil {
		return err
	}

	doc.Find("table.summary").Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	return nil
}
