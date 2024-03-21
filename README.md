# What is this?

`ngore` is a library written in Go to interact with the https://ncore.pro website. Since the ncore website does not provide an official api, this library uses website scraping to parse the responses.

# How to use

## Add a dependency

To import the library, make sure you are using a recent version of Go.

Use `go get` to add a dependency to the latest release:

```
go get github.com/gar-r/ngore@latest
```

Add a dependency to a specific version:

```
 go get github.com/gar-r/ngore@v0.9.1
```

## Initialize the api

To create an instance of the api with a default client:

```go
api := ngore.Default("https://ncore.pro")	
```

Optionally, you may want to customize the client used by the api:

```go
client := &http.Client{
    Timeout: 30 * time.Second,
}
api := ngore.New(client, "https://ncore.pro")
```

# Usage Examples

## Login

In order to use the api, you will first have to log in. Using the api without logging in, or with an expired login will result in an error being returned.

```go
err := api.Login(&login.BasicAuth{
    UserName: "user",
    Password: "pass",
})
if err != nil {
	// ...
}
```

## Search

### basic search

The search function takes a `search.Params` pointer, which specifies what to search for.
At a minimum, you want to add a search phrase:

```go
res, err := api.Search(&search.Params{
    SearchPhrase: "hortobagy",
})
```

In most cases you will want to narrow down you search a bit more:

```go
func SearchMovieByName(api ngore.Api) error {
    params := &search.Params{
        SearchPhrase: "hortobagy",
        Field:        search.Name,
        Category:     search.MovieSdHu,
    }
	res, err := api.Search(params)
    if err != nil {
        return err
    }
    // process results
    return nil
}
```

Search returns a `search.Result` pointer, which contains the actual result list, and paging information. The following example prints out the name of every search result on the first page:

```go
for _, t := range res.Torrents {
    fmt.Printf("name: %s\n", t.Title)
}
```

### paging

By default, search will request the first page of the results. This can be changed by specifying the value of the `Page` field in the `search.Params`.

The following example performs several searches in order to iterate through all results in all pages:

```go
func SearchWithPaging(api ngore.Api) error {
	params := &search.Params{
		SearchPhrase: "game",
		Field:        search.Name,
		Category:     search.SeriesSdEn,
	}
	var res *search.Result
	var err error
	for params.Page = 0; res == nil || res.Page.HasMore(); params.Page++ {
		fmt.Printf("page %d\n", params.Page)
		res, err = api.Search(params)
		if err != nil {
			return err
		}
		// process results on page
	}
	return nil
}
```

### sorting

Server-side sorting can be requested using the `SortField` and `SortMode` fields:

```go
func SearchWithSort(api ngore.Api) error {
	params := &search.Params{
		SearchPhrase: "penguin",
		SortField:    search.ByDownloaded,
		SortMode:     search.Descending,
	}
	res, err := api.Search(params)
	if err != nil {
		return err
	}
	// process results
	return nil
}
```

## User Activity 

Several user activity stats can be requested from the server. This is represented by the `activity.Info` struct.

The following example demonstrates how to fetch the user activity details:

```go
func PrintActivity(api ngore.Api) error {
	info, err := api.Activity()
	if err != nil {
		return err
	}
	PrintJson(info)
	return nil
}
```

Result :

```json
{
  "Rank": {
    "Daily": "111",
    "Weekly": "222",
    "Monthly": "333",
    "PrevMonth": "444"
  },
  "Stats": {
    "Current": "0",
    "Allowed": "korlátlan",
    "PenMonths": "5",
    "PenTorrents": "0"
  },
  "CanDownload": "igen",
  "History": []
}
```


## Download

In order to download, you only need to provide the torrent id, and have a valid api key (automatically acquired by the api after logging in).

The following example downloads a torrent:

```go
func Download(api ngore.Api) error {
	t := // acquired by searching previously
	bytes, err := api.Download(t.Id)
	if err != nil {
		return err
	}
	// process downloaded data
	return nil
}
```