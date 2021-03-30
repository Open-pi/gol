/*
Package gol implements an easy interface to make calls to the OpenLibrary API

gol uses the WorkAPI, the EditionAPI, and the CoverAPI
*/
package gol

type Time struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AuthorKey struct {
	Key string `json:"key"`
}

type Type struct {
	Key string `json:"key"`
}

type AuthorKeyAndType struct {
	Type      Type      `json:"type"`
	AuthorKey AuthorKey `json:"Author"`
}

type TypeValue struct {
	Type  string `json:"type"`
	Value string `josn:"value"`
}
