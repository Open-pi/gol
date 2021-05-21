# SearchAPI
This section will have all things related to searching and querying the OpenLibrary API.

## SubjectAPI
To start searching for a subject, you use `GetSubject`; which will return a `Subject` struct type. For a more detailed search on subject, you could use `GetSubjectDetails`.
### Examples
*Coming soon*
## SearchAPI
When searching you are able to specify which data you are looking for (author, subject, title, etc.) as specified by the OpenLibrary API. Firstly, you have to construct the SearchURL, using `SearchURL` struct type and methods, that would be fed to `Search`.
### Examples
```go
    // Construct the SearchUrl
    url := gol.SearchUrl().All("the selfish gene").Author("Richard Dawkins").Construct()
    // search
    search, err := gol.Search(url)
```

## QueryAPI
When using the queryAPI, you have to first construct the url with `QueryUrl` and then make the query with `gol.Query()`.
### Examples
```go
    // Construct the QueryUrl
    url := gol.QueryUrl().Author("OL236174A").Limit(2).Construct()
    // make the query
    query, err := gol.Query(url)
    // map[string]interface{}
```