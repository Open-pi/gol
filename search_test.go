package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	// TODO: Add search by Text, Lists, and Avcanced search
	t.Run("simple search", func(t *testing.T) {
		tr, err := gol.Search(gol.Url().All("the selfish gene s").Construct())
		if !cmp.Equal(searchQ, tr) || err != nil {
			t.Error("Incorrect testresult Search(the selfish gene s)")
		}
	})
	t.Run("title search", func(t *testing.T) {
		tr, err := gol.Search(gol.Url().Title("spellslinger 6").Construct())
		if !cmp.Equal(searchTitle, tr) || err != nil {
			t.Error("Incorrect testresult Search(title:spellslinger 6)")
		}
	})
	t.Run("author search", func(t *testing.T) {
		tr, err := gol.Search(gol.Url().Author("Sarah Penner").Construct())
		if !cmp.Equal(searchAuthor, tr) || err != nil {
			t.Error("Incorrect testresult Search(author:Sarah Penner)")
		}
	})
	t.Run("subject search", func(t *testing.T) {
		tr, err := gol.Search(gol.Url().Subject("abcd").Construct())
		if !cmp.Equal(searchSubject, tr) || err != nil {
			t.Error("Incorrect testresult Search(subject:abcd)")
		}
	})
	t.Run("mixed search", func(t *testing.T) {
		tr, err := gol.Search(gol.Url().Author("Richard Dawkins").Title("The Selfish Gene").Subject("evolution").Construct())
		if !cmp.Equal(searchMixed, tr) || err != nil {
			t.Error("Incorrect testresult Search(author:Richard Dawkins & title:The Selfish Gene & subject:evolution)")
		}
	})
}
