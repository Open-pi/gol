package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

var a gol.Author = gol.Author{
	Bio:            gol.Bio{Type: "/type/text", Value: "Clinton Richard Dawkins, FRS, FRSL is a British ethologist, evolutionary biologist and popular science author. He was formerly Professor for Public Understanding of Science at Oxford and was a fellow of New College, Oxford.\r\n\r\nDawkins came to prominence with his 1976 book *The Selfish Gene*, which popularised the gene-centred view of evolution and introduced the term meme. In 1982, he made a widely cited contribution to evolutionary biology with the concept, presented in his book *The Extended Phenotype*, that the phenotypic effects of a gene are not necessarily limited to an organism's body, but can stretch far into the environment, including the bodies of other organisms.\r\n\r\nDawkins is well known for his candid criticism of creationism and intelligent design. In his 1986 book *The Blind Watchmaker*, he argued against the watchmaker analogy, an argument for the existence of a supernatural creator based upon the complexity of living organisms. Instead, he described evolutionary processes as analogous to a blind watchmaker. He has since written several popular science books, and makes regular television and radio appearances, predominantly discussing these topics.\r\n\r\n([Source][1])\r\n\r\n\r\n  [1]: http://en.wikipedia.org/wiki/Richard_Dawkins"},
	Name:           "Richard Dawkins",
	Title:          "FRS, FRSL",
	PersonalName:   "Richard Dawkins",
	Wikipedia:      "http://en.wikipedia.org/wiki/Richard_Dawkins",
	Created:        gol.Time{Type: "/type/datetime", Value: "2008-04-01T03:28:50.625462"},
	Photos:         []int{6954866, -1},
	LastModified:   gol.Time{Type: "/type/datetime", Value: "2020-09-27T08:02:43.330024"},
	LatestRevision: 13,
	Key:            "/authors/OL236174A",
	BirthDate:      "26 March 1941",
	Revision:       13,
	Type:           gol.Type{"/type/author"},
	RemoteIds:      gol.RemoteIds{Viaf: "12307054", Wikidata: "Q44461", Isni: "0000000121208079"},
}

func TestGetAuthor(t *testing.T) {
	tt := []struct {
		name  string
		input string
		tr    gol.Author
	}{
		{"Test Author Richard Dawkins", "OL236174A", a},
		{"Test Author Sam Harris", "OL709883A", harris},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a, err := gol.GetAuthor(tc.input)
			if err != nil {
				t.Fatalf("%s returned error : %s when Author ID is correct", tc.name, err)
			}
			if !cmp.Equal(a, tc.tr) {
				t.Fatalf("Returned Author is different from expected author struct")
			}
		})
	}
}

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
