package option

// Option represents a value that may or may not be present. The zero value of
// Option is equivalent to [None].
type Option[T any] struct {
	value T
	given bool
}

// None returns an [Option] containing no value.
func None[T any]() Option[T] {
	return Option[T]{given: false}
}

// Some returns an [Option] containing value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: value, given: true}
}

// IsNone reports whether opt contains no value.
func (opt Option[T]) IsNone() bool {
	return !opt.given
}

// IsSome reports whether opt contains a value.
func (opt Option[T]) IsSome() bool {
	return opt.given
}
