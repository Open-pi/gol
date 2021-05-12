package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestQuery(t *testing.T) {
	url := gol.QueryUrl().Type("edition").Author("OL236174A").Limit(2).Construct()
	tr, err := gol.Query(url)

	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}

	if !cmp.Equal(tr, query) {
		t.Errorf("incorrent result")
	}
}

func TestQueryURL(t *testing.T) {
	tt := []struct {
		name     string
		expected string
		output   string
	}{
		{"query editions by author with limit", gol.QueryUrl().Type("edition").Author("OL236174A").Limit(2).Construct(), "https://openlibrary.org/query.json?&type=/type/edition&authors=/authors/OL236174A&limit=2"},
		{"query editions by author and get title", gol.QueryUrl().Type("edition").Author("OL236174A").Title("").Construct(), "https://openlibrary.org/query.json?&type=/type/edition&authors=/authors/OL236174A&title="},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != tc.output {
				t.Fatalf("got unexpected result")
			}
		})
	}
}
