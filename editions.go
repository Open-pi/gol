package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Book map[string]interface{}

// GetEdition returns a book from its open library id
func GetEdition(olid string) (b Book, err error) {
	s := fmt.Sprintf("https://openlibrary.org/books/%s.json", olid)
	resp, err := http.Get(s)
	if err != nil {
		return b, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &b)

	if error, ok := b["error"]; ok {
		return b, fmt.Errorf("GetEdition: Error fetching book; %s", error)
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
	if error, ok := b["error"]; ok {
		return b, fmt.Errorf("Book/Edition not found; %s", error)
	}
	return
}

/*
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
*/
