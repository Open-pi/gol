package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Author struct {
	Bio            Bio       `json:"bio"`
	Name           string    `json:"name"`
	Title          string    `json:"title"`
	PersonalName   string    `json:"personal_name"`
	Wikipedia      string    `json:"wikipedia"`
	Created        Time      `json:"created"`
	Photos         []int     `json:"photos"`
	LastModified   Time      `json:"last_modified"`
	LatestRevision int       `json:"latest_revision"`
	Key            string    `json:"key"`
	BirthDate      string    `json:"birth_date"`
	Revision       int       `json:"revision"`
	Type           Type      `json:"type"`
	RemoteIds      RemoteIds `json:"remote_ids"`
	Error          string    `json:"error"`
}

type Bio struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RemoteIds struct {
	Viaf     string `json:"viaf"`
	Wikidata string `json:"wikidata"`
	Isni     string `json:"isni"`
}

// GetAuthor returns an Author struct
func GetAuthor(id string) (a Author, e error) {
	s := fmt.Sprintf("https://openlibrary.org/authors/%s.json", id)
	resp, err := http.Get(s)
	if err != nil {
		return a, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &a)
	if a.Error == "notfound" {
		return a, errors.New("Author not found")
	}
	return
}

// works returns all the works of the author
func (a Author) Works() ([]Work, error) {
	return a.works("")
}

func (a Author) works(offset string) ([]Work, error) {
	var s string
	works := struct {
		Entries []Work `json:"entries"`
		Number  int    `json:"size"`
		Links   struct {
			Next string `json:"next"`
		}
		Error string `json:"error"`
	}{}

	if offset != "" {
		s = fmt.Sprintf("https://openlibrary.org%s", offset)
	} else {
		s = fmt.Sprintf("https://openlibrary.org%s/works.json", a.Key)
	}

	resp, err := http.Get(s)
	if err != nil {
		return works.Entries, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &works)
	if works.Error == "notfound" {
		return works.Entries, fmt.Errorf("Works of  %s not found", a.Key)
	}

	// Use the next field If there are still another works to request from the API
	if works.Links.Next != "" {
		entries, err := a.works(works.Links.Next)
		if err != nil {
			return works.Entries, err
		}
		works.Entries = append(works.Entries, entries...)
	}

	return works.Entries, err
}

// KeyCover returns (if it exists) the URL of the Author's Photo/Cover
func (a Author) KeyCover() string {
	if len(a.Photos) > 0 {
		return strconv.Itoa(a.Photos[0])
	}
	return ""
}

// Cover returns (if it exists) the URL of the Author's Phote/Cover
func (a Author) Cover(size string) string {
	return Cover(a, size)
}
