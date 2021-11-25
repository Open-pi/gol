package gol

import (
	"fmt"
)

type QueryURL struct {
	url       string
	querytype string
	author    string
	limit     string
	work      string
	title     string
}

// QueryURL returns QueryURL type that holds OpenLibrary Query URL
func QueryUrl() QueryURL {
	return QueryURL{url: "https://openlibrary.org/query.json?"}
}

// Construct returns the string represtation of the QueryURL
func (q QueryURL) Construct() string {
	return fmt.Sprintf("%s%s%s%s%s%s", q.url, q.querytype, q.author, q.limit, q.work, q.title)
}

// Type returns a QueryURL with the type that has be passed as an arg
func (q QueryURL) Type(s string) QueryURL {
	q.querytype = fmt.Sprintf("&type=/type/%s", s)
	return q
}

// Author returns a QueryURL with the author that has be passed as an arg
func (q QueryURL) Author(author string) QueryURL {
	q.author = fmt.Sprintf("&authors=/authors/%s", author)
	return q
}

// Limit sets the number limit of the Query
func (q QueryURL) Limit(l int) QueryURL {
	q.limit = fmt.Sprintf("&limit=%d", l)
	return q
}

// Work sets the work key of the Query
func (q QueryURL) Work(w string) QueryURL {
	q.work = fmt.Sprintf("&works=/works/%s", w)
	return q
}

// Title sets the title key of the query
func (q QueryURL) Title(t string) QueryURL {
	q.title = fmt.Sprintf("&title=%s", t)
	return q
}

// Query returns the result of the query from a url -- Constructed by Construct()
func Query(url string) (Container, error) {
	result, err := Request(url)
	if err != nil {
		return nil, err
	}
	if err := HasError(result); err != nil {
		return nil, err
	}
	return result, nil
}
