package gol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Subject holds all the information about a particular subject
type Subject struct {
	Key               string          `json:"key"`
	Name              string          `json:"name"`
	SubjectType       string          `json:"subject_type"`
	WorkCount         int             `json:"work_count"`
	Works             []SubjectWork   `json:"works"`
	EbookCount        int             `json:"ebook_count"`
	Subjects          []SubjectDetail `json:"subjects"`
	Places            []SubjectDetail `json:"places"`
	People            []SubjectDetail `json:"people"`
	Times             []SubjectDetail `json:"times"`
	Authors           []SubjectDetail `json:"authors"`
	Publishers        []SubjectDetail `json:"publishers"`
	Languages         []SubjectDetail `json:"languages"`
	PublishingHistory [][]int         `json:"publishing_history"`
	Error             string          `json:"error"`
}

// SubjectWork holds all the information about works of a particular subject
type SubjectWork struct {
	Key               string                  `json:"key"`
	Title             string                  `json:"title"`
	EditionCount      int                     `json:"edition_count"`
	CoverID           int                     `json:"cover_id"`
	CoverEditionKey   string                  `json:"cover_edition_key"`
	Subject           []string                `json:"subject"`
	IaCollection      []string                `json:"ia_collection"`
	Lendinglibrary    bool                    `json:"lendinglibrary"`
	Printdisabled     bool                    `json:"printdisabled"`
	LendingEdition    string                  `json:"lending_edition"`
	LendingIdentifier string                  `json:"lending_identifier"`
	Authors           []SubjectDetail         `json:"authors"`
	FirstPublishYear  interface{}             `json:"first_publish_year"`
	Ia                string                  `json:"ia"`
	PublicScan        bool                    `json:"public_scan"`
	HasFulltext       bool                    `json:"has_fulltext"`
	CheckedOut        bool                    `json:"checked_out"`
	Availability      SubjectWorkAvailability `json:"availability"`
}

// SubjectDetail holds informations about subject details
type SubjectDetail struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// SubjectWorkAvailability holds information about the work of a subject
type SubjectWorkAvailability struct {
	Status              string      `json:"status"`
	AvailableToBrowse   bool        `json:"available_to_browse"`
	AvailableToBorrow   bool        `json:"available_to_borrow"`
	AvailableToWaitlist bool        `json:"available_to_waitlist"`
	IsPrintdisabled     bool        `json:"is_printdisabled"`
	IsReadable          bool        `json:"is_readable"`
	IsLendable          bool        `json:"is_lendable"`
	Identifier          string      `json:"identifier"`
	Isbn                interface{} `json:"isbn"`
	Oclc                interface{} `json:"oclc"`
	OpenlibraryWork     string      `json:"openlibrary_work"`
	OpenlibraryEdition  string      `json:"openlibrary_edition"`
	LastLoanDate        interface{} `json:"last_loan_date"`
	NumWaitlist         interface{} `json:"num_waitlist"`
	LastWaitlistDate    interface{} `json:"last_waitlist_date"`
	Collection          string      `json:"collection"`
	IsRestricted        bool        `json:"is_restricted"`
	IsBrowseable        bool        `json:"is_browseable"`
	Src                 string      `json:"__src__"`
}

// GetSubject returns the chosen subject's information
func GetSubject(subject string) (sbj Subject, err error) {
	return GeneralGetSubject(fmt.Sprintf("https://openlibrary.org/subjects/%s.json", subject))
}

// GetSubject returns the chosen subject's information with more details
func GetSubjectDetails(subject string) (sbj Subject, err error) {
	return GeneralGetSubject(fmt.Sprintf("https://openlibrary.org/subjects/%s.json?details=true", subject))
}

// GeneralGetSubject is a general function for GetSubject and GetSubjectDetails functions
func GeneralGetSubject(s string) (sbj Subject, err error) {
	resp, err := http.Get(s)
	if err != nil {
		return sbj, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &sbj)
	return
}
