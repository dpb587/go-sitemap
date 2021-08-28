package xhtml_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
	"go.dpb.io/sitemap/xmlencoding/extensions/xhtml"
	"go.dpb.io/sitemap/xmlencoding/sitemapxml"
)

var decoderSet = xmlencoding.NewElementDecoderSet(
	sitemapxml.DefaultSitemapindexDecoder,
	sitemapxml.UrlsetDecoder(
		sitemapxml.UrlsetUrlDecoder(
			xhtml.UrlLinkDecoder,
		),
	),
)

func decode(r io.Reader, callback data.EntryCallback) error {
	return decoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}

func TestParse_DocumentationExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := decode(
		// https://developers.google.com/search/docs/advanced/crawling/localized-versions#sitemap
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlns:xhtml="http://www.w3.org/1999/xhtml">
	<url>
		<loc>http://www.example.com/english/page.html</loc>
		<xhtml:link
								rel="alternate"
								hreflang="de"
								href="http://www.example.com/deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="de-ch"
								href="http://www.example.com/schweiz-deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="en"
								href="http://www.example.com/english/page.html"/>
	</url>
	<url>
		<loc>http://www.example.com/deutsch/page.html</loc>
		<xhtml:link
								rel="alternate"
								hreflang="de"
								href="http://www.example.com/deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="de-ch"
								href="http://www.example.com/schweiz-deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="en"
								href="http://www.example.com/english/page.html"/>
	</url>
	<url>
		<loc>http://www.example.com/schweiz-deutsch/page.html</loc>
		<xhtml:link
								rel="alternate"
								hreflang="de"
								href="http://www.example.com/deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="de-ch"
								href="http://www.example.com/schweiz-deutsch/page.html"/>
		<xhtml:link
								rel="alternate"
								hreflang="en"
								href="http://www.example.com/english/page.html"/>
	</url>
</urlset>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 3, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	testSubentries := func(subentries data.EntryList) {
		if _e, _a := 3, len(subentries); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok := subentries[0].(*xhtml.LinkEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://www.example.com/deutsch/page.html", subentry.Href; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "de", subentry.Hreflang; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "alternate", subentry.Rel; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok = subentries[1].(*xhtml.LinkEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://www.example.com/schweiz-deutsch/page.html", subentry.Href; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "de-ch", subentry.Hreflang; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "alternate", subentry.Rel; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		subentry, ok = subentries[2].(*xhtml.LinkEntry)
		if !ok {
			t.Fatalf("unexpected type: %T", subentries[0])
		} else if _e, _a := "http://www.example.com/english/page.html", subentry.Href; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "en", subentry.Hreflang; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		} else if _e, _a := "alternate", subentry.Rel; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}

	{
		entry, ok := aggregate[0].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/english/page.html", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		testSubentries(entry.GetExtensions())
	}

	{
		entry, ok := aggregate[1].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/deutsch/page.html", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		testSubentries(entry.GetExtensions())
	}
}
