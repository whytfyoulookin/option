package option

// Flatten unwraps an [Option] that contains another [Option].
func Flatten[T any](opt Option[Option[T]]) Option[T] {
	return AndThen(opt, func(inner Option[T]) Option[T] { return inner })
}

// Map applies f to the contained value and returns the result as an [Option].
// If opt contains no value, it returns [None].
//
// The function f is not called when opt contains no value.
func Map[T, U any](opt Option[T], f func(T) U) Option[U] {
	return AndThen(opt, func(v T) Option[U] { return Some(f(v)) })
}

// MapOr applies f to the contained value and returns the result. If opt
// contains no value, it returns defaultValue.
//
// defaultValue is evaluated before MapOr is called. Use [MapOrElse] to compute
// a fallback only when it is needed.
//
// The function f is not called when opt contains no value.
func MapOr[T, U any](opt Option[T], defaultValue U, f func(T) U) U {
	return Map(opt, f).UnwrapOr(defaultValue)
}

// MapOrDefault applies f to the contained value and returns the result. If opt
// contains no value, it returns the zero value of U.
//
// The function f is not called when opt contains no value.
func MapOrDefault[T, U any](opt Option[T], f func(T) U) U {
	return Map(opt, f).UnwrapOrDefault()
}

// MapOrElse applies f to the contained value and returns the result. If opt
// contains no value, it calls defaultF and returns that result.
//
// The function defaultF is not called when opt contains a value. The function
// f is not called when opt contains no value.
func MapOrElse[T, U any](opt Option[T], defaultF func() U, f func(T) U) U {
	return Map(opt, f).UnwrapOrElse(defaultF)
}

// ZipWith applies f to the contained values of opt and other and returns
// the result as an [Option]. If either contains no value, it returns [None].
//
// The function f is not called unless both contain a value.
func ZipWith[T, U, V any](
	opt Option[T], other Option[U], f func(T, U) V,
) Option[V] {
	return AndThen(opt, func(t T) Option[V] {
		return Map(other, func(u U) V { return f(t, u) })
	})
}

// Filter returns opt if it contains a value that satisfies f. Otherwise, it
// returns [None].
//
// The function f is not called when opt contains no value.
func (opt Option[T]) Filter(f func(T) bool) Option[T] {
	return AndThen(opt, func(v T) Option[T] {
		if f(v) {
			return Some(v)
		}

		return None[T]()
	})
}

// Inspect calls f with the contained value if opt contains a value.
// It returns opt unchanged.
//
// The function f is not called when opt contains no value.
func (opt Option[T]) Inspect(f func(T)) Option[T] {
	if opt.IsSome() {
		f(opt.value)
	}

	return opt
}

// Reduce combines opt and other with f if both contain a value. If only one
// contains a value, that [Option] is returned. If neither contains a value,
// it returns [None].
//
// The function f is not called unless both contain a value.
//
// Unlike [ZipWith], a single contained value is kept instead of returning
// [None].
func (opt Option[T]) Reduce(other Option[T], f func(T, T) T) Option[T] {
	return ZipWith(opt, other, f).Or(opt).Or(other)
}
