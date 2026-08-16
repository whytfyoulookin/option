package option

// And returns [None] if opt contains no value. Otherwise, it returns other.
//
// other is evaluated before And is called. Use [AndThen] when the second
// [Option] should be computed only if opt contains a value.
func And[T, U any](opt Option[T], other Option[U]) Option[U] {
	return AndThen(opt, func(T) Option[U] { return other })
}

// AndThen returns [None] if opt contains no value. Otherwise, it calls f with
// the contained value and returns the result.
//
// The function f is not called when opt contains no value.
func AndThen[T, U any](opt Option[T], f func(T) Option[U]) Option[U] {
	if opt.IsNone() {
		return None[U]()
	}

	return f(opt.value)
}

// Or returns opt if it contains a value. Otherwise, it returns other.
//
// other is evaluated before Or is called. Use [Option.OrElse] when the fallback
// [Option] should be computed only if opt contains no value.
func (opt Option[T]) Or(other Option[T]) Option[T] {
	return opt.OrElse(func() Option[T] { return other })
}

// OrElse returns opt if it contains a value. Otherwise, it calls f and returns
// the result.
//
// The function f is not called when opt contains a value.
func (opt Option[T]) OrElse(f func() Option[T]) Option[T] {
	if opt.IsSome() {
		return opt
	}

	return f()
}

// Xor returns opt if only opt contains a value, or other if only other
// contains a value. It returns [None] otherwise.
func (opt Option[T]) Xor(other Option[T]) Option[T] {
	if opt.IsSome() == other.IsSome() {
		return None[T]()
	}

	return opt.Or(other)
}
