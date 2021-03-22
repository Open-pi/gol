package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

// GetBook returns a book from its olid
func GetBook(olid string) (b Book, err error) {
	s := fmt.Sprintf("https://openlibrary.org/books/%s.json", olid)
	resp, err := http.Get(s)
	if err != nil {
		return b, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &b)
	if b.Error == "notfound" {
		return b, errors.New("Book/Edition not found")
	}
	return
}

// GetCoverURL returns the url of the requested cover
// GetCoverURL accepts the cover type i.e b(book), a(author),
// the key of the book i.e ISBN, OCLC, LCCN, OLID and ID,
// the value of the key, and
// the size of the cover (S, M, or L)
func GetCoverURL(coverType string, key string, value string, size string) string {
	return fmt.Sprintf("http://covers.openlibrary.org/%s/%s/%s-%s.jpg", coverType, key, value, size)
}

func GetBookCoverURL(key string, value string, size string) string {
	return GetCoverURL("b", key, value, size)
}

func GetAuthorCoverURL(key string, value string, size string) string {
	return GetCoverURL("a", key, value, size)
}
