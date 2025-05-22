package rocketbox

// StringIterator defines a simple iterator over strings.
//
// It mirrors the commonbox.StringIterator interface so existing implementations
// remain compatible.
type StringIterator interface {
	Len() int32
	HasNext() bool
	Next() string
}
