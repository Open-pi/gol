# Depricated! (needs updating)
# WorkAPI
This is the documentation page related to Work pages on Open Library that begin with the URL prefix "/works".
You will find here all the operations you can do on works through the `Work` struct.

In it's simplest form, `GetWork` will fetch all the data and returns a filled `Work` struct. And from the struct you can make other calls to get Editions, Authors, and Covers etc.

### List of methods
| Methods | Args | Returns  |
|---|---|--|
| GetWork   | WorkId | (w Work, err error)  |
| (w Work) Cover   | size | URL of cover  |
| (w Work) Authors   |  | []Authors, err  |
| (w Work) Editions |  | []Book, err |

### WorkAPI Examples
```go
    work := gol.GetWork("OL45583W")
    // Output:
    // gol.Work{
    //  Created:        gol.Time{Type: "/type/datetime", Value: "2009-10-15T11:23:34.130855"},
    //	Subjects:       []string{"History and criticism", "Russian literature", "Russian literature, history and criticism"},
    //	LatestRevision: 4,
    //	Key:            "/works/OL45583W",
    //	Title:          "An outline of Russian literature",
    //	Authors:        []gol.AuthorAndType{{gol.Type{"/type/author_role"}, gol.Author{"/authors/OL18295A"}}},
    //  ...
    // }
    
    cover := work.Cover("L") // Get a large cover of the work
    // Output:
    // http://covers.openlibrary.org/b/id/5917705-L.jpg 
    
    authors, err := work.Authors() // Get the list of authors that contributed to the work.
    // Output:
    // []Authors

    editions, err := work.Editions() // Get the list of editions liked to the work.
    // Output:
    // []Book
```
