package xmlencoding

import (
	"encoding/xml"

	"go.dpb.io/sitemap/data"
)

// ElementDecoder adds support for decoding a particular element type. It may optionally modify the current scope's
// entry or invoke the callback function.
type ElementDecoder interface {
	// Names returns the node type supported by the decoder.
	Names() []xml.Name

	// Decode receives a starting XML element and must process through the end element. It may modify the scope or invoke
	// the callback function at any point(s).
	Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *Element) error
}
