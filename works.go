package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Work holds all the information about works coming from openlibrary
type Work struct {
	Created          Time               `json:"created"`
	Subjects         []string           `json:"subjects"`
	LatestRevision   int                `json:"latest_revision"`
	Key              string             `json:"key"`
	Title            string             `json:"title"`
	Subtitle         string             `json:"subtitle"`
	FirstPublishDate string             `json:"first_publish_date"`
	AuthorsKey       []AuthorKeyAndType `json:"authors"`
	Type             Type               `json:"type"`
	LastModified     Time               `json:"last_modified"`
	Covers           []int              `json:"covers"`
	Revision         int                `json:"revision"`
	Error            string             `json:"error"`
	NumberOfEditions int
}

// GetWork returns the work from the workID
func GetWork(id string) (w Work, err error) {
	s := fmt.Sprintf("https://openlibrary.org/works/%s.json", id)
	resp, err := http.Get(s)
	if err != nil {
		return w, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &w)
	if w.Error == "notfound" {
		return w, errors.New("Work not found")
	}
	return
}

// KeyCover returns (if it exists) the ID of the work's cover
func (w Work) KeyCover() string {
	if len(w.Covers) > 0 {
		return strconv.Itoa(w.Covers[0])
	}
	return ""
}

// Cover returns the cover url to the "first" edition.
// It takes size as an argument; it can be (S, M, or L)
func (w Work) Cover(size string) string {
	return Cover(w, size)
}

// KeyAuthors returns array of all authors keys
func (w Work) KeyAuthors() []string {
	a := make([]string, len(w.AuthorsKey))
	for i, AuthorKey := range w.AuthorsKey {
		a[i] = AuthorKey.AuthorKey.Key[9:]
	}
	return a
}

// Authors returns more information about the authors (using AuthorsKey)
func (w Work) Authors() (a []Author, err error) {
	return Authors(w)
}

// Editions returns an array of books linked to the work
func (w *Work) Editions() ([]Book, error) {
	editions := struct {
		Entries []Book `json:"entries"`
		Number  int    `json:"size"`
		Error   string `json:"error"`
	}{}

	s := fmt.Sprintf("https://openlibrary.org%s/editions.json", w.Key)
	resp, err := http.Get(s)
	if err != nil {
		return editions.Entries, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &editions)
	if editions.Error == "notfound" {
		return editions.Entries, errors.New("Editions of work not found")
	}
	// Populate the NumberOfEditions
	w.NumberOfEditions = editions.Number

	return editions.Entries, err
}
