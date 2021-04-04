package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Book holds all the information returned from the EditionAPI
// TODO: Add Id field, table of contents, series
type Book struct {
	Publishers        []string    `json:"publishers"`
	Identifiers       Identifiers `json:"identifiers"`
	IaBoxID           []string    `json:"ia_box_id"`
	Covers            []int       `json:"covers"`
	LocalID           []string    `json:"local_id"`
	IaLoadedID        []string    `json:"ia_loaded_id"`
	LcClassifications []string    `json:"lc_classifications"`
	Key               string      `json:"key"`
	AuthorsKey        []AuthorKey `json:"authors"`
	Ocaid             string      `json:"ocaid"`
	PublishPlaces     []string    `json:"publish_places"`
	Subjects          []string    `json:"subjects"`
	Pagination        string      `json:"pagination"`
	SourceRecords     []string    `json:"source_records"`
	Title             string      `json:"title"`
	SubTitle          string      `json:"subtitle"`
	EditionName       string      `json:"edition_name"`
	DeweyDecimalClass []string    `json:"dewey_decimal_class"`
	Notes             Notes       `json:"notes"`
	NumberOfPages     int         `json:"number_of_pages"`
	TranslatedFrom    []Language  `json:"translated_from"`
	Languages         []Language  `json:"languages"`
	Lccn              []string    `json:"lccn"`
	Isbn10            []string    `json:"isbn_10"`
	Isbn13            []string    `json:"isbn_13"`
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

// Identifiers of the books
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

// Notes on books available in OpenLibrary
type Notes struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Language of book
type Language struct {
	Key string `json:"key"`
}

// Works
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

// GetEditionISBN returns a book from its isbnid
func GetEditionISBN(isbnid string) (b Book, err error) {
	isbnid = strings.ReplaceAll(isbnid, "-", "")
	if len(isbnid) != 10 && len(isbnid) != 13 {
		return b, errors.New("incorrect ISBN ID length, must be 10 or 13")
	} else if len(isbnid) == 13 && isbnid[:3] != "978" {
		return b, errors.New("incorrect ISBN-13 ID prefix, must be 978")
	}
	s := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbnid)
	resp, err := http.Get(s)
	if resp.StatusCode == 404 {
		return b, errors.New("ISBN not found")
	}
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

// KeyAuthors returns array of all authors keys
func (b Book) KeyAuthors() []string {
	a := make([]string, len(b.AuthorsKey))
	for i, AuthorKey := range b.AuthorsKey {
		a[i] = AuthorKey.Key[9:]
	}
	return a
}

// Authors returns all the information related to the book's authors
func (b Book) Authors() (a []Author, err error) {
	return Authors(b)
}

// KeyCover returns (if it exists) the ID of the work's cover
func (b Book) KeyCover() string {
	if len(b.Covers) > 0 {
		return strconv.Itoa(b.Covers[0])
	}
	return ""
}

// Cover returns (if it exists) the URL of the Book's Cover
func (b Book) Cover(size string) string {
	return Cover(b, size)
}
