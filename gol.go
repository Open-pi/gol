/*
Package gol implements an easy interface to make calls to the OpenLibrary API

gol uses the WorkAPI, the EditionAPI, and the CoverAPI
*/
package gol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

type Container = *gabs.Container

func HasError(data Container) error {
	// verify if an error field is present in the returned data
	if err, ok := data.Path("error").Data().(string); ok {
		return fmt.Errorf("Error fetching data; %s", err)
	}
	return nil
}

func Request(url string) (Container, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("404 Error")
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	dec := json.NewDecoder(bytes.NewReader(bodyBytes))
	dec.UseNumber()
	container, err := gabs.ParseJSONDecoder(dec)
	if err != nil {
		return nil, err
	}
	return container, nil
}

func MakeRequest(api string, id string, params ...string) (Container, error) {
	var s string
	if len(params) == 0 {
		s = fmt.Sprintf("https://openlibrary.org/%s/%s.json", api, id)
	} else {
		s = fmt.Sprintf("https://openlibrary.org/%s/%s.json?%s", api, id, strings.Join(params, "&"))
	}
	return Request(s)
}

func MakeBookRequest(id string) (Container, error) {
	return MakeRequest("books", id)
}

func MakeAuthorRequest(id string) (Container, error) {
	return MakeRequest("authors", id)
}

func MakeISBNRequest(isbn string) (Container, error) {
	return MakeRequest("isbn", isbn)
}

func MakeSubjectRequest(subject string) (Container, error) {
	return MakeRequest("subjects", subject)
}

func MakeDetailedSubjectRequest(subject string) (Container, error) {
	return MakeRequest("subjects", subject, "details=true")
}

func MakeWorkRequest(work string) (Container, error) {
	return MakeRequest("works", work)
}
