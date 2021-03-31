package gol

type IKeyAuthors interface {
	KeyAuthors() []string
}

type Coverer interface {
	KeyCover() string
}

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

func Cover(c Coverer, size string) string {
	if id := c.KeyCover(); id != "" {
		return GetBookCoverURL("id", id, size)
	}
	return ""
}
