package sitemapxml

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

func SitemapindexDecoder(extensions ...xmlencoding.ElementDecoder) sitemapindexDecoder {
	return sitemapindexDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				SitemapindexSitemapDecoder(),
			},
			extensions...,
		)...),
	}
}

type sitemapindexDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = sitemapindexDecoder{}

func (sitemapindexDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "sitemapindex"}}
}

func (d sitemapindexDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	return d.ds.Decode(decoder, callback, scope, element)
}

//

func SitemapindexSitemapDecoder(extensions ...xmlencoding.ElementDecoder) sitemapindexSitemapDecoder {
	return sitemapindexSitemapDecoder{
		ds: xmlencoding.NewElementDecoderSet(append(
			[]xmlencoding.ElementDecoder{
				sitemapindexSitemapLocDecoderInstance,
				sitemapindexSitemapLastmodDecoderInstance,
			},
			extensions...,
		)...),
	}
}

type sitemapindexSitemapDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = sitemapindexSitemapDecoder{}

func (sitemapindexSitemapDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "sitemap"}}
}

func (d sitemapindexSitemapDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	sitemapEntry := data.NewSitemap(
		xmlencoding.NameToEntryType(element.StartElement.Name),
		element.Offset,
		LastmodTimeFormats,
	)

	err := d.ds.Decode(decoder, callback, sitemapEntry, element)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	err = callback.WithEntry(sitemapEntry)
	if err != nil {
		return errors.Wrap(err, "invoking callback")
	}

	return nil
}

//

var sitemapindexSitemapLocDecoderInstance = sitemapindexSitemapLocDecoder{}

type sitemapindexSitemapLocDecoder struct{}

var _ xmlencoding.ElementDecoder = sitemapindexSitemapLocDecoder{}

func (sitemapindexSitemapLocDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "loc"}}
}

func (d sitemapindexSitemapLocDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.Sitemap)).Location, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var sitemapindexSitemapLastmodDecoderInstance = sitemapindexSitemapLastmodDecoder{}

type sitemapindexSitemapLastmodDecoder struct{}

var _ xmlencoding.ElementDecoder = sitemapindexSitemapLastmodDecoder{}

func (sitemapindexSitemapLastmodDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "lastmod"}}
}

func (d sitemapindexSitemapLastmodDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*data.Sitemap)).LastModified, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
