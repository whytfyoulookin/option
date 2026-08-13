package option

// IsNone reports whether opt contains no value.
func (opt Option[T]) IsNone() bool {
	return !opt.given
}

// IsSome reports whether opt contains a value.
func (opt Option[T]) IsSome() bool {
	return opt.given
}

// IsNoneOr reports whether opt contains no value or the value matches a
// predicate.
func (opt Option[T]) IsNoneOr(f func(T) bool) bool {
	return opt.IsNone() || f(opt.value)
}

// IsSomeAnd reports whether opt contains a value and the value matches a
// predicate.
func (opt Option[T]) IsSomeAnd(f func(T) bool) bool {
	return opt.IsSome() && f(opt.value)
}
