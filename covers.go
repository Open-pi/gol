package gol

import "fmt"

// GetCoverURL returns the url of the requested cover
// GetCoverURL accepts the cover type i.e b(book), a(author),
// the key of the book i.e ISBN, OCLC, LCCN, OLID and ID,
// the value of the key, and
// the size of the cover (S, M, or L)
func GetCoverURL(coverType string, key string, value string, size string) string {
	return fmt.Sprintf("http://covers.openlibrary.org/%s/%s/%s-%s.jpg", coverType, key, value, size)
}

// GetBookCoverURL returns the url to book cover from an identifier.
// The size can be specified (S, M, or L)
func GetBookCoverURL(key string, value string, size string) string {
	return GetCoverURL("b", key, value, size)
}

// GetAuthorCoverURL returns the url of author cover (picture) from an identifier.
// The size can be specified (S, M, or L)
func GetAuthorCoverURL(key string, value string, size string) string {
	return GetCoverURL("a", key, value, size)
}
