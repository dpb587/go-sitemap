package sitemapxml_test

import (
	"bytes"
	"testing"
	"time"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding/sitemapxml"
)

func TestParse_DocumentationUrlsetExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := sitemapxml.Decode(
		// https://www.sitemaps.org/protocol.html#sitemapXMLExample
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>http://www.example.com/</loc>
		<lastmod>2005-01-01</lastmod>
		<changefreq>monthly</changefreq>
		<priority>0.8</priority>
	</url>
	<url>
		<loc>http://www.example.com/catalog?item=12&amp;desc=vacation_hawaii</loc>
		<changefreq>weekly</changefreq>
	</url>
	<url>
		<loc>http://www.example.com/catalog?item=73&amp;desc=vacation_new_zealand</loc>
		<lastmod>2004-12-23</lastmod>
		<changefreq>weekly</changefreq>
	</url>
	<url>
		<loc>http://www.example.com/catalog?item=74&amp;desc=vacation_newfoundland</loc>
		<lastmod>2004-12-23T18:00:15+00:00</lastmod>
		<priority>0.3</priority>
	</url>
	<url>
		<loc>http://www.example.com/catalog?item=83&amp;desc=vacation_usa</loc>
		<lastmod>2004-11-23</lastmod>
	</url>
</urlset>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 5, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	{
		entry, ok := aggregate[0].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		lastModifiedAsTime, present, err := entry.LastModifiedAsTime()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: LastModified")
		} else if _e, _a := "2005-01-01T00:00:00Z", lastModifiedAsTime.Format(time.RFC3339); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		changeFrequencyAsString, present, err := entry.ChangeFrequencyAsString()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: ChangeFrequency")
		} else if _e, _a := "monthly", changeFrequencyAsString; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		priorityAsFloat32, present, err := entry.PriorityAsFloat32()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: Priority")
		} else if _e, _a := float32(0.8), priorityAsFloat32; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}

	{
		entry, ok := aggregate[1].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/catalog?item=12&desc=vacation_hawaii", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		_, present, err := entry.LastModifiedAsTime()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if present {
			t.Fatalf("expected field to not be present: LastModified")
		}

		changeFrequencyAsString, present, err := entry.ChangeFrequencyAsString()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: ChangeFrequency")
		} else if _e, _a := "weekly", changeFrequencyAsString; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		_, present, err = entry.PriorityAsFloat32()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if present {
			t.Fatalf("expected field to not be present: Priority")
		}
	}
}

func TestParse_DocumentationSitemapindexExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := sitemapxml.Decode(
		// https://www.sitemaps.org/protocol.html#sitemapIndexXMLExample
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<sitemap>
		<loc>http://www.example.com/sitemap1.xml.gz</loc>
		<lastmod>2004-10-01T18:23:17+00:00</lastmod>
	</sitemap>
	<sitemap>
		<loc>http://www.example.com/sitemap2.xml.gz</loc>
		<lastmod>2005-01-01</lastmod>
	</sitemap>
</sitemapindex>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 2, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	{
		entry, ok := aggregate[0].(*data.Sitemap)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/sitemap1.xml.gz", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		lastModifiedAsTime, present, err := entry.LastModifiedAsTime()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: LastModified")
		} else if _e, _a := "2004-10-01T18:23:17Z", lastModifiedAsTime.Format(time.RFC3339); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}

	{
		entry, ok := aggregate[1].(*data.Sitemap)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/sitemap2.xml.gz", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}

		lastModifiedAsTime, present, err := entry.LastModifiedAsTime()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else if !present {
			t.Fatalf("expected field to be present: LastModified")
		} else if _e, _a := "2005-01-01T00:00:00Z", lastModifiedAsTime.Format(time.RFC3339); _e != _a {
			t.Fatalf("expected `%v` but got: %v", _e, _a)
		}
	}
}
