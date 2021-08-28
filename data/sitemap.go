package data

import (
	"errors"
	"time"
)

type Sitemap struct {
	Location     string
	LastModified *string

	entryType               EntryType
	offset                  int64
	extensions              EntryList
	lastModifiedTimeFormats []string
}

var _ Entry = Sitemap{}

func NewSitemap(entryType EntryType, offset int64, lastModifiedTimeFormats []string) *Sitemap {
	return &Sitemap{
		entryType:               entryType,
		offset:                  offset,
		lastModifiedTimeFormats: lastModifiedTimeFormats,
	}
}

func (e Sitemap) GetType() EntryType {
	return e.entryType
}

func (e Sitemap) GetOffset() int64 {
	return e.offset
}

func (e Sitemap) GetExtensions() EntryList {
	return e.extensions
}

// LastModifiedAsTime will return a parsed form of the LastModified field. If there was no value for the field, the
// bool return will be false. If there was an error parsing as a time, the error return will be non-nil.
func (e Sitemap) LastModifiedAsTime() (time.Time, bool, error) {
	if e.LastModified == nil {
		return time.Time{}, false, nil
	}

	for _, format := range e.lastModifiedTimeFormats {
		t, err := time.Parse(format, *e.LastModified)
		if err == nil {
			return t, true, nil
		}
	}

	return time.Time{}, true, errors.New("unrecognized format")
}
