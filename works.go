package gol

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Work struct {
	Container
	subjects         []string
	key              string
	title            string
	desc             string
	keyAuthors       []string
	keyCovers        []string
	NumberOfEditions int
	editions         []Book
}

// GetWork returns the work from the workID
// After making the request, the fields are loaded with Load
func GetWork(id string) (w Work, err error) {
	w.Container, err = MakeWorkRequest(id)
	if err != nil {
		return w, err
	}
	if err := HasError(w.Container); err != nil {
		return w, err
	}
	w.Load()

	return
}

// Load parses the json container and fills all the fields
func (w *Work) Load() {
	w.Key()
	w.Subjects()
	w.Title()
	w.Desc()
	w.KeyAuthors()
	w.KeyCovers()
}

func (w *Work) Key() (string, error) {
	if w.key != "" {
		return w.key, nil
	}

	if key, ok := w.Path("key").Data().(string); ok {
		w.key = key
		return w.key, nil
	}

	if w.key == "" {
		return w.key, fmt.Errorf("Key not found")
	}
	return w.key, nil
}

func (w *Work) Desc() (string, error) {
	if w.desc != "" {
		return w.desc, nil
	}

	if desc, ok := w.Path("description.value").Data().(string); ok {
		w.desc = desc
		return w.desc, nil
	} else {
		return "", fmt.Errorf("Description not found")
	}
}

func (w *Work) Subjects() ([]string, error) {
	if len(w.subjects) > 0 {
		return w.subjects, nil
	}

	for _, child := range w.S("subjects").Children() {
		w.subjects = append(w.subjects, child.Data().(string))
	}
	if len(w.subjects) == 0 {
		return nil, fmt.Errorf("subjects not found")
	}

	return w.subjects, nil
}

func (w *Work) Title() (string, error) {
	if w.title != "" {
		return w.title, nil
	}
	if title, ok := w.Path("title").Data().(string); ok {
		w.title = title
		return w.title, nil
	} else {
		return "", fmt.Errorf("Title not found")
	}
}

func (w *Work) KeyAuthors() ([]string, error) {
	if len(w.keyAuthors) > 0 {
		return w.keyAuthors, nil
	}
	for _, child := range w.Path("authors").Children() {
		for _, v := range child.S("author").ChildrenMap() {
			w.keyAuthors = append(w.keyAuthors, v.Data().(string))
		}
	}
	if len(w.keyAuthors) == 0 {
		return nil, fmt.Errorf("Key Authors not found")
	}

	return w.keyAuthors, nil
}

func (w *Work) KeyCovers() ([]string, error) {
	if len(w.keyCovers) > 0 {
		return w.keyCovers, nil
	}

	for _, child := range w.S("covers").Children() {
		id, err := child.Data().(json.Number).Int64()
		if err == nil {
			w.keyCovers = append(w.keyCovers, fmt.Sprintf("%v", id))
		}
	}
	if len(w.keyCovers) == 0 {
		return nil, fmt.Errorf("Key covers not found")
	}
	return w.keyCovers, nil
}

// FirstCoverKey return the first cover key (if it exists)
func (w Work) FirstCoverKey() string {
	if keys, ok := w.KeyCovers(); ok == nil {
		return keys[0]
	} else {
		return ""
	}
}

// Cover returns the cover url to the "first" edition.
// It takes size as an argument; it can be (S, M, or L)
func (w Work) Cover(size string) string {
	return Cover(w, size)
}

// Authors returns more information about the authors (using AuthorsKey)
func (w Work) Authors() (a []Author, err error) {
	return Authors(&w)
}

func (w *Work) Editions() ([]Book, error) {
	if len(w.editions) > 0 {
		return w.editions, nil
	}
	editions, err := w._Editions("")
	if err != nil {
		return nil, err
	}
	w.editions = editions
	return w.editions, nil
}

func (w *Work) _Editions(offset string) ([]Book, error) {
	var s string
	if offset != "" {
		s = fmt.Sprintf("https://openlibrary.org%s", offset)
	} else {
		key, err := w.Key()
		if err != nil {
			return nil, err
		}
		s = fmt.Sprintf("https://openlibrary.org%s/editions.json", key)
	}

	container, err := Request(s)
	if err != nil {
		return nil, err
	}

	if err := HasError(container); err != nil {
		return nil, err
	}

	size, _ := container.Path("size").Data().(json.Number).Int64()
	w.NumberOfEditions = int(size)

	entries := []Book{}
	var wg sync.WaitGroup
	entriesJSON := container.Path("entries").Children()
	wg.Add(len(entriesJSON))

	for _, edition := range entriesJSON {
		b := Book{
			Container: edition,
		}
		go func(wg *sync.WaitGroup) {
			b.Load()
			wg.Done()
		}(&wg)

		entries = append(entries, b)
	}

	// Use the next field If there are still another works to request from the API
	if next, ok := container.Path("links.next").Data().(string); ok && next != "" {
		nextEntries, err := w._Editions(next)
		if err != nil {
			return entries, err
		}
		entries = append(entries, nextEntries...)
	}

	wg.Wait()
	return entries, err
}
