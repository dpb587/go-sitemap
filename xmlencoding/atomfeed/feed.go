package atomfeed

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

func FeedDecoder(extensions ...xmlencoding.ElementDecoder) feedDecoder {
	return feedDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				FeedEntryDecoder(),
			},
			extensions...,
		)...),
	}
}

type feedDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = feedDecoder{}

func (feedDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema, Local: "feed"}}
}

func (d feedDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	return d.ds.Decode(decoder, callback, scope, element)
}

//

func FeedEntryDecoder(extensions ...xmlencoding.ElementDecoder) feedEntryDecoder {
	return feedEntryDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				feedEntryLinkDecoderInstance,
				feedEntryUpdatedDecoderInstance,
			},
			extensions...,
		)...),
	}
}

type feedEntryDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = feedEntryDecoder{}

func (feedEntryDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema, Local: "entry"}}
}

func (d feedEntryDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	uriEntry := data.NewURL(
		xmlencoding.NameToEntryType(element.StartElement.Name),
		element.Offset,
		UpdatedTimeFormats,
	)

	err := d.ds.Decode(decoder, callback, uriEntry, element)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	err = callback.WithEntry(uriEntry)
	if err != nil {
		return errors.Wrap(err, "invoking callback")
	}

	return nil
}

//

var feedEntryLinkDecoderInstance = feedEntryLinkDecoder{}

type feedEntryLinkDecoder struct{}

var _ xmlencoding.ElementDecoder = feedEntryLinkDecoder{}

func (feedEntryLinkDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema, Local: "link"}}
}

func (d feedEntryLinkDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	var rel, href string

	for _, attr := range element.StartElement.Attr {
		if attr.Name == (xml.Name{Local: "rel"}) {
			rel = attr.Value
		} else if attr.Name == (xml.Name{Local: "href"}) {
			href = attr.Value
		}
	}

	if len(rel) == 0 {
		scope.(*data.URL).Location = href
	}

	return xmlencoding.SkipElement(decoder, element)
}

//

var feedEntryUpdatedDecoderInstance = feedEntryUpdatedDecoder{}

type feedEntryUpdatedDecoder struct{}

var _ xmlencoding.ElementDecoder = feedEntryUpdatedDecoder{}

func (feedEntryUpdatedDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema, Local: "updated"}}
}

func (d feedEntryUpdatedDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).LastModified, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
