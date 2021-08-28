package atomfeed_test

import (
	"bytes"
	"testing"
	"time"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding/atomfeed"
)

func TestParse_WikipediaExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := atomfeed.Decode(
		// https://en.wikipedia.org/wiki/Atom_(Web_standard)#Example_of_an_Atom_1.0_feed
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Example Feed</title>
	<subtitle>A subtitle.</subtitle>
	<link href="http://example.org/feed/" rel="self" />
	<link href="http://example.org/" />
	<id>urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6</id>
	<updated>2003-12-13T18:30:02Z</updated>
	<entry>
		<title>Atom-Powered Robots Run Amok</title>
		<link href="http://example.org/2003/12/13/atom03" />
		<link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
		<link rel="edit" href="http://example.org/2003/12/13/atom03/edit"/>
		<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
		<updated>2003-12-13T18:30:02Z</updated>
		<summary>Some text.</summary>
		<content type="xhtml">
			<div xmlns="http://www.w3.org/1999/xhtml">
				<p>This is the entry content.</p>
			</div>
		</content>
		<author>
			<name>John Doe</name>
			<email>johndoe@example.com</email>
		</author>
	</entry>
</feed>
`)),
		&aggregate,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if _e, _a := 1, len(aggregate); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}

	entry, ok := aggregate[0].(*data.URL)
	if !ok {
		t.Fatalf("unexpected type: %T", aggregate[0])
	} else if _e, _a := "http://example.org/2003/12/13/atom03", entry.Location; _e != _a {
		t.Fatalf("expected `%v` but got: `%v`", _e, _a)
	}

	lastModifiedAsTime, present, err := entry.LastModifiedAsTime()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if !present {
		t.Fatalf("expected field to be present: LastModified")
	} else if _e, _a := "2003-12-13T18:30:02Z", lastModifiedAsTime.Format(time.RFC3339); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}
