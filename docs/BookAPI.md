# Book/EditionAPI
Here you can find all the information on how to get Books and what methods the `Book` struct has.

### EditionAPI Examples
In these examples you will find the most useful methods/function, for more low level on Book structs you can look at the documentation.

```go
    book, err := gol.GetEdition("OL4554174M") // Get the edition using the openlibrary ID
    // Output:
    // Book, error
    
    book, err := gol.GetEditionISBN("978-3-16-148410-0") // Get the edition from the ISBN key
    // Output:
    // Book, error
    
    // To load "all" information from the JSON data in the struct
    book.Load()

    // Returns all the information of the book's authors

    // Get the keys/ids of the book's author
    keyauthors, err := book.KeyAuthors()

    // Get the Authors of the book
    authors, err := book.Authors() // Alternatively you can use Authors(book)
    // Output:
    // []Author, error
    
    // To get the cover keys/ids of the book
    keys, err := book.KeyCovers()
    // Output:
    // []String, error
    
    // To get the first (index 0) cover of the book
    key := book.FirstCoverKey()
    // Output:
    // string

    cover := book.Cover("S") // Returns the URL of the book's cover (size Small)
    // Output:
    // url
```