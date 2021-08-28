package sitemapxml

import (
	"encoding/xml"
	"io"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

// DefaultSitemapindexDecoder is the standard decoder for the <sitemapindex /> root element.
var DefaultSitemapindexDecoder = SitemapindexDecoder()

// DefaultUrlsetDecoder is the standard decoder for the <urlset /> root element.
var DefaultUrlsetDecoder = UrlsetDecoder()

// DefaultDecoderSet allows processing either of the standard sitemap.xml root elements.
var DefaultDecoderSet = xmlencoding.NewElementDecoderSet(
	DefaultSitemapindexDecoder,
	DefaultUrlsetDecoder,
)

// Decode processes the input according to the default sitemap.xml rules.
func Decode(r io.Reader, callback data.EntryCallback) error {
	return DefaultDecoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}
