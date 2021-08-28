package xmlencoding

import (
	"encoding/xml"
	"io"

	"github.com/pkg/errors"
	"go.dpb.io/sitemap/data"
)

var emptyDecoderSet = ElementDecoderSet{}

// SkipElement will drop any elements until it reaches the end of the passed element.
func SkipElement(decoder *xml.Decoder, element *Element) error {
	return emptyDecoderSet.Decode(decoder, nil, nil, element)
}

// ElementDecoderSet defines a set of expected element types which can be decoded.
type ElementDecoderSet struct {
	elementDecoders map[xml.Name]ElementDecoder
}

// NewElementDecoderSet uses the passed ElementDecoder list to create a new ElementDecoderSet.
func NewElementDecoderSet(elementDecoders ...ElementDecoder) ElementDecoderSet {
	eds := ElementDecoderSet{
		elementDecoders: map[xml.Name]ElementDecoder{},
	}

	for _, elementDecoder := range elementDecoders {
		for _, name := range elementDecoder.Names() {
			eds.elementDecoders[name] = elementDecoder
		}
	}

	return eds
}

// Decode uses the passed Decoder for processing all children which are of the supported element types.
func (eds *ElementDecoderSet) Decode(decoder *xml.Decoder, callback data.EntryCallback, scope data.Entry, element *Element) error {
	for {
		offset := decoder.InputOffset()
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if _, ok := token.(xml.EndElement); ok {
			return nil
		}

		startElement, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		element := &Element{
			StartElement: &startElement,
			Offset:       offset,
		}

		elementDecoder, ok := eds.elementDecoders[startElement.Name]
		if !ok {
			err = SkipElement(decoder, element)
			if err != nil {
				return errors.Wrap(err, "decoding skip")
			}

			continue
		}

		err = elementDecoder.Decode(decoder, callback, scope, element)
		if err != nil {
			return errors.Wrapf(err, "decoding %s[offset=%d]", startElement.Name.Local, offset)
		}
	}

	return nil
}
