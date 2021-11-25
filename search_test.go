package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
)

func TestSearch(t *testing.T) {
	//t.Parallel()
	//tr, err := gol.Search(gol.SearchUrl().Subject("abcd").Construct())
	//if err != nil {
	//log.Println(err)
	//}
	//log.Println(tr)
}

func TestSearchURL(t *testing.T) {
	tr := gol.SearchUrl().All("the selfish gene s").Construct()
	var searchQURL = "https://openlibrary.org/search.json?&q=the+selfish+gene+s"

	if tr != searchQURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	var searchTitleURL = "https://openlibrary.org/search.json?&title=spellslinger+6"
	tr = gol.SearchUrl().Title("spellslinger 6").Construct()
	if tr != searchTitleURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	var searchAuthorURL = "https://openlibrary.org/search.json?&author=Sarah+Penner"
	tr = gol.SearchUrl().Author("Sarah Penner").Construct()
	if tr != searchAuthorURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	var searchSubjectURL = "https://openlibrary.org/search.json?&subject=abcd"
	tr = gol.SearchUrl().Subject("abcd").Construct()
	if tr != searchSubjectURL {
		t.Errorf("Incorrect URL construction of Search query")
	}

	tr = gol.SearchUrl().Author("Richard Dawkins").Title("The Selfish Gene").Subject("evolution").Construct()
	var searchMixedURL = "https://openlibrary.org/search.json?&title=The+Selfish+Gene&author=Richard+Dawkins&subject=evolution"
	if tr != searchMixedURL {
		t.Errorf("Incorrect URL construction of Search query")
	}
}
