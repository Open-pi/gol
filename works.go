package gol

import (
	"encoding/json"
	"fmt"
)

type Work struct {
	Container
	subjects   []string
	key        string
	title      string
	desc       string
	keyAuthors []string
	keyCovers  []string
}

// LoadWork parses the json container and fills all the fields
func (w *Work) Load() {
	w.Subjects()
	w.Title()
	w.Desc()
	w.KeyAuthors()
	w.KeyCovers()
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
	for _, child := range w.S("author").Children() {
		for _, v := range child.ChildrenMap() {
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

/*
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
*/
