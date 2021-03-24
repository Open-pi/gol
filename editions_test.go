package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestGetEdition(t *testing.T) {
	b := gol.Book{
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
		Languages:         []gol.Languages{{"/languages/eng"}},
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
	tr, err := gol.GetEdition("OL4554174M")
	if !cmp.Equal(b, tr) || err != nil {
		t.Error("Incorrect testresult GetEdition(OL4554174M)")
	}

	tr, err = gol.GetEdition("OL4554174")
	if err == nil {
		t.Error("GetEdition did not return an err when calling an inexistent book")
	}
}
