package googlenews_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"
	"time"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
	"go.dpb.io/sitemap/xmlencoding/extensions/googlenews"
	"go.dpb.io/sitemap/xmlencoding/sitemapxml"
)

var decoderSet = xmlencoding.NewElementDecoderSet(
	sitemapxml.DefaultSitemapindexDecoder,
	sitemapxml.UrlsetDecoder(
		sitemapxml.UrlsetUrlDecoder(
			googlenews.UrlNewsDecoder,
		),
	),
)

func decode(r io.Reader, callback data.EntryCallback) error {
	return decoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}

func TestParse_DocumentationExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := decode(
		// https://developers.google.com/search/docs/advanced/sitemaps/news-sitemap#example-sitemap-entry
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
				xmlns:news="http://www.google.com/schemas/sitemap-news/0.9">
	<url>
		<loc>http://www.example.org/business/article55.html</loc>
		<news:news>
		<news:publication>
			<news:name>The Example Times</news:name>
			<news:language>en</news:language>
		</news:publication>
		<news:publication_date>2008-12-23</news:publication_date>
			<news:title>Companies A, B in Merger Talks</news:title>
		</news:news>
	</url>
</urlset>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 1, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	{
		entry, ok := aggregate[0].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.org/business/article55.html", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentries := entry.GetExtensions()
		if _e, _a := 1, len(subentries); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok := subentries[0].(*googlenews.NewsEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "The Example Times", subentry.PublicationName; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "en", subentry.PublicationLanguage; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "Companies A, B in Merger Talks", subentry.Title; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		PublicationDateAsTime, present, err := subentry.PublicationDateAsTime()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: LastModified")
		} else if _e, _a := "2008-12-23T00:00:00Z", PublicationDateAsTime.Format(time.RFC3339); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}
}
