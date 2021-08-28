package xhtml

import (
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

const Schema = "http://www.w3.org/1999/xhtml"

// UrlLinkDecoder supports decoding <link> elements which may be within a <url>.
//
// https://developers.google.com/search/docs/advanced/crawling/localized-versions#sitemap
var UrlLinkDecoder = urlLinkDecoder{}

type urlLinkDecoder struct{}

var _ xmlencoding.ElementDecoder = urlLinkDecoder{}

func (urlLinkDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema, Local: "link"}}
}

func (d urlLinkDecoder) Decode(decoder *xml.Decoder, _ data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	urlEntry, ok := scope.(*data.URL)
	if !ok {
		return fmt.Errorf("received non-URL entry: %T", scope)
	}

	linkEntry := NewLinkEntry(xmlencoding.NameToEntryType(element.StartElement.Name), element.Offset)

	for _, attr := range element.StartElement.Attr {
		switch attr.Name {
		case xml.Name{Space: Schema, Local: "rel"}, xml.Name{Local: "rel"}:
			linkEntry.Rel = attr.Value
		case xml.Name{Space: Schema, Local: "href"}, xml.Name{Local: "href"}:
			linkEntry.Href = attr.Value
		case xml.Name{Space: Schema, Local: "hreflang"}, xml.Name{Local: "hreflang"}:
			linkEntry.Hreflang = attr.Value
		}
	}

	err := xmlencoding.SkipElement(decoder, element)
	if err != nil {
		return errors.Wrap(err, "skipping")
	}

	urlEntry.AddExtension(linkEntry)

	return nil
}
