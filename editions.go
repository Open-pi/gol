package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Book struct {
	Container
	title      string
	keyAuthors []string
	keyCovers  []string
	workKeys   []string
	goodreads  string
	isbn10     string
	subjects   []string
	publishers []string
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
	b.GoodReads()
	b.Isbn10()
}

// KeyAuthors returns array of all authors keys
func (b *Book) KeyAuthors() ([]string, error) {
	if len(b.keyAuthors) > 0 {
		return b.keyAuthors, nil
	}
	for _, child := range b.S("authors").Children() {
		for _, v := range child.ChildrenMap() {
			b.keyAuthors = append(b.keyAuthors, v.Data().(string))
		}
	}

	if len(b.keyAuthors) == 0 {
		return b.keyAuthors, fmt.Errorf("Could not find any authors")
	}
	return b.keyAuthors, nil
}

// Authors returns the authors of the book
func (b Book) Authors() ([]Author, error) {
	return Authors(&b)
}

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
		return b.keyCovers, fmt.Errorf("could not find key covers")
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
func (b Book) Cover(size string) string {
	return Cover(b, size)
}

// GoodReads returns the goodreads identifier
func (b Book) GoodReads() (string, error) {
	if b.goodreads != "" {
		return b.goodreads, nil
	} else {
		for _, child := range b.Path("identifiers.goodreads").Children() {
			b.goodreads = child.Data().(string)
			return b.goodreads, nil
		}
		return "", fmt.Errorf("could not find goodreads identifier")
	}
}

// Isbn10 returns the isbn10 identifier
func (b *Book) Isbn10() (string, error) {
	if b.isbn10 != "" {
		return b.goodreads, nil
	} else {
		for _, child := range b.Path("isbn_10").Children() {
			b.isbn10 = child.Data().(string)
			return b.isbn10, nil
		}
		return "", fmt.Errorf("could not find isbn10 identifier")
	}
}

// Subjects returns the subjects of the book
func (b *Book) Subjects() ([]string, error) {
	if len(b.subjects) > 0 {
		return b.subjects, nil
	} else {
		var ss []string
		for _, child := range b.Path("subjects").Children() {
			ss = append(ss, child.Data().(string))
		}
		if len(ss) > 0 {
			b.subjects = ss
			return b.subjects, nil
		} else {
			return nil, fmt.Errorf("could not find subjects")
		}
	}
}

// returns the publishers of the book
func (b *Book) Publishers() ([]string, error) {
	if len(b.publishers) > 0 {
		return b.publishers, nil
	} else {
		var ps []string
		for _, child := range b.Path("publishers").Children() {
			ps = append(ps, child.Data().(string))
		}
		if len(ps) > 0 {
			b.publishers = ps
			return b.publishers, nil
		} else {
			return nil, fmt.Errorf("could not find publishers")
		}
	}
}

func (b *Book) Title() (string, error) {
	if len(b.title) > 0 {
		return b.title, nil
	} else if title, ok := b.Path("title").Data().(string); ok {
		b.title = title
		return b.title, nil
	} else {
		return "", fmt.Errorf("could not find title")
	}
}

func (b *Book) WorkKeys() ([]string, error) {
	if len(b.workKeys) > 0 {
		return b.workKeys, nil
	} else {
		for _, child := range b.Path("works").Children() {
			for _, key := range child.ChildrenMap() {
				b.workKeys = append(b.workKeys, key.Data().(string))
			}
		}
		if len(b.workKeys) == 0 {
			return nil, fmt.Errorf("could not find work keys")
		} else {
			return b.workKeys, nil
		}
	}
}
