package sitemapxml

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

func UrlsetDecoder(extensions ...xmlencoding.ElementDecoder) urlsetDecoder {
	return urlsetDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				UrlsetUrlDecoder(),
			},
			extensions...,
		)...),
	}
}

type urlsetDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = urlsetDecoder{}

func (urlsetDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "urlset"}}
}

func (d urlsetDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	return d.ds.Decode(decoder, callback, scope, element)
}

//

func UrlsetUrlDecoder(extensions ...xmlencoding.ElementDecoder) urlsetUrlDecoder {
	return urlsetUrlDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				urlsetUrlLocDecoderInstance,
				urlsetUrlLastmodDecoderInstance,
				urlsetUrlChangefreqDecoderInstance,
				urlsetUrlPriorityDecoderInstance,
			},
			extensions...,
		)...),
	}
}

type urlsetUrlDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = urlsetUrlDecoder{}

func (urlsetUrlDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "url"}}
}

func (d urlsetUrlDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	uriEntry := data.NewURL(
		xmlencoding.NameToEntryType(element.StartElement.Name),
		element.Offset,
		LastmodTimeFormats,
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

var urlsetUrlLocDecoderInstance = urlsetUrlLocDecoder{}

type urlsetUrlLocDecoder struct{}

var _ xmlencoding.ElementDecoder = urlsetUrlLocDecoder{}

func (urlsetUrlLocDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "loc"}}
}

func (d urlsetUrlLocDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).Location, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlsetUrlLastmodDecoderInstance = urlsetUrlLastmodDecoder{}

type urlsetUrlLastmodDecoder struct{}

var _ xmlencoding.ElementDecoder = urlsetUrlLastmodDecoder{}

func (urlsetUrlLastmodDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "lastmod"}}
}

func (d urlsetUrlLastmodDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).LastModified, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlsetUrlChangefreqDecoderInstance = urlsetUrlChangefreqDecoder{}

type urlsetUrlChangefreqDecoder struct{}

var _ xmlencoding.ElementDecoder = urlsetUrlChangefreqDecoder{}

func (urlsetUrlChangefreqDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "changefreq"}}
}

func (d urlsetUrlChangefreqDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).ChangeFrequency, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlsetUrlPriorityDecoderInstance = urlsetUrlPriorityDecoder{}

type urlsetUrlPriorityDecoder struct{}

var _ xmlencoding.ElementDecoder = urlsetUrlPriorityDecoder{}

func (urlsetUrlPriorityDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "priority"}}
}

func (d urlsetUrlPriorityDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.URL)).Priority, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
