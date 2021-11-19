/*
Package gol implements an easy interface to make calls to the OpenLibrary API

gol uses the WorkAPI, the EditionAPI, and the CoverAPI
*/
package gol

import (
	"fmt"

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
