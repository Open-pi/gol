package gol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// SearchData holds all information about Search coming from openlibrary
type SearchData struct {
	NumFound int         `json:"num_found"`
	Start    int         `json:"start"`
	Docs     []SearchDoc `json:"docs"`
	Error    string      `json:"error"`
}

// SearchDoc holds all information about Search Docs
type SearchDoc struct {
	Key                               string   `json:"key"`
	Type                              string   `json:"type"`
	Title                             string   `json:"title"`
	TitleSuggest                      string   `json:"title_suggest"`
	HasFulltext                       bool     `json:"has_fulltext"`
	EditionCount                      int      `json:"edition_count"`
	FirstPublishYear                  int      `json:"first_publish_year"`
	LastModifiedI                     int      `json:"last_modified_i"`
	EbookCountI                       int      `json:"ebook_count_i"`
	PublicScanB                       bool     `json:"public_scan_b,omitempty"`
	IaCollectionS                     string   `json:"ia_collection_s,omitempty"`
	LendingEditionS                   string   `json:"lending_edition_s,omitempty"`
	LendingIdentifierS                string   `json:"lending_identifier_s,omitempty"`
	PrintdisabledS                    string   `json:"printdisabled_s,omitempty"`
	CoverEditionKey                   string   `json:"cover_edition_key,omitempty"`
	CoverI                            int      `json:"cover_i,omitempty"`
	PublishYear                       []int    `json:"publish_year"`
	AuthorName                        []string `json:"author_name"`
	IDAmazon                          []string `json:"id_amazon,omitempty"`
	Seed                              []string `json:"seed"`
	AuthorAlternativeName             []string `json:"author_alternative_name"`
	Subject                           []string `json:"subject,omitempty"`
	Isbn                              []string `json:"isbn,omitempty"`
	IaLoadedID                        []string `json:"ia_loaded_id,omitempty"`
	EditionKey                        []string `json:"edition_key"`
	Language                          []string `json:"language,omitempty"`
	IDLibrarything                    []string `json:"id_librarything,omitempty"`
	Lcc                               []string `json:"lcc,omitempty"`
	Lccn                              []string `json:"lccn,omitempty"`
	IDGoodreads                       []string `json:"id_goodreads,omitempty"`
	PublishPlace                      []string `json:"publish_place,omitempty"`
	Contributor                       []string `json:"contributor,omitempty"`
	IDGoogle                          []string `json:"id_google,omitempty"`
	Ia                                []string `json:"ia,omitempty"`
	Text                              []string `json:"text"`
	Place                             []string `json:"place,omitempty"`
	Ddc                               []string `json:"ddc,omitempty"`
	AuthorKey                         []string `json:"author_key"`
	IDLibris                          []string `json:"id_libris,omitempty"`
	IDOverdrive                       []string `json:"id_overdrive,omitempty"`
	IDDepSitoLegal                    []string `json:"id_depósito_legal,omitempty"`
	IDDnb                             []string `json:"id_dnb,omitempty"`
	IDAlibrisID                       []string `json:"id_alibris_id,omitempty"`
	IaBoxID                           []string `json:"ia_box_id,omitempty"`
	FirstSentence                     []string `json:"first_sentence,omitempty"`
	Person                            []string `json:"person,omitempty"`
	IDWikidata                        []string `json:"id_wikidata,omitempty"`
	Oclc                              []string `json:"oclc,omitempty"`
	Publisher                         []string `json:"publisher,omitempty"`
	IDBcid                            []string `json:"id_bcid,omitempty"`
	PublishDate                       []string `json:"publish_date"`
	IDAmazonDeAsin                    []string `json:"id_amazon_de_asin,omitempty"`
	IDAmazonItAsin                    []string `json:"id_amazon_it_asin,omitempty"`
	IDNla                             []string `json:"id_nla,omitempty"`
	IDBritishNationalBibliography     []string `json:"id_british_national_bibliography,omitempty"`
	IDAmazonCoUkAsin                  []string `json:"id_amazon_co_uk_asin,omitempty"`
	Time                              []string `json:"time,omitempty"`
	IDAmazonCaAsin                    []string `json:"id_amazon_ca_asin,omitempty"`
	IDPaperbackSwap                   []string `json:"id_paperback_swap,omitempty"`
	IDBibliothQueNationaleDeFranceBnf []string `json:"id_bibliothèque_nationale_de_france_bnf,omitempty"`
	IDBritishLibrary                  []string `json:"id_british_library,omitempty"`
	IDScribd                          []string `json:"id_scribd,omitempty"`
	IDHathiTrust                      []string `json:"id_hathi_trust,omitempty"`
	IDCanadianNationalLibraryArchive  []string `json:"id_canadian_national_library_archive,omitempty"`
	Subtitle                          string   `json:"subtitle,omitempty"`
}

// SearchURL holds information about the Search URL
type SearchURL struct {
	url     string
	all     string
	title   string
	author  string
	subject string
}

// Search returns search API data from a constructed url
func Search(s string) (sd SearchData, err error) {
	resp, err := http.Get(s)
	if err != nil {
		return sd, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &sd)
	return
}

// Url returns OpenLibrary's search url
func SearchUrl() SearchURL {
	return SearchURL{url: "https://openlibrary.org/search.json?"}
}

// All allows to search in everything
func (s SearchURL) All(q string) SearchURL {
	q = strings.ReplaceAll(q, " ", "+")
	s.all = fmt.Sprintf("&q=%s", q)
	return s
}

// Title allows to only search in titles
func (s SearchURL) Title(t string) SearchURL {
	t = strings.ReplaceAll(t, " ", "+")
	s.title = fmt.Sprintf("&title=%s", t)
	return s
}

// Author allows to only search in authors
func (s SearchURL) Author(a string) SearchURL {
	a = strings.ReplaceAll(a, " ", "+")
	s.author = fmt.Sprintf("&author=%s", a)
	return s
}

// Subject allow to search in a specific subject
func (s SearchURL) Subject(sbj string) SearchURL {
	sbj = strings.ReplaceAll(sbj, " ", "+")
	s.subject = fmt.Sprintf("&subject=%s", sbj)
	return s
}

// Construct contructs and returns a complete searchable url
func (su SearchURL) Construct() string {
	return fmt.Sprintf("%s%s%s%s%s", su.url, su.all, su.title, su.author, su.subject)
}
