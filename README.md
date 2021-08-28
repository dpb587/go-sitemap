# go.dpb.io/sitemap

Stream decoders for `sitemap.xml` data and link feeds.

 * supports multiple file formats - traditional `sitemap.xml` files, syndication feeds (Atom, RSS), and plain text files
 * supports extensions - [`xhtml:link` localizations](https://developers.google.com/search/docs/advanced/crawling/localized-versions#sitemap), Google-specific ([image sitemaps](https://developers.google.com/search/docs/advanced/sitemaps/image-sitemaps) and [Google News sitemap](https://developers.google.com/search/docs/advanced/sitemaps/image-sitemaps))
 * supports stream parsing (vs parsing all records into memory)
 * utilities for resolving relative URLs, ignoring invalid locations, and fetching directly from a server

## Usage

Import the base packages...

```go
import (
  "go.dpb.io/sitemap"
  "go.dpb.io/sitemap/data"
)
```

Create an `EntryCallback` for processing each `Entry` as it is decoded...

```go
callback := data.EntryCallbackFunc(func (entry data.Entry) error {
  if entryURL, ok := e.(*data.URL); ok {
    fmt.Println(entryURL.Location)
  }

  return nil
})
```

Use the default decoder which will auto-detect the encoding and supports some common extensions...

```go
err := sitemap.Decode(reader, callback)
```

Learn more from the [`examples` directory](examples), [test files](https://github.com/dpb587/go-sitemap/search?q=filename%3A_test.go), and [code documentation](https://pkg.go.dev/go.dpb.io/sitemap).

## Alternatives

 * [oxffaa/gopher-parse-sitemap](https://github.com/oxffaa/gopher-parse-sitemap) - stream parsing but doesn't support extensions
 * [yterajima/go-sitemap](https://github.com/yterajima/go-sitemap) - loads everything into memory and doesn't support extensions

## License

[MIT License](LICENSE)
