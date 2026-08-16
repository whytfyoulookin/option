package option

import "iter"

// Collect returns a slice of every contained value in opts. If any element
// contains no value, it returns [None].
//
// If opts is empty, Collect returns [Some] of an empty slice.
func Collect[T any](opts []Option[T]) Option[[]T] {
	ts := make([]T, len(opts))

	for i, opt := range opts {
		if opt.IsNone() {
			return None[[]T]()
		}
		ts[i] = opt.value
	}

	return Some(ts)
}

// Seq returns an iterator over the contained value. If opt contains no value,
// the iterator is empty.
func (opt Option[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		if opt.IsSome() {
			yield(opt.value)
		}
	}
}
