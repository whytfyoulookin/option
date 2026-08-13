package option

func And[T, U any](optA Option[T], optB Option[U]) Option[U] {
	if optA.IsNone() {
		return None[U]()
	}

	return optB
}

func Or[T any](optA, optB Option[T]) Option[T] {
	if optA.IsSome() {
		return optA
	}

	return optB
}

func Xor[T any](optA, optB Option[T]) Option[T] {
	if optA.IsSome() == optB.IsSome() {
		return None[T]()
	}

	if optB.IsNone() {
		return optA
	}

	return optB
}

func AndThen[T, U any](opt Option[T], f func(T) Option[U]) Option[U] {
	if opt.IsNone() {
		return None[U]()
	}

	return f(opt.value)
}

func OrElse[T any](opt Option[T], f func() Option[T]) Option[T] {
	if opt.IsSome() {
		return opt
	}

	return f()
}
