package xmlencoding

import (
	"encoding/xml"

	"go.dpb.io/sitemap/data"
)

func NameToEntryType(in xml.Name) data.EntryType {
	return data.EntryType{
		Space: in.Space,
		Local: in.Local,
	}
}
