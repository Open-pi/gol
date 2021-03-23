package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Publishers        []string    `json:"publishers"`
	Identifiers       Identifiers `json:"identifiers"`
	IaBoxID           []string    `json:"ia_box_id"`
	Covers            []int       `json:"covers"`
	LocalID           []string    `json:"local_id"`
	IaLoadedID        []string    `json:"ia_loaded_id"`
	LcClassifications []string    `json:"lc_classifications"`
	Key               string      `json:"key"`
	Authors           []Author    `json:"authors"`
	Ocaid             string      `json:"ocaid"`
	PublishPlaces     []string    `json:"publish_places"`
	Subjects          []string    `json:"subjects"`
	Pagination        string      `json:"pagination"`
	SourceRecords     []string    `json:"source_records"`
	Title             string      `json:"title"`
	DeweyDecimalClass []string    `json:"dewey_decimal_class"`
	Notes             Notes       `json:"notes"`
	NumberOfPages     int         `json:"number_of_pages"`
	Languages         []Languages `json:"languages"`
	Lccn              []string    `json:"lccn"`
	Isbn10            []string    `json:"isbn_10"`
	PublishDate       string      `json:"publish_date"`
	PublishCountry    string      `json:"publish_country"`
	ByStatement       string      `json:"by_statement"`
	OclcNumbers       []string    `json:"oclc_numbers"`
	Works             []Works     `json:"works"`
	Type              Type        `json:"type"`
	LatestRevision    int         `json:"latest_revision"`
	Revision          int         `json:"revision"`
	Created           Time        `json:"created"`
	LastModified      Time        `json:"last_modified"`
	Error             string      `json:"error"`
}

type Identifiers struct {
	Google           []string `json:"google"`
	Lccn             []string `json:"lccn"`
	Isbn13           []string `json:"isbn_13"`
	Amazon           []string `json:"amazon"`
	Isbn10           []string `json:"isbn_10"`
	Oclc             []string `json:"oclc"`
	Librarything     []string `json:"librarything"`
	ProjectGutenberg []string `json:"project_gutenberg"`
	Goodreads        []string `json:"goodreads"`
}

type Notes struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Languages struct {
	Key string `json:"key"`
}

type Works struct {
	Key string `json:"key"`
}

// GetEdition returns a book from its olid
func GetEdition(olid string) (b Book, err error) {
	s := fmt.Sprintf("https://openlibrary.org/books/%s.json", olid)
	resp, err := http.Get(s)
	if err != nil {
		return b, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &b)
	if b.Error == "notfound" {
		return b, errors.New("Book/Edition not found")
	}
	return
}
