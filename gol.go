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

//s := fmt.Sprintf("https://openlibrary.org/authors/%s.json", id)
//s := fmt.Sprintf("https://openlibrary.org/books/%s.json", olid)
func MakeRequest(api string, id string) (Container, error) {
	s := fmt.Sprintf("https://openlibrary.org/%s/%s.json", api, id)
	resp, err := http.Get(s)
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

func MakeBookRequest(id string) (Container, error) {
	return MakeRequest("books", id)
}

func MakeAuthorRequest(id string) (Container, error) {
	return MakeRequest("authors", id)
}

func MakeISBNRequest(isbn string) (Container, error) {
	return MakeRequest("isbn", isbn)
}
