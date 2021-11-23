package gol

type IKeyAuthors interface {
	KeyAuthors() ([]string, error)
}

type HasCovers interface {
	FirstCoverKey() string
}

func Authors(i IKeyAuthors) (as []Author, err error) {
	keys, err := i.KeyAuthors()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		author, err := GetAuthor(key[9:])
		if err != nil {
			continue
		}
		as = append(as, author)
	}
	return
}

func Cover(c HasCovers, size string) string {
	if id := c.FirstCoverKey(); id != "" {
		return GetBookCoverURL("id", id, size)
	}
	return ""
}
