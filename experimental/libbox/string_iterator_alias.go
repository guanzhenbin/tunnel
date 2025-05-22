package libbox

// StringIterator defines a simple iterator over strings.
//
// It matches the commonbox.StringIterator interface so implementations of
// that interface automatically satisfy this one as well.
type StringIterator interface {
	Len() int32
	HasNext() bool
	Next() string
}
