package option

// Insert inserts value into opt, then returns a pointer to the contained
// value.
//
// If opt already contains a value, the old value is overwritten.
//
// The returned pointer aliases opt's storage. It is valid until the next
// [Option.Take] or [Option.Replace] on opt. Use [Option.Ptr] for a pointer to
// a copy that does not alias opt.
//
// Insert must be called on an addressable [Option]. Map elements cannot take
// pointer methods in place.
//
// See also [Option.GetOrInsert], which does not overwrite a contained value.
func (opt *Option[T]) Insert(value T) *T {
	opt.value = value
	opt.given = true
	return &opt.value
}

// GetOrInsert inserts value into opt if opt contains no value, then returns a
// pointer to the contained value.
//
// The returned pointer aliases opt's storage. See [Option.Insert].
//
// value is evaluated before GetOrInsert is called. Use [Option.GetOrInsertWith]
// to compute a fallback only when it is needed.
//
// See also [Option.Insert], which overwrites the contained value even if opt
// already contains one.
func (opt *Option[T]) GetOrInsert(value T) *T {
	return opt.GetOrInsertWith(func() T { return value })
}

// GetOrInsertDefault inserts the zero value of T into opt if opt contains no
// value, then returns a pointer to the contained value.
//
// The returned pointer aliases opt's storage. See [Option.Insert].
func (opt *Option[T]) GetOrInsertDefault() *T {
	return opt.GetOrInsertWith(func() T { var t T; return t })
}

// GetOrInsertWith inserts a value computed from f into opt if opt contains no
// value, then returns a pointer to the contained value.
//
// The returned pointer aliases opt's storage. See [Option.Insert].
//
// The function f is not called when opt contains a value.
func (opt *Option[T]) GetOrInsertWith(f func() T) *T {
	if opt.IsNone() {
		opt.value = f()
		opt.given = true
	}

	return &opt.value
}

// Take returns the previous [Option] and leaves [None] in its place.
func (opt *Option[T]) Take() Option[T] {
	old := *opt
	*opt = None[T]()
	return old
}

// TakeIf takes the value out of opt if p returns true for the contained
// value.
//
// If opt contains a value and p returns true, TakeIf is equivalent to
// [Option.Take]. Otherwise, it leaves opt unchanged and returns [None].
//
// p receives a copy of the contained value and cannot mutate opt, even when
// it returns false.
//
// The function p is not called when opt contains no value.
func (opt *Option[T]) TakeIf(p func(T) bool) Option[T] {
	if opt.IsNone() || !p(opt.value) {
		return None[T]()
	}

	return opt.Take()
}

// Replace replaces the contained value with value and returns the previous
// [Option]. After Replace, opt contains value.
func (opt *Option[T]) Replace(value T) Option[T] {
	old := *opt
	opt.Insert(value)
	return old
}
