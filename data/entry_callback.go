package data

type EntryCallback interface {
	WithEntry(entry Entry) error
}

type EntryCallbackFunc func(entry Entry) error

var _ EntryCallback = EntryCallbackFunc(nil)

func (v EntryCallbackFunc) WithEntry(entry Entry) error {
	return v(entry)
}
