package data

import (
	"errors"
	"strconv"
	"time"
)

type URL struct {
	Location        string
	LastModified    *string
	ChangeFrequency *string
	Priority        *string

	entryType               EntryType
	offset                  int64
	extensions              EntryList
	lastModifiedTimeFormats []string
}

var _ Entry = URL{}

func NewURL(entryType EntryType, offset int64, lastModifiedTimeFormats []string) *URL {
	return &URL{
		entryType:               entryType,
		offset:                  offset,
		lastModifiedTimeFormats: lastModifiedTimeFormats,
	}
}

func (e URL) GetType() EntryType {
	return e.entryType
}

func (e URL) GetOffset() int64 {
	return e.offset
}

func (e URL) GetExtensions() EntryList {
	return e.extensions
}

func (e *URL) AddExtension(entry Entry) {
	e.extensions = append(e.extensions, entry)
}

// LastModifiedAsTime will return a parsed form of the LastModified field. If there was no value for the field, the
// bool return will be false. If there was an error parsing as a time, the error return will be non-nil.
func (e URL) LastModifiedAsTime() (time.Time, bool, error) {
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

var sitemapChangefreqValues = map[string]struct{}{
	"always":  {},
	"hourly":  {},
	"daily":   {},
	"weekly":  {},
	"monthly": {},
	"yearly":  {},
	"never":   {},
}

// ChangeFrequencyAsString will return a validated form of the ChangeFrequency field. If there was no value for the
// field, the bool return will be false. If the value is non-standard, the error return will be non-nil.
func (e URL) ChangeFrequencyAsString() (string, bool, error) {
	if e.ChangeFrequency == nil {
		return "", false, nil
	}

	changefreq := *e.ChangeFrequency

	if _, known := sitemapChangefreqValues[changefreq]; known {
		return changefreq, true, nil
	}

	return "", true, errors.New("unrecognized value")
}

// PriorityAsFloat32 will return a parsed form of the Priority field. If there was no value for the field, the bool
// return will be false. If there was an error parsing as a float, the error return will be non-nil.
func (e URL) PriorityAsFloat32() (float32, bool, error) {
	if e.Priority == nil {
		return 0, false, nil
	}

	val, err := strconv.ParseFloat(*e.Priority, 32)
	if err != nil {
		return 0, true, err
	}

	return float32(val), true, nil
}
