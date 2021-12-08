package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
)

func TestGetEdition(t *testing.T) {
	_, err := gol.GetEdition("OL0000000")
	if err == nil {
		t.Error("GetEdition did not return an err when calling an inexistent book")
	}
}

func TestGetEditionISBN(t *testing.T) {
	t.Parallel()

	tt1 := []struct {
		name  string
		input string
	}{
		{"Incorrect ISBN ID Length", "9984"},
		{"Incorrect ISBN-13 ID prefix", "9870140328721"},
		{"Inexistent ISBN/Book", "9780140328725"},
	}

	for _, tc := range tt1 {
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
	t.Parallel()
	t.Run("Existing Author", func(t *testing.T) {
		t.Parallel()
		b, _ := gol.GetEdition("OL4554174M")
		_, ok := b.KeyAuthors()
		if ok != nil {
			t.Fatalf("KeyAuthors() did not return authors for an existing book")
		}
	})

	t.Run("Non-Existing Author", func(t *testing.T) {
		t.Parallel()
		b, _ := gol.GetEdition("OL0000000M")
		_, ok := b.KeyAuthors()
		if ok == nil {
			t.Fatalf("KeyAuthors() did not return an error for non existing book")
		}
	})
}

func TestEditionRefresh(t *testing.T) {
}

func TestKeyCovers(t *testing.T) {
}

func TestEditionGoodReads(t *testing.T) {
}

func TestEditionISBN(t *testing.T) {
}

func TestEditionSubjects(t *testing.T) {
}

func TestEditionPublishers(t *testing.T) {
}

func TestEditionTitle(t *testing.T) {
}

func TestEditionWorkKeys(t *testing.T) {
}

func TestEditionNumberOfPages(t *testing.T) {
}

/*
//func TestEditionAuthors(t *testing.T) {
//tr, err := editions[0].Authors()
//if err != nil {
//t.Errorf("b.Authors() returned an error, %v expecting an Authors slice", err)
//}
//if !cmp.Equal(tr, authors) {
//t.Errorf("Expected set of authors incorrect")
//}
//}
*/
