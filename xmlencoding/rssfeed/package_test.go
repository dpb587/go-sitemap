package rssfeed_test

import (
	"bytes"
	"testing"
	"time"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding/rssfeed"
)

func TestParse_WikipediaExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := rssfeed.Decode(
		// https://en.wikipedia.org/wiki/RSS#Example
		bytes.NewReader([]byte(`
<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
<channel>
	<title>RSS Title</title>
	<description>This is an example of an RSS feed</description>
	<link>http://www.example.com/main.html</link>
	<copyright>2020 Example.com All rights reserved</copyright>
	<lastBuildDate>Mon, 06 Sep 2010 00:01:00 +0000</lastBuildDate>
	<pubDate>Sun, 06 Sep 2009 16:20:00 +0000</pubDate>
	<ttl>1800</ttl>

	<item>
	<title>Example entry</title>
	<description>Here is some text containing an interesting description.</description>
	<link>http://www.example.com/blog/post/1</link>
	<guid isPermaLink="false">7bd204c6-1655-4c27-aeee-53f933c5395f</guid>
	<pubDate>Sun, 06 Sep 2009 16:20:00 +0000</pubDate>
	</item>

</channel>
</rss>
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
	} else if _e, _a := "http://www.example.com/blog/post/1", entry.Location; _e != _a {
		t.Fatalf("expected `%v` but got: `%v`", _e, _a)
	}

	lastModifiedAsTime, present, err := entry.LastModifiedAsTime()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else if !present {
		t.Fatalf("expected field to be present: LastModified")
	} else if _e, _a := "2009-09-06T16:20:00Z", lastModifiedAsTime.Format(time.RFC3339); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}
