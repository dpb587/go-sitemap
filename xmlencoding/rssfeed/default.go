package rssfeed

import (
	"encoding/xml"
	"io"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

// DefaultDecoder is the standard decoder for the <rss /> root element.
var DefaultDecoder = RssDecoder()

// DefaultDecoderSet allows processing a standard RSS feed.
var DefaultDecoderSet = xmlencoding.NewElementDecoderSet(
	DefaultDecoder,
)

// Decode processes the input according to the default RSS rules.
func Decode(r io.Reader, callback data.EntryCallback) error {
	return DefaultDecoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}
