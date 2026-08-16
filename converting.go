package option

// Get returns the contained value and true. If opt contains no value, it
// returns the zero value of T and false.
func (opt Option[T]) Get() (T, bool) {
	return opt.value, opt.given
}

// Ptr returns a pointer to a copy of the contained value. If opt contains no
// value, it returns nil.
func (opt Option[T]) Ptr() *T {
	if opt.IsNone() {
		return nil
	}

	t := opt.value
	return &t
}

// OkOr returns the contained value and a nil error. If opt contains no value,
// it returns the zero value of T and err.
//
// err is evaluated before OkOr is called. Use [Option.OkOrElse] to compute an
// error only when it is needed.
func (opt Option[T]) OkOr(err error) (T, error) {
	return opt.OkOrElse(func() error { return err })
}

// OkOrElse returns the contained value and a nil error. If opt contains no
// value, it calls errF and returns the zero value of T and that error.
//
// The function errF is not called when opt contains a value.
func (opt Option[T]) OkOrElse(errF func() error) (T, error) {
	if opt.IsSome() {
		return opt.value, nil
	}

	var t T
	return t, errF()
}
