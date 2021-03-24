package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

var w gol.Work = gol.Work{
	Created:        gol.Time{Type: "/type/datetime", Value: "2009-10-15T11:23:34.130855"},
	Subjects:       []string{"History and criticism", "Russian literature", "Russian literature, history and criticism"},
	LatestRevision: 4,
	Key:            "/works/OL45583W",
	Title:          "An outline of Russian literature",
	Authors:        []gol.AuthorAndType{{gol.Type{"/type/author_role"}, gol.Author{"/authors/OL18295A"}}},
	Type:           gol.Type{"/type/work"},
	LastModified:   gol.Time{Type: "/type/datetime", Value: "2020-08-20T03:30:30.325116"},
	Covers:         []int{5917705},
	Revision:       4,
}

func TestGetWork(t *testing.T) {
	// Test GetWork when WorkId is valid
	result, err := gol.GetWork("OL45583W")
	if !cmp.Equal(w, result) || err != nil {
		t.Error("Incorrect result GetWork(OL45583W)")
	}

	// Test GetWork when workId is invalid
	result, err = gol.GetWork("notAnId")
	if err == nil {
		t.Error("GetWork did not return an err when calling an inexistent work")
	}
}

func TestCover(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		result string
	}{
		{"Test Small cover", "S", "http://covers.openlibrary.org/b/id/5917705-S.jpg"},
		{"Test Medium cover", "M", "http://covers.openlibrary.org/b/id/5917705-M.jpg"},
		{"Test Large cover", "L", "http://covers.openlibrary.org/b/id/5917705-L.jpg"},
	}

	// Test if returned covers are correct
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := w.Cover(tc.input)
			if c != tc.result {
				t.Fatalf("Cover returned is incorrect, %s instead of %s", c, tc.result)
			}
		})
	}
}
