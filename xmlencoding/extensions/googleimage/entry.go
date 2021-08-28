package googleimage

import "go.dpb.io/sitemap/data"

type ImageEntry struct {
	Location    string
	Caption     *string
	GeoLocation *string
	Title       *string
	License     *string

	entryType data.EntryType
	offset    int64
}

var _ data.Entry = ImageEntry{}

func NewImageEntry(entryType data.EntryType, offset int64) *ImageEntry {
	return &ImageEntry{
		entryType: entryType,
		offset:    offset,
	}
}

func (e ImageEntry) GetType() data.EntryType {
	return e.entryType
}

func (e ImageEntry) GetOffset() int64 {
	return e.offset
}

func (e ImageEntry) GetExtensions() data.EntryList {
	return nil
}
