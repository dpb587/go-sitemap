package googleimage

import (
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
	"go.dpb.io/sitemap/xmlencoding"
)

// UrlImageDecoder supports decoding <image> elements which may be within a <url>.
//
// https://developers.google.com/search/docs/advanced/sitemaps/image-sitemaps
var UrlImageDecoder = urlImageDecoder{
	ds: xmlencoding.NewElementDecoderSet(
		urlImageLocDecoderInstance,
		urlImageCaptionDecoderInstance,
		urlImageGeoLocationDecoderInstance,
		urlImageTitleDecoderInstance,
		urlImageLicenseDecoderInstance,
	),
}

type urlImageDecoder struct {
	ds xmlencoding.ElementDecoderSet
}

var _ xmlencoding.ElementDecoder = urlImageDecoder{}

func (urlImageDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "image"}}
}

func (d urlImageDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	urlEntry, ok := scope.(*data.URL)
	if !ok {
		return fmt.Errorf("received non-URL entry: %T", scope)
	}

	imageEntry := NewImageEntry(xmlencoding.NameToEntryType(element.StartElement.Name), element.Offset)

	err := d.ds.Decode(decoder, callback, imageEntry, element)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	urlEntry.AddExtension(imageEntry)

	return nil
}

//

var urlImageLocDecoderInstance = urlImageLocDecoder{}

type urlImageLocDecoder struct{}

var _ xmlencoding.ElementDecoder = urlImageLocDecoder{}

func (urlImageLocDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "loc"}}
}

func (d urlImageLocDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*ImageEntry)).Location, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlImageCaptionDecoderInstance = urlImageCaptionDecoder{}

type urlImageCaptionDecoder struct{}

var _ xmlencoding.ElementDecoder = urlImageCaptionDecoder{}

func (urlImageCaptionDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "caption"}}
}

func (d urlImageCaptionDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*ImageEntry)).Caption, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlImageGeoLocationDecoderInstance = urlImageGeoLocationDecoder{}

type urlImageGeoLocationDecoder struct{}

var _ xmlencoding.ElementDecoder = urlImageGeoLocationDecoder{}

func (urlImageGeoLocationDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "geo_location"}}
}

func (d urlImageGeoLocationDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*ImageEntry)).GeoLocation, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlImageTitleDecoderInstance = urlImageTitleDecoder{}

type urlImageTitleDecoder struct{}

var _ xmlencoding.ElementDecoder = urlImageTitleDecoder{}

func (urlImageTitleDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "title"}}
}

func (d urlImageTitleDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*ImageEntry)).Title, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}

//

var urlImageLicenseDecoderInstance = urlImageLicenseDecoder{}

type urlImageLicenseDecoder struct{}

var _ xmlencoding.ElementDecoder = urlImageLicenseDecoder{}

func (urlImageLicenseDecoder) Names() []xml.Name {
	return []xml.Name{{Space: "http://www.google.com/schemas/sitemap-image/1.1", Local: "license"}}
}

func (d urlImageLicenseDecoder) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *xmlencoding.Element) error {
	err := decoder.DecodeElement(&(scope.(*ImageEntry)).License, element.StartElement)
	if err != nil {
		return errors.Wrap(err, "decoding")
	}

	return nil
}
