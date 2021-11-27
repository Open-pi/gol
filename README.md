# Gol
**This is still WIP**

gol (**G**olang **O**pen **L**ibrary) is an interface for the OpenLibrary API. For more information about the API itself please visit [openlibrary.org](https://openlibrary.org/developers/api).

[![Go Reference](https://pkg.go.dev/badge/github.com/Open-pi/gol.svg)](https://pkg.go.dev/github.com/Open-pi/gol)
![go](https://github.com/Open-pi/gol/actions/workflows/go.yml/badge.svg)

As OpenLibrary's data is always changing, under the hood all the JSON data is handled through [\`gabs\`](https://github.com/Jeffail/gabs/). Subsequently if a field is not accessible with a method, `gabs`'s container can be used instead from any `struct`.

## Functions
For more information, browse the `docs` folder where you can find additional about every API and its subsequent methods.
* [Work API](docs/WorkAPI.md)
* [Book API](docs/BookAPI.md)
* [Cover API](docs/CoverAPI.md)
* [Author API](docs/AuthorAPI.md)
* [Subject/Search/Query API](docs/SearchAPI.md)
