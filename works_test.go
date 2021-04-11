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
	AuthorsKey:     []gol.AuthorKeyAndType{{gol.Type{"/type/author_role"}, gol.AuthorKey{"/authors/OL18295A"}}},
	Type:           gol.Type{"/type/work"},
	LastModified:   gol.Time{Type: "/type/datetime", Value: "2020-08-20T03:30:30.325116"},
	Covers:         []int{5917705},
	Revision:       4,
}

var authors []gol.Author = []gol.Author{{
	Bio:            gol.Bio{Type: "/type/text", Value: "Maurice Baring (<a href=http://en.wikipedia.org/wiki/Maurice_Baring>Wikipedia</a>) was a versatile English man of letters, known as a dramatist, poet, novelist, translator and essayist, and also as a travel writer and war correspondent. He was the eighth child, and fifth son, of Edward Charles Baring, first Baron Revelstoke, of the Baring banking family, and his wife Louisa Emily Charlotte Bulteel, granddaughter of the second Earl Grey. He was educated at Eton College and Trinity College, Cambridge. After an abortive start on a diplomatic career, he travelled widely, particularly in Russia. He reported as an eye-witness on the Russo-Japanese War for the London Morning Post. At the start of World War I he joined the Royal Flying Corps, where he served as assistant to Trenchard. In 1918 Baring served as a staff officer in the Royal Air Force and was appointed OBE. In 1925 Baring received an honorary commission as a wing commander in the Reserve of Air Force Officers. After the war he enjoyed a period of success as a dramatist, and began to write novels. He suffered from chronic illness in the last years of his life; for the final 15 years of his life he was debilitated by Parkinson's Disease. He was widely connected socially, to some of the Cambridge Apostles, to The Coterie, and to the literary group around G. K. Chesterton and Hilaire Belloc in particular. He was staunch in his anti-intellectualism with respect to the arts, and a convinced practical joker. He became a Roman Catholic convert in 1909, being received into the church by Fr. Bowden at the Brompton Oratory. He described his conversion as \"the only action in my life which I am quite certain I have never regretted.\" Baring became a novelist late in his life, but went on to find success in that genre, as well as in playwriting. In France his novel Daphne Adene ran to over twenty printings."},
	Name:           "Maurice Baring",
	Title:          "comp.",
	Wikipedia:      "http://en.wikipedia.org/wiki/Maurice_Baring",
	Created:        gol.Time{Type: "/type/datetime", Value: "2008-04-01T03:28:50.625462"},
	Photos:         []int{6259227},
	LastModified:   gol.Time{Type: "/type/datetime", Value: "2020-09-30T13:46:23.364857"},
	LatestRevision: 9,
	Key:            "/authors/OL18295A",
	BirthDate:      "27 April 1874",
	PersonalName:   "Maurice Baring",
	Revision:       9,
	Type:           gol.Type{Key: "/type/author"},
	RemoteIds:      gol.RemoteIds{Viaf: "44386303", Wikidata: "Q1363814", Isni: "0000000108570698"},
}}

func TestGetWork(t *testing.T) {
	// Test GetWork when WorkId is valid
	tr, err := gol.GetWork("OL45583W")
	if err != nil {
		t.Errorf("got unexpected error: %v", err)
	}
	if !cmp.Equal(w, tr) {
		t.Error("Incorrect testresult GetWork(OL45583W)")
	}

	// Test GetWork when workId is invalid
	tr, err = gol.GetWork("notAnId")
	if err == nil {
		t.Error("GetWork did not return an err when calling an inexistent work")
	}
}

func TestCover(t *testing.T) {
	tt := []struct {
		name  string
		input string
		tr    string
	}{
		{"Test Small cover", "S", "http://covers.openlibrary.org/b/id/5917705-S.jpg"},
		{"Test Medium cover", "M", "http://covers.openlibrary.org/b/id/5917705-M.jpg"},
		{"Test Large cover", "L", "http://covers.openlibrary.org/b/id/5917705-L.jpg"},
	}

	// Test if returned covers are correct
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := w.Cover(tc.input)
			if c != tc.tr {
				t.Fatalf("Cover returned is incorrect, %s instead of %s", c, tc.tr)
			}
		})
	}
}

func TestWorkKeyAuthors(t *testing.T) {
	a := w.KeyAuthors()
	if !cmp.Equal(a, []string{"OL18295A"}) {
		t.Errorf("Unexpected returned array. Expecting [OL18295A] got %v", a)
	}
}

func TestWorkAuthors(t *testing.T) {
	tr, err := w.Authors()
	if err != nil {
		t.Errorf("w.Authors() returned an error, %v expecting an Author slice", err)
	} else if !cmp.Equal(tr, authors) {
		t.Errorf("Expected set of authors not correct")
	}
}

func TestWork_Editions(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name  string
		input gol.Work
		tr    []gol.Book
	}{
		{"Test Editions of same work", w, editions},
	}
	for _, tc := range tt {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tr, err := tc.input.Editions()
			if err != nil {
				t.Fatalf("%s returned an error %v", tc.name, err)
			}
			if !cmp.Equal(tr, tc.tr) {
				t.Fatalf("Unexpected result: test result different from testdata")
			}
		})
	}

	name := "Test Editions of inexistent work"
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		naw, _ := gol.GetWork("notAndId")
		_, err := naw.Editions()
		if err == nil {
			t.Fatalf("%s, should return error; got %v", name, err)
		}
	})
}
