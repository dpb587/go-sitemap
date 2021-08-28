package googlenews

import (
	"errors"
	"time"

	"go.dpb.io/sitemap/data"
)

type NewsEntry struct {
	PublicationName     string
	PublicationLanguage string
	PublicationDate     string
	Title               string

	entryType data.EntryType
	offset    int64
}

var _ data.Entry = NewsEntry{}

func NewNewsEntry(entryType data.EntryType, offset int64) *NewsEntry {
	return &NewsEntry{
		entryType: entryType,
		offset:    offset,
	}
}

func (e NewsEntry) GetType() data.EntryType {
	return e.entryType
}

func (e NewsEntry) GetOffset() int64 {
	return e.offset
}

func (e NewsEntry) GetExtensions() data.EntryList {
	return nil
}

func (e NewsEntry) PublicationDateAsTime() (time.Time, bool, error) {
	for _, format := range []string{time.RFC3339, time.RFC3339Nano, "2006-01-02T15:04Z07:00", "2006-01-02"} {
		t, err := time.Parse(format, e.PublicationDate)
		if err == nil {
			return t, true, nil
		}
	}

	return time.Time{}, true, errors.New("unrecognized format")
}
