package googleimage_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
	"go.dpb.io/sitemap/xmlencoding/extensions/googleimage"
	"go.dpb.io/sitemap/xmlencoding/sitemapxml"
)

var decoderSet = xmlencoding.NewElementDecoderSet(
	sitemapxml.DefaultSitemapindexDecoder,
	sitemapxml.UrlsetDecoder(
		sitemapxml.UrlsetUrlDecoder(
			googleimage.UrlImageDecoder,
		),
	),
)

func decode(r io.Reader, callback data.EntryCallback) error {
	return decoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}

func TestParse_DocumentationExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := decode(
		// https://developers.google.com/search/docs/advanced/sitemaps/image-sitemaps#example-sitemap
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
				xmlns:image="http://www.google.com/schemas/sitemap-image/1.1">
	<url>
		<loc>http://example.com/sample1.html</loc>
		<image:image>
			<image:loc>http://example.com/image.jpg</image:loc>
		</image:image>
		<image:image>
			<image:loc>http://example.com/photo.jpg</image:loc>
		</image:image>
	</url>
	<url>
		<loc>http://example.com/sample2.html</loc>
		<image:image>
			<image:loc>http://example.com/picture.jpg</image:loc>
			<image:caption>A funny picture of a cat eating cabbage</image:caption>
			<image:geo_location>Lyon, France</image:geo_location>
			<image:title>Cat vs Cabbage</image:title>
			<image:license>http://example.com/image-license</image:license>
		</image:image>
	</url>
</urlset>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 2, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	{
		entry, ok := aggregate[0].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://example.com/sample1.html", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentries := entry.GetExtensions()
		if _e, _a := 2, len(subentries); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok := subentries[0].(*googleimage.ImageEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://example.com/image.jpg", subentry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok = subentries[1].(*googleimage.ImageEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://example.com/photo.jpg", subentry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}

	{
		entry, ok := aggregate[1].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://example.com/sample2.html", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentries := entry.GetExtensions()
		if _e, _a := 1, len(subentries); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok := subentries[0].(*googleimage.ImageEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://example.com/picture.jpg", subentry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if subentry.Caption == nil {
			t.Fatalf("expected field to be non-nil: Caption")
		} else if _e, _a := "A funny picture of a cat eating cabbage", *subentry.Caption; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if subentry.GeoLocation == nil {
			t.Fatalf("expected field to be non-nil: GeoLocation")
		} else if _e, _a := "Lyon, France", *subentry.GeoLocation; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if subentry.Title == nil {
			t.Fatalf("expected field to be non-nil: Title")
		} else if _e, _a := "Cat vs Cabbage", *subentry.Title; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if subentry.License == nil {
			t.Fatalf("expected field to be non-nil: License")
		} else if _e, _a := "http://example.com/image-license", *subentry.License; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}
}
