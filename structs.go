package gol

type Time struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Author struct {
	Key string `json:"key"`
}

type Type struct {
	Key string `json:"key"`
}

type AuthorAndType struct {
	Type   Type
	Author Author
}

type Work struct {
	Created        Time            `json:"created"`
	Subjects       []string        `json:"subjects"`
	LatestRevision int             `json:"latest_revision"`
	Key            string          `json:"key"`
	Title          string          `json:"title"`
	Authors        []AuthorAndType `json:"authors"`
	Type           Type            `json:"type"`
	LastModified   Time            `json:"last_modified"`
	Covers         []int           `json:"covers"`
	Revision       int             `json:"revision"`
	Error          string          `json:"error"`
}
