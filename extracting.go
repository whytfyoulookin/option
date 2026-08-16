package option

// Expect returns the contained value. It panics with a custom panic message
// provided by msg if opt contains no value.
func (opt Option[T]) Expect(msg string) T {
	if opt.IsNone() {
		panic(msg)
	}

	return opt.value
}

// Unwrap returns the contained value. It panics if opt contains no value.
//
// Because this function may panic, its use is generally discouraged. Panics
// are meant for unrecoverable errors, and may abort the entire program.
// Instead, use [Option.UnwrapOr], [Option.UnwrapOrElse], or
// [Option.UnwrapOrDefault].
func (opt Option[T]) Unwrap() T {
	return opt.Expect("option: Unwrap called on None")
}

// UnwrapOr returns the contained value or a provided default.
//
// Arguments passed to UnwrapOr are eagerly evaluated; if you are passing the
// result of a function call, it is recommended to use [Option.UnwrapOrElse],
// which is lazily evaluated.
func (opt Option[T]) UnwrapOr(defaultValue T) T {
	if opt.IsNone() {
		return defaultValue
	}

	return opt.value
}

// UnwrapOrDefault returns the contained value. If opt contains no value, it
// returns the zero value of T.
func (opt Option[T]) UnwrapOrDefault() T {
	if opt.IsNone() {
		var defaultValue T
		return defaultValue
	}

	return opt.value
}

// UnwrapOrElse returns the contained value. If opt contains no value it calls
// f and returns its result. The function f is not called when opt contains a
// value.
func (opt Option[T]) UnwrapOrElse(f func() T) T {
	if opt.IsNone() {
		return f()
	}

	return opt.value
}
