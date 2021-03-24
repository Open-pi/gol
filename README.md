# Gol
gol (**G**olang **O**pen **L**ibrary) is an interface for the OpenLibrary API. For more information about the API itself please visit [openlibrary.org](https://openlibrary.org/developers/api)

## Functions
These are the API functions (WorkAPI, EditionAPI, etc) to get the data. With them you can make other calls. For more information, browse the `docs` folder where you can find additional information.

| Functions | Args | Returns  |
|---|---|--|
| GetWork   | WorkId | (w Work, err error)  |
| GetEdition   | OLID | (b Book, err error)  |
| GetCoverURL   | coverType, identifierType, identifier, size | string  |
| GetBookCoverURL   | identifierType, identifier, size | string  |
| GetAuthorCoverURL   | identifierType, identifier, size | string  |

### WorkAPI Examples
```go
    work := gol.GetWork("OL45583W")
    // Output:
    // gol.Work{
    //  ...
    // }
```

### EditionAPI Examples
```go
    book := gol.GetEdition("OL4554174M")
    // Output:
    // gol.Book{
    //  Publishers:        []string{"Oxford University Press"},
    //  ...
    //	Key:               "/books/OL4554174M",
    //	Authors:           []gol.Author{{"/authors/OL236174A"}},
    //  ...
    //	Subjects:          []string{"Genetics.", "Evolution (Biology)"},
    //  ...
    //	Title:             "The selfish gene",
    //  ...
    //	NumberOfPages:     224,
    //	Languages:         []gol.Languages{{"/languages/eng"}},
    //	...
    //	Isbn10:            []string{"0195200004"},
    //  ...
    //	OclcNumbers:       []string{"3167790"},
    //	Works:             []gol.Works{{"/works/OL1966513W"}},
    //  ...
    //	Created:           gol.Time{Type: "/type/datetime", Value: "2008-04-01T03:28:50.625462"},
    //	LastModified:      gol.Time{Type: "/type/datetime", Value: "2021-03-03T05:21:06.382367"},
    // }
```

### CoverAPI Examples
```go
    cover := gol.GetBookCoverURL("olid", "OL4554174M", "S")
    // Output:
    // http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg
    
    cover := gol.GetAuthorCoverURL("olid", "OL229501A", "S")
    // Output:
    // http://covers.openlibrary.org/a/olid/OL229501A-S.jpg
```
