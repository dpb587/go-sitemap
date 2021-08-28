package xmlencoding

import "encoding/xml"

type Element struct {
	StartElement *xml.StartElement
	Offset       int64
}
