package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Book struct {
	Container
	keyAuthors []string
	keyCovers  []string
}

// GetEdition returns a book from its open library id
func GetEdition(olid string) (b Book, err error) {
	b.Container, err = MakeBookRequest(olid)
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

	b.Container, err = MakeISBNRequest(isbnid)
	if err != nil {
		return b, err
	}

	// verify if an error field is present in the returned data
	if err := HasError(b.Container); err != nil {
		return b, err
	}

	return
}

// Load tries to load the fields from the json container
func (b *Book) Load() {
	b.KeyAuthors()
	b.KeyCovers()
}

// KeyAuthors returns array of all authors keys
func (b *Book) KeyAuthors() (keys []string, err error) {
	if len(b.keyAuthors) > 0 {
		return b.keyAuthors, err
	}
	for _, child := range b.S("authors").Children() {
		for _, v := range child.ChildrenMap() {
			b.keyAuthors = append(b.keyAuthors, v.Data().(string))
		}
	}

	if len(b.keyAuthors) == 0 {
		return b.keyAuthors, fmt.Errorf("Could not find any authors")
	}
	return
}

//TODO: add this after dealing with authors API
// Authors returns all the information related to the book's authors
//func (b Book) Authors() (a []Author, err error) {
//return Authors(b)
//}

// KeyCover returns (if it exists) the ID of the work's cover
func (b *Book) KeyCovers() ([]string, error) {
	if len(b.keyCovers) > 0 {
		return b.keyCovers, nil
	}

	for _, child := range b.S("covers").Children() {
		id, err := child.Data().(json.Number).Int64()
		if err == nil {
			b.keyCovers = append(b.keyCovers, fmt.Sprintf("%v", id))
		}
	}

	if len(b.keyCovers) == 0 {
		return b.keyCovers, fmt.Errorf("Could not find key covers")
	}
	return b.keyCovers, nil
}

// FirstCoverKey returns the first cover if it exists
func (b Book) FirstCoverKey() string {
	if keys, ok := b.KeyCovers(); ok == nil {
		return keys[0]
	} else {
		return ""
	}
}

// Cover returns (if it exists) the URL of the Book's Cover
func (b Book) FirstURLCover(size string) string {
	return Cover(b, size)
}
