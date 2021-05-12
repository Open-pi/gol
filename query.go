package gol

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type QueryURL struct {
	url       string
	querytype string
	author    string
	limit     string
	work      string
}

// QueryURL returns QueryURL type that holds OpenLibrary Query URL
func QueryUrl() QueryURL {
	return QueryURL{url: "https://openlibrary.org/query.json?"}
}

// Construct returns the string represtation of the QueryURL
func (q QueryURL) Construct() string {
	return fmt.Sprintf("%s%s%s%s%s", q.url, q.querytype, q.author, q.limit, q.work)
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

// Query returns the result of the query from a url -- Constructed by Construct()
func Query(url string) (result map[string]interface{}, err error) {
	resp, err := http.Get(s)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	
	json.Unmarshal(bodyBytes, &result)
	return
}
