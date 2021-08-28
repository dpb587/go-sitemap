package textencoding_test

import (
	"bytes"
	"testing"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/textencoding"
)

func TestParse_DocumentationExample(t *testing.T) {
	aggregate := data.NewEntryAggregate()

	err := textencoding.Decode(
		// https://www.sitemaps.org/protocol.html
		bytes.NewReader([]byte(`
http://www.example.com/catalog?item=1
http://www.example.com/catalog?item=11
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
		} else if _e, _a := "http://www.example.com/catalog?item=1", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: `%v`", _e, _a)
		}
	}

	{
		entry, ok := aggregate[1].(*data.URL)
		if !ok {
			t.Fatalf("unexpected type: %T", aggregate[0])
		} else if _e, _a := "http://www.example.com/catalog?item=11", entry.Location; _e != _a {
			t.Fatalf("expected `%v` but got: `%v`", _e, _a)
		}
	}
}
