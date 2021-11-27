# WorkAPI
This is the documentation page related to Work pages on Open Library that begin with the URL prefix "/works".
You will find here all the operations you can do on works through the `Work` struct.

In it's simplest form, `GetWork` will fetch all the data and returns a filled* `Work` struct. And from the struct you can make other calls to get Editions, Authors, and Covers etc.

\* filling the struct requires using the Load() method; When using `GetWork()`, `Load()` is called.

### List of methods
| Methods | Args | Returns  |
|---|---|--|
| GetWork   | WorkId | (w Work, err error)  |
| (w *Work) Load | | |
| (w *Work) Key | | string, error |
| (w *Work) Desc | | string, error |
| (w *Work) Subjects | | string, error |
| (w *Work) Title | | string, error |
| (w *Work) KeyAuthors | | []string, error |
| (w *Work) KeyCovers | | []string, error |
|---|---|--|
| (w Work) FirstCoverKey | | string  |
| (w Work) Cover   | size | URL of cover  |
| (w Work) Authors   |  | []Authors, err  |
| (w Work) Editions |  | []Book, err |

### WorkAPI Examples
```go
    work, err:= gol.GetWork("OL45583W")
    // Output:
    // Work, error
    
    // To get fields
    title, err := work.Title()
    
    cover := work.Cover("L") // Get a large cover of the work
    // Output:
    // http://covers.openlibrary.org/b/id/5917705-L.jpg 
    
    authors, err := work.Authors() // Get the list of authors that contributed to the work.
    // Output:
    // []Authors, error

    editions, err := work.Editions() // Get the list of editions liked to the work.
    // Output:
    // []Book, error
```
