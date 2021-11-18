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

//func TestEditionKeyAuthors(t *testing.T) {
//a := b.KeyAuthors()
//if !cmp.Equal(a, []string{"OL236174A"}) {
//t.Errorf("Unexpected returned array. Expecting [OL18295A] got %v", a)
//}
//}

//func TestEditionAuthors(t *testing.T) {
//tr, err := editions[0].Authors()
//if err != nil {
//t.Errorf("b.Authors() returned an error, %v expecting an Authors slice", err)
//}
//if !cmp.Equal(tr, authors) {
//t.Errorf("Expected set of authors incorrect")
//}
//}
