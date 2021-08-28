package xhtml

import "go.dpb.io/sitemap/data"

type LinkEntry struct {
	Rel      string
	Href     string
	Hreflang string

	entryType data.EntryType
	offset    int64
}

var _ data.Entry = LinkEntry{}

func NewLinkEntry(entryType data.EntryType, offset int64) *LinkEntry {
	return &LinkEntry{
		entryType: entryType,
		offset:    offset,
	}
}

func (e LinkEntry) GetType() data.EntryType {
	return e.entryType
}

func (e LinkEntry) GetOffset() int64 {
	return e.offset
}

func (e LinkEntry) GetExtensions() data.EntryList {
	return nil
}
