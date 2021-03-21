package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
