package option

// And returns [None] if a contains no value. Otherwise, it returns b.
//
// b is evaluated before And is called. Use [AndThen] when the second Option
// should be computed only if a contains a value.
func And[T, U any](a Option[T], b Option[U]) Option[U] {
	return AndThen(a, func(T) Option[U] { return b })
}

// Or returns a if it contains a value. Otherwise, it returns b.
//
// b is evaluated before Or is called. Use [OrElse] when the fallback Option
// should be computed only if a contains no value.
func Or[T any](a, b Option[T]) Option[T] {
	return OrElse(a, func() Option[T] { return b })
}

// Xor returns a if only a contains a value, or b if only b contains a value.
// It returns [None] otherwise.
func Xor[T any](a, b Option[T]) Option[T] {
	if a.IsSome() == b.IsSome() {
		return None[T]()
	}

	return Or(a, b)
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

// OrElse returns opt if it contains a value. Otherwise, it calls f and returns
// the result.
//
// The function f is not called when opt contains a value.
func OrElse[T any](opt Option[T], f func() Option[T]) Option[T] {
	if opt.IsSome() {
		return opt
	}

	return f()
}
