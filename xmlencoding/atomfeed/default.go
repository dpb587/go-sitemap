package atomfeed

import (
	"encoding/xml"
	"io"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

// DefaultDecoder is the standard decoder for the <feed /> root element.
var DefaultDecoder = FeedDecoder()

// DefaultDecoderSet allows processing a standard Atom feed.
var DefaultDecoderSet = xmlencoding.NewElementDecoderSet(
	DefaultDecoder,
)

// Decode processes the input according to the default Atom rules.
func Decode(r io.Reader, callback data.EntryCallback) error {
	return DefaultDecoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}
