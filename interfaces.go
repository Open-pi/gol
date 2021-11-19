package gol

/*
type IKeyAuthors interface {
	KeyAuthors() []string
}
*/

type HasCovers interface {
	KeyCovers() ([]string, error)
	FirstCoverKey() string
}

/*
func Authors(i IKeyAuthors) (a []Author, err error) {
	for _, key := range i.KeyAuthors() {
		author, err := GetAuthor(key)
		if err != nil {
			return a, err
		}
		a = append(a, author)
	}
	return
}
*/

func Cover(c HasCovers, size string) string {
	if id := c.FirstCoverKey(); id != "" {
		return GetBookCoverURL("id", id, size)
	}
	return ""
}
