package gol

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Author struct {
	Container
	keyCovers []string
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
	return
}

/*
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
}*/

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

/*
// Cover returns (if it exists) the URL of the Author's Phote/Cover
func (a Author) Cover(size string) string {
	return Cover(a, size)
}
*/
