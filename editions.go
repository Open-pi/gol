package gol

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

type Book struct {
	Container
	Authors []string
}

// GetEdition returns a book from its open library id
func GetEdition(olid string) (b Book, err error) {
	s := fmt.Sprintf("https://openlibrary.org/books/%s.json", olid)
	resp, err := http.Get(s)
	if err != nil {
		return b, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	b.Container, err = gabs.ParseJSON(bodyBytes)
	if err != nil {
		return b, err
	}

	// verify if an error field is present in the returned data
	if err := HasError(b.Container); err != nil {
		return b, err
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
	b.Container, err = gabs.ParseJSON(bodyBytes)
	if err != nil {
		return b, err
	}
	if err != nil {
		return b, err
	}

	// verify if an error field is present in the returned data
	if err := HasError(b.Container); err != nil {
		return b, err
	}

	return
}

// KeyAuthors returns array of all authors keys
func (b *Book) KeyAuthors() (err error) {
	for _, child := range b.S("authors").Children() {
		for _, v := range child.ChildrenMap() {
			b.Authors = append(b.Authors, v.Data().(string))
		}
	}

	if len(b.Authors) == 0 {
		return fmt.Errorf("Could not find any authors")
	}
	return
}

/*
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
