package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
)

func TestGetCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg"
	tr := gol.GetCoverURL("b", "olid", "OL4554174M", "S")
	if tr != cover {
		t.Error("URL returned is not correct")
	}
}

func TestGetBookCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg"
	tr := gol.GetBookCoverURL("olid", "OL4554174M", "S")
	if tr != cover {
		t.Error("URL returned is not correct")
	}
}

func TestGetAuthorCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/a/olid/OL229501A-S.jpg"
	tr := gol.GetAuthorCoverURL("olid", "OL229501A", "S")
	if tr != cover {
		t.Error("URL returned is not correct")
	}
}
