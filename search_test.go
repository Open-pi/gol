package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	// TODO: Add search by Text, Lists, and Avcanced search
	tt := []struct {
		name string
		input string
		tr gol.SearchData
	}{
		{"Test Simple Search", gol.Url().All("the selfish gene s").Construct(), searchQ},
		{"Test Title Search", gol.Url().Title("spellslinger 6").Construct(), searchTitle},
		{"Test Author Search", gol.Url().Author("Sarah Penner").Construct(), searchAuthor},
		{"Test Subject Search", gol.Url().Subject("abcd").Construct(), searchSubject},
		{"Test Mixed Search", gol.Url().Author("Richard Dawkins").Title("The Selfish Gene").Subject("evolution").Construct(), searchMixed},
	}

	for _, tc := range tt {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tr, err := gol.Search(tc.input)

			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			if !cmp.Equal(tc.tr, tr) {
				t.Fatalf("Unexpected result for Search()")
			} 
		})
	}
}
