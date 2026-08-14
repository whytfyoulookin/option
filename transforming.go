package option

// Filter returns opt if it contains a value that satisfies f. Otherwise, it
// returns [None].
//
// The function f is not called when opt contains no value.
func Filter[T any](opt Option[T], f func(T) bool) Option[T] {
	return AndThen(opt, func(v T) Option[T] {
		if f(v) {
			return Some(v)
		}

		return None[T]()
	})
}

// Flatten unwraps an [Option] that contains another [Option].
func Flatten[T any](opt Option[Option[T]]) Option[T] {
	return AndThen(opt, func(inner Option[T]) Option[T] { return inner })
}

// Inspect calls f with the contained value if opt contains a value.
// It returns opt unchanged.
//
// The function f is not called when opt contains no value.
func Inspect[T any](opt Option[T], f func(T)) Option[T] {
	if opt.IsSome() {
		f(opt.value)
	}

	return opt
}

// Map applies f to the contained value and returns the result as an [Option].
// If opt contains no value, it returns [None].
func Map[T, U any](opt Option[T], f func(T) U) Option[U] {
	return AndThen(opt, func(v T) Option[U] { return Some(f(v)) })
}

// MapOr applies f to the contained value and returns the result. If opt
// contains no value, it returns defaultValue.
//
// defaultValue is evaluated before MapOr is called. Use [MapOrElse] to compute
// a fallback only when it is needed.
func MapOr[T, U any](opt Option[T], defaultValue U, f func(T) U) U {
	return Map(opt, f).UnwrapOr(defaultValue)
}

// MapOrDefault applies f to the contained value and returns the result. If opt
// contains no value, it returns the zero value of U.
func MapOrDefault[T, U any](opt Option[T], f func(T) U) U {
	return Map(opt, f).UnwrapOrDefault()
}

// MapOrElse applies f to the contained value and returns the result. If opt
// contains no value, it calls defaultF and returns that result.
//
// The function defaultF is not called when opt contains a value.
func MapOrElse[T, U any](opt Option[T], defaultF func() U, f func(T) U) U {
	return Map(opt, f).UnwrapOrElse(defaultF)
}
