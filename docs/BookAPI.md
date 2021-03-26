# Book/EditionAPI
Here you can find all the information on how to get Books and what methods the `Book` struct has.

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