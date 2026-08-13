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
