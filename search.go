package gol

import (
	"fmt"
	"strings"
)

// SearchURL holds information about the Search URL
type SearchURL struct {
	url     string
	all     string
	title   string
	author  string
	subject string
}

// Search returns search API data from a constructed url
func Search(s string) (Container, error) {
	container, err := Request(s)
	if err != nil {
		return nil, err
	}

	if err := HasError(container); err != nil {
		return nil, err
	}
	return container, nil
}

// Url returns OpenLibrary's search url
func SearchUrl() SearchURL {
	return SearchURL{url: "https://openlibrary.org/search.json?"}
}

// All allows to search in everything
func (s SearchURL) All(q string) SearchURL {
	q = strings.ReplaceAll(q, " ", "+")
	s.all = fmt.Sprintf("&q=%s", q)
	return s
}

// Title allows to only search in titles
func (s SearchURL) Title(t string) SearchURL {
	t = strings.ReplaceAll(t, " ", "+")
	s.title = fmt.Sprintf("&title=%s", t)
	return s
}

// Author allows to only search in authors
func (s SearchURL) Author(a string) SearchURL {
	a = strings.ReplaceAll(a, " ", "+")
	s.author = fmt.Sprintf("&author=%s", a)
	return s
}

// Subject allow to search in a specific subject
func (s SearchURL) Subject(sbj string) SearchURL {
	sbj = strings.ReplaceAll(sbj, " ", "+")
	s.subject = fmt.Sprintf("&subject=%s", sbj)
	return s
}

// Construct constructs and returns a complete searchable url
func (su SearchURL) Construct() string {
	return fmt.Sprintf("%s%s%s%s%s", su.url, su.all, su.title, su.author, su.subject)
}
