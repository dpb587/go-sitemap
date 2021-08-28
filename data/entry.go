package data

import "github.com/pkg/errors"

type EntryType struct {
	Space, Local string
}

type Entry interface {
	GetType() EntryType
	GetOffset() int64
	GetExtensions() EntryList
}

type EntryList []Entry

func (el EntryList) ByType(ref EntryType) EntryList {
	var result EntryList

	for _, e := range el {
		if e.GetType() == ref {
			result = append(result, e)
		}
	}

	return result
}

func (el EntryList) Each(callback EntryCallback) error {
	for eIdx, e := range el {
		err := callback.WithEntry(e)
		if err != nil {
			return errors.Wrapf(err, "entry #%d", eIdx+1)
		}
	}

	return nil
}
