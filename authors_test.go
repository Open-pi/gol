package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
)

func TestGetAuthor(t *testing.T) {
}

func TestAuthorKeyCovers(t *testing.T) {
	t.Parallel()

	t.Run("Valid Author Input", func(t *testing.T) {
		t.Parallel()
		a, _ := gol.GetAuthor("OL236174A")
		_, err := a.KeyCovers()
		if err != nil {
			t.Fatalf("Author.KeyCovers() did not return cover keys")
		}
	})

	t.Run("Invalid Author Input", func(t *testing.T) {
		a, _ := gol.GetAuthor("OL000000A")
		_, err := a.KeyCovers()
		if err == nil {
			t.Fatalf("Author.KeyCovers() no error returned for invalid input")
		}
	})
}

/*
func TestWorks(t *testing.T) {
	t.Parallel()
	tr, err := harris.Works()
	if err != nil {
		t.Errorf("Expecting []Works, got error %v", err)
	}
	if !cmp.Equal(tr, works) {
		t.Error("Returned Works are different from expected works")
	}
}
*/
