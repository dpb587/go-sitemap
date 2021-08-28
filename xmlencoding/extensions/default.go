package extensions

import (
	"encoding/xml"
	"io"

	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
	"go.dpb.io/sitemap/xmlencoding/extensions/googleimage"
	"go.dpb.io/sitemap/xmlencoding/extensions/googlenews"
	"go.dpb.io/sitemap/xmlencoding/extensions/xhtml"
	"go.dpb.io/sitemap/xmlencoding/sitemapxml"
)

// SitemapxmlDecoderSet allows processing sitemap.xml and includes extension schemas.
var SitemapxmlDecoderSet = xmlencoding.NewElementDecoderSet(
	sitemapxml.DefaultSitemapindexDecoder,
	sitemapxml.UrlsetDecoder(
		sitemapxml.UrlsetUrlDecoder(
			googleimage.UrlImageDecoder,
			googlenews.UrlNewsDecoder,
			xhtml.UrlLinkDecoder,
		),
	),
)

// Decode processes the input according to the sitemap.xml with extension schemas rules.
func Decode(r io.Reader, callback data.EntryCallback) error {
	return SitemapxmlDecoderSet.Decode(xml.NewDecoder(r), callback, nil, nil)
}
