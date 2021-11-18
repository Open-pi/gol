package gol_test

/*
import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	// TODO: Add search by Text, Lists, and Avcanced search
	tt := []struct {
		name  string
		input string
		tr    gol.SearchData
	}{
		{"Test Simple Search", gol.SearchUrl().All("the selfish gene s").Construct(), searchQ},
		{"Test Title Search", gol.SearchUrl().Title("spellslinger 6").Construct(), searchTitle},
		{"Test Author Search", gol.SearchUrl().Author("Sarah Penner").Construct(), searchAuthor},
		{"Test Subject Search", gol.SearchUrl().Subject("abcd").Construct(), searchSubject},
		{"Test Mixed Search", gol.SearchUrl().Author("Richard Dawkins").Title("The Selfish Gene").Subject("evolution").Construct(), searchMixed}}

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

func TestSearchURL(t *testing.T) {
	tr := gol.SearchUrl().All("the selfish gene s").Construct()
	if tr != searchQURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	tr = gol.SearchUrl().Title("spellslinger 6").Construct()
	if tr != searchTitleURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	tr = gol.SearchUrl().Author("Sarah Penner").Construct()
	if tr != searchAuthorURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	tr = gol.SearchUrl().Subject("abcd").Construct()
	if tr != searchSubjectURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	tr = gol.SearchUrl().Author("Richard Dawkins").Title("The Selfish Gene").Subject("evolution").Construct()
	if tr != searchMixedURL {
		t.Errorf("Incorrect URL construction of Search query")
	}
}
*/
