package option

// Option represents a value that may or may not be present. The zero value of
// Option is equivalent to [None].
type Option[T any] struct {
	value T
	given bool
}

// FromOK returns [Some] of v if err is nil. Otherwise, it returns [None].
func FromOK[T any](v T, err error) Option[T] {
	if err != nil {
		return None[T]()
	}
	return Some(v)
}

// FromPtr returns [Some] of the pointed-to value. If ptr is nil, it returns
// [None].
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}

	return Some(*ptr)
}

// None returns an [Option] containing no value.
func None[T any]() Option[T] {
	return Option[T]{given: false}
}

// Some returns an [Option] containing the value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: value, given: true}
}
