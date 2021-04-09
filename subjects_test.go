package gol_test

import (
	"testing"

	"github.com/Open-pi/gol"
	"github.com/google/go-cmp/cmp"
)

func TestGetSubject(t *testing.T) {
	tr, err := gol.GetSubject("fake")
	if !cmp.Equal(sbj, tr) || err != nil {
		t.Error("Incorrect testresult GetSubject(fake)")
	}
}

func TestGetSubjectDetails(t *testing.T) {
	tr, err := gol.GetSubjectDetails("fake")
	if !cmp.Equal(sbjDetails, tr) || err != nil {
		t.Error("Incorrect testresult GetSubjectDetails(fake)")
	}
}
