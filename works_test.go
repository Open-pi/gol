package gol_test

import (
	"testing"
)

func TestGetWork(t *testing.T) {
}

func TestWorkAuthors(t *testing.T) {
	//w, err := gol.GetWork("OL45583W")
	//if err != nil {
	//fmt.Println(err)
	//}
	//as, err := w.Authors()
	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(as)
}

func TestWork_Editions(t *testing.T) {
	//t.Parallel()
	//w, _ := gol.GetWork("OL257943W")
	//start := time.Now()
	//w.Editions()
	//fmt.Println(w.NumberOfEditions)
}

/*
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

*/
