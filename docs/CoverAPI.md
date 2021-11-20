# Depricated! (needs updating)
# CoverAPI
This part of the library returns the urls covers of authors and books. This is mainly a "low level function" as it's easier to call, for example [`work.Cover()`](WorkAPI.md) to get the url cover of a certain work.

### CoverAPI Examples
```go
    cover := gol.GetBookCoverURL("olid", "OL4554174M", "S")
    // Output:
    // http://covers.openlibrary.org/b/olid/OL4554174M-S.jpg
    
    cover := gol.GetAuthorCoverURL("olid", "OL229501A", "S")
    // Output:
    // http://covers.openlibrary.org/a/olid/OL229501A-S.jpg
```
