package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
