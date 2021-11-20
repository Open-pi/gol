# Depricated! (needs updating)
# AuthorAPI
The AuthorAPI deals with everything related to Author pages on Open Library that begin with the URL prefix `/authors`.

In these docs you will find all the methods of the `Author` struct. To populate the struct you have to firt use the the `GetAuthor`.

*If you are looking to get the Author from a work or edition, head out to that API's documentation i.e (`docs/WorkAPI.md` or `docs/BookAPI.md` etc.)*

### More about `Author` struct
The fields can be found [here](https://pkg.go.dev/github.com/Open-pi/gol#Author)

Author satisfy these interfaces:
* Coverer


### AuthorAPI Examples
Here are some examples of how you could use the AuthorAPI.
```go
    author := gol.GetAuthor("OL236174A") // Get the Author Richard Dawkins
	// Output:
	// gol.Author {
	//  Bio:    gol.Bio{Type: "/type/text", Value: "Clinton Richard Dawkins, FRS, FRSL is a British ethologist, evolutionary biologist and popular science author..."},
	//  Name:           "Richard Dawkins",
	//  Title:          "FRS, FRSL",
	//  PersonalName:   "Richard Dawkins",
	//  ....
	// }
	
	works := author.Works() // Get the Works of the author
	// Output:
	// []Works
	
	// Get the Author's Photo
	cover := author.Cover("L") // Alternatively: Cover(author, size)
	// Output:
	// string
```