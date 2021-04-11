package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

var b = gol.Book{
	Publishers:        []string{"Oxford University Press"},
	Identifiers:       gol.Identifiers{Librarything: []string{"23538"}, Goodreads: []string{"3109896"}},
	IaBoxID:           []string{"IA144904"},
	Covers:            []int{7891217},
	LocalID:           []string{"urn:cst:10017032837"},
	IaLoadedID:        []string{"selfishgene00dawk"},
	LcClassifications: []string{"QH437 .D38 1978"},
	Key:               "/books/OL4554174M",
	AuthorsKey:        []gol.AuthorKey{{"/authors/OL236174A"}},
	Ocaid:             "selfishgene00dawk",
	PublishPlaces:     []string{"New York"},
	Subjects:          []string{"Genetics.", "Evolution (Biology)"},
	Pagination:        "xi, 224 p. ;",
	SourceRecords:     []string{"ia:selfishgene00dawk", "marc:marc_claremont_school_theology/CSTMARC1_barcode.mrc:83476818:2271", "marc:marc_loc_2016/BooksAll.2016.part10.utf8:115807776:727", "marc:marc_claremont_school_theology/CSTMARC1_multibarcode.mrc:83581761:2271"},
	Title:             "The selfish gene",
	DeweyDecimalClass: []string{"591.5"},
	Notes:             gol.Notes{"/type/text", "Bibliography: p. [217]-220.\nIncludes index."},
	NumberOfPages:     224,
	Languages:         []gol.Language{{"/languages/eng"}},
	Lccn:              []string{"77023844"},
	Isbn10:            []string{"0195200004"},
	PublishDate:       "1978",
	PublishCountry:    "enk",
	ByStatement:       "Richard Dawkins.",
	OclcNumbers:       []string{"3167790"},
	Works:             []gol.Works{{"/works/OL1966513W"}},
	Type:              gol.Type{"/type/edition"},
	LatestRevision:    13,
	Revision:          13,
	Created:           gol.Time{Type: "/type/datetime", Value: "2008-04-01T03:28:50.625462"},
	LastModified:      gol.Time{Type: "/type/datetime", Value: "2021-03-03T05:21:06.382367"},
}

func TestGetEdition(t *testing.T) {
	tr, err := gol.GetEdition("OL4554174M")

	if err != nil {
		t.Errorf("got unexpected error: %v", err)
	}
	if !cmp.Equal(b, tr) {
		t.Error("Incorrect testresult GetEdition(OL4554174M)")
	}

	tr, err = gol.GetEdition("OL4554174")
	if err == nil {
		t.Error("GetEdition did not return an err when calling an inexistent book")
	}
}

func TestGetEditionISBN(t *testing.T) {
	t.Parallel()

	tt1 := []struct {
		name  string
		input string
		tr    gol.Book
	}{
		{"Correct ISBN-10 ID prefix", "0195200004", b},
		{"Correct ISBN-13 ID prefix", "9780140328721", b13},
		{"Correct ISBN-13 Dashed Format", "978-0-14-032872-1", b13},
	}

	for _, tc := range tt1 {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tr, err := gol.GetEditionISBN(tc.input)
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !cmp.Equal(tc.tr, tr) {
				t.Fatalf("Unexpected result : GetEditionISBN(%s)", tc.input)
			}
		})
	}

	tt2 := []struct {
		name  string
		input string
	}{
		{"Incorrect ISBN ID Length", "9984"},
		{"Incorrect ISBN-13 ID prefix", "9870140328721"},
		{"Inexistent ISBN/Book", "9780140328725"},
	}

	for _, tc := range tt2 {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := gol.GetEditionISBN(tc.input)
			if err == nil {
				t.Fatalf("GetEditionISBN(%s) did not return an error for incorrect/inexistent ISBN", tc.input)
			}
		})
	}
}

func TestEditionKeyAuthors(t *testing.T) {
	a := b.KeyAuthors()
	if !cmp.Equal(a, []string{"OL236174A"}) {
		t.Errorf("Unexpected returned array. Expecting [OL18295A] got %v", a)
	}
}

func TestEditionAuthors(t *testing.T) {
	tr, err := editions[0].Authors()
	if err != nil {
		t.Errorf("b.Authors() returned an error, %v expecting an Authors slice", err)
	}
	if !cmp.Equal(tr, authors) {
		t.Errorf("Expected set of authors incorrect")
	}
}
