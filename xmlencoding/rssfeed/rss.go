package rssfeed

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

func RssDecoder(extensions ...xmlencoding.ElementDecoder) rssDecoder {
	return rssDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				RssChannelDecoder(),
			},
			extensions...,
		)...),
	}
}

type rssDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = rssDecoder{}

func (rssDecoder) Names() []xml.Name {
	return []xml.Name{{Local: "rss"}}
}

func (d rssDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	return d.ds.Decode(decoder, callback, scope, element)
}

//

func RssChannelDecoder(extensions ...xmlencoding.ElementDecoder) rssChannelDecoder {
	return rssChannelDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				RssChannelItemDecoder(),
			},
			extensions...,
		)...),
	}
}

type rssChannelDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = rssChannelDecoder{}

func (rssChannelDecoder) Names() []xml.Name {
	return []xml.Name{{Local: "channel"}}
}

func (d rssChannelDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	return d.ds.Decode(decoder, callback, scope, element)
}

//

func RssChannelItemDecoder(extensions ...xmlencoding.ElementDecoder) rssChannelItemDecoder {
	return rssChannelItemDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				rssChannelItemLinkDecoderInstance,
				rssChannelItemPubDateDecoderInstance,
			},
			extensions...,
		)...),
	}
}

type rssChannelItemDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = rssChannelItemDecoder{}

func (rssChannelItemDecoder) Names() []xml.Name {
	return []xml.Name{{Local: "item"}}
}

func (d rssChannelItemDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	uriEntry := data.NewURL(
		xmlencoding.NameToEntryType(element.StartElement.Name),
		element.Offset,
		PubDateTimeFormats,
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

var rssChannelItemLinkDecoderInstance = rssChannelItemLinkDecoder{}

type rssChannelItemLinkDecoder struct{}

var _ xmlencoding.ElementDecoder = rssChannelItemLinkDecoder{}

func (rssChannelItemLinkDecoder) Names() []xml.Name {
	return []xml.Name{{Local: "link"}}
}

func (d rssChannelItemLinkDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).Location, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var rssChannelItemPubDateDecoderInstance = rssChannelItemPubDateDecoder{}

type rssChannelItemPubDateDecoder struct{}

var _ xmlencoding.ElementDecoder = rssChannelItemPubDateDecoder{}

func (rssChannelItemPubDateDecoder) Names() []xml.Name {
	return []xml.Name{{Local: "pubDate"}}
}

func (d rssChannelItemPubDateDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).LastModified, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
