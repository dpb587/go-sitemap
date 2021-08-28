package googlenews

import (
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

const Schema_0_9 = "http://www.google.com/schemas/sitemap-news/0.9"

// UrlNewsDecoder supports decoding <news> elements which may be within a <url>.
//
// https://developers.google.com/search/docs/advanced/sitemaps/news-sitemap
var UrlNewsDecoder = urlNewsDecoder{
	ds: xmlencoding.NewElementDecoderSet(
		urlNewsPublicationDecoderInstance,
		urlNewsPublicationDateDecoderInstance,
		urlNewsTitleDecoderInstance,
	),
}

type urlNewsDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = urlNewsDecoder{}

func (urlNewsDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "news"}}
}

func (d urlNewsDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	urlEntry, ok := scope.(*data.URL)
	if !ok {
		return fmt.Errorf("received non-URL entry: %T", scope)
	}

	newsEntry := NewNewsEntry(xmlencoding.NameToEntryType(element.StartElement.Name), element.Offset)

	err := d.ds.Decode(decoder, callback, newsEntry, element)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	urlEntry.AddExtension(newsEntry)

	return nil
}

//

var urlNewsPublicationDecoderInstance = urlNewsPublicationDecoder{
	ds: xmlencoding.NewElementDecoderSet(
		urlNewsPublicationNameDecoderInstance,
		urlNewsPublicationLanguageDecoderInstance,
	),
}

type urlNewsPublicationDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = urlNewsPublicationDecoder{}

func (urlNewsPublicationDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "publication"}}
}

func (d urlNewsPublicationDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := d.ds.Decode(decoder, callback, scope, element)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlNewsPublicationNameDecoderInstance = urlNewsPublicationNameDecoder{}

type urlNewsPublicationNameDecoder struct{}

var _ xmlencoding.ElementDecoder = urlNewsPublicationNameDecoder{}

func (urlNewsPublicationNameDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "name"}}
}

func (d urlNewsPublicationNameDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*NewsEntry)).PublicationName, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlNewsPublicationLanguageDecoderInstance = urlNewsPublicationLanguageDecoder{}

type urlNewsPublicationLanguageDecoder struct{}

var _ xmlencoding.ElementDecoder = urlNewsPublicationLanguageDecoder{}

func (urlNewsPublicationLanguageDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "language"}}
}

func (d urlNewsPublicationLanguageDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*NewsEntry)).PublicationLanguage, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlNewsPublicationDateDecoderInstance = urlNewsPublicationDateDecoder{}

type urlNewsPublicationDateDecoder struct{}

var _ xmlencoding.ElementDecoder = urlNewsPublicationDateDecoder{}

func (urlNewsPublicationDateDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "publication_date"}}
}

func (d urlNewsPublicationDateDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*NewsEntry)).PublicationDate, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlNewsTitleDecoderInstance = urlNewsTitleDecoder{}

type urlNewsTitleDecoder struct{}

var _ xmlencoding.ElementDecoder = urlNewsTitleDecoder{}

func (urlNewsTitleDecoder) Names() []xml.Name {
	return []xml.Name{{Space: Schema_0_9, Local: "title"}}
}

func (d urlNewsTitleDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*NewsEntry)).Title, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
