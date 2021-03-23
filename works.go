package gol

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Work struct {
	Created        Time            `json:"created"`
	Subjects       []string        `json:"subjects"`
	LatestRevision int             `json:"latest_revision"`
	Key            string          `json:"key"`
	Title          string          `json:"title"`
	Authors        []AuthorAndType `json:"authors"`
	Type           Type            `json:"type"`
	LastModified   Time            `json:"last_modified"`
	Covers         []int           `json:"covers"`
	Revision       int             `json:"revision"`
	Error          string          `json:"error"`
}

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
