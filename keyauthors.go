package gol

type IKeyAuthors interface {
	KeyAuthors() []string
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
