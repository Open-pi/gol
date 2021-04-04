# Book/EditionAPI
Here you can find all the information on how to get Books and what methods the `Book` struct has.

### EditionAPI Examples
In these examples you will find the most useful methods/function, for more low level on Book structs you can look at the documentation.

```go
    book := gol.GetEdition("OL4554174M") // Get the edition using the openlibrary ID
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
    
    book := gol.GetEditionISBN("978-3-16-148410-0") // Get the edition from the ISBN key
    // Output:
    // gol.Book{
    // ...
    // }
    
    // Returns all the information of the book's authors
    authors := book.Authors() // Alternatively you can use Authors(book)
    // Output:
    // []Author
    
    cover := book.Cover("S") // Returns the URL of the book's cover (size Small)
    // Output:
    // string
```