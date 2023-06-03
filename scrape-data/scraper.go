package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

type Scraper struct {
	Client *http.Client
}

func (s *Scraper) Init() {
	// Initialize any necessary configurations or dependencies for the scraper.
	s.Client = &http.Client{}
}

func (s *Scraper) FetchPage(url string) (*http.Response, error) {
	// Send an HTTP GET request to fetch the page.
	resp, err := s.Client.Get(url)
	if err != nil {
		return nil, err
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	return resp, nil
}

func (s *Scraper) ParsePage(resp *http.Response) (*goquery.Document, error) {
	// Parse the HTML document using goquery.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (s *Scraper) ExtractData(doc *goquery.Document) {
	// Extract the desired data from the parsed document.
	doc.Find("h1").Each(func(i int, sel *goquery.Selection) {
		fmt.Printf("Heading 1 - %d: %s\n", i+1, sel.Text())
	})

	doc.Find("h2").Each(func(i int, sel *goquery.Selection) {
		fmt.Printf("Heading 2 - %d: %s\n", i+1, sel.Text())
	})
}

func main() {
	scraper := Scraper{}
	scraper.Init()

	// Read the URL from command-line argument
	if len(os.Args) < 2 {
		log.Fatal("URL argument is required")
	}
	url := os.Args[1]
	resp, err := scraper.FetchPage(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := scraper.ParsePage(resp)
	if err != nil {
		log.Fatal(err)
	}

	scraper.ExtractData(doc)
}
