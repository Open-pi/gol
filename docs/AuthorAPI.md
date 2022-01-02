# AuthorAPI
The AuthorAPI deals with everything related to Author pages on Open Library that begin with the URL prefix `/authors`.

In these docs you will find all the methods of the `Author` struct. To populate the struct you must first call `GetAuthor`.

*If you are looking to get the Author from a work or edition, head out to that API's documentation i.e (`docs/WorkAPI.md` or `docs/BookAPI.md` etc.)*

### More about `Author` struct
The fields can be found [here](https://pkg.go.dev/github.com/Open-pi/gol#Author)

Author satisfy these interfaces:
* HasCovers
* IKeyAuthors


### AuthorAPI Examples
Here are some examples of how you could use the AuthorAPI.
```go

	author := gol.GetAuthor("OL236174A") // Get the Author Richard Dawkins
	// Output:
	// Author
	
	works := author.Works() // Get the Works of the author
	// Output:
	// []Works
	
	// Get the Author's Photo
	cover := author.Cover("L") // Alternatively: Cover(author, size)
	// Output:
	// url to the cover
```
