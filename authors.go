package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type Author struct {
	Container
	keyCovers []string
	Key       string
}

// GetAuthor returns an Author struct
func GetAuthor(id string) (a Author, err error) {
	a.Container, err = MakeAuthorRequest(id)
	if err != nil {
		return a, err
	}

	if HasError(a.Container) != nil {
		return a, errors.New("Author not found")
	}
	a.Key = id
	return
}

// works returns all the works of the author
func (a Author) Works() ([]Work, error) {
	return a.works("")
}

func (a Author) works(offset string) ([]Work, error) {
	var s string

	if offset != "" {
		s = fmt.Sprintf("https://openlibrary.org%s", offset)
	} else {
		s = fmt.Sprintf("https://openlibrary.org/authors/%s/works.json", a.Key)
	}

	container, err := Request(s)
	if err != nil {
		return nil, err
	}

	if HasError(container) != nil {
		return nil, errors.New("Author not found")
	}

	entries := []Work{}
	var wg sync.WaitGroup
	entriesJSON := container.Path("entries").Children()
	wg.Add(len(entriesJSON))

	for _, child := range entriesJSON {
		work := Work{
			Container: child,
		}
		go func(wg *sync.WaitGroup) {
			work.Load()
			wg.Done()
		}(&wg)

		entries = append(entries, work)
	}

	// Use the next field If there are still another works to request from the API
	if next, ok := container.Path("links.next").Data().(string); ok && next != "" {
		nextEntries, err := a.works(next)
		if err != nil {
			return entries, err
		}
		entries = append(entries, nextEntries...)
	}

	wg.Wait()
	return entries, err
}

// KeyCovers returns (if they exists) the key covers/photo of the author
func (a Author) KeyCovers() ([]string, error) {
	if len(a.keyCovers) > 0 {
		return a.keyCovers, nil
	}

	for _, child := range a.S("photos").Children() {
		id, err := child.Data().(json.Number).Int64()
		if err == nil {
			a.keyCovers = append(a.keyCovers, fmt.Sprintf("%v", id))
		}
	}

	if len(a.keyCovers) == 0 {
		return a.keyCovers, fmt.Errorf("Could not find the key covers")
	}
	return a.keyCovers, nil
}

func (a Author) FirstCoverKey() string {
	if keys, ok := a.KeyCovers(); ok == nil {
		return keys[0]
	}
	return ""
}

// Cover returns (if it exists) the URL of the first Author's Photo/Cover
func (a Author) Cover(size string) string {
	return Cover(a, size)
}
