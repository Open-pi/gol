package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
)

func TestGetCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg"
	result := gol.GetCoverURL("b", "olid", "OL4554174M", "S")
	if result != cover {
		t.Error("URL returned is not correct")
	}
}

func TestGetBookCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg"
	result := gol.GetBookCoverURL("olid", "OL4554174M", "S")
	if result != cover {
		t.Error("URL returned is not correct")
	}
}

func TestGetAuthorCoverURL(t *testing.T) {
	cover := "http://covers.openlibrary.org/a/olid/OL229501A-S.jpg"
	result := gol.GetAuthorCoverURL("olid", "OL229501A", "S")
	if result != cover {
		t.Error("URL returned is not correct")
	}
}
