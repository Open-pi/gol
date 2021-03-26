# Gol
**This is still WIP**

gol (**G**olang **O**pen **L**ibrary) is an interface for the OpenLibrary API. For more information about the API itself please visit [openlibrary.org](https://openlibrary.org/developers/api).

## Functions
These are the API functions (WorkAPI, EditionAPI, etc) to get the data. With them you can make other calls.

| Functions | Args | Returns  |
|---|---|--|
| GetWork   | WorkId | (w Work, err error)  |
| GetEdition   | OLID | (b Book, err error)  |
| GetCoverURL   | coverType, identifierType, identifier, size | string  |
| GetBookCoverURL   | identifierType, identifier, size | string  |
| GetAuthorCoverURL   | identifierType, identifier, size | string  |
| GetAuthor | Author Id | (a Author)  |

For more information, browse the `docs` folder where you can find additional about every API and its subsequent methods.
* [WorkAPI](docs/WorkAPI.md)
* [BookAPI](docs/BookAPI.md)
* [CoverAPI](docs/CoverAPI.md)
* [AuthorAPI](docs/AuthorAPI.md)
