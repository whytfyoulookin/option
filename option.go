package option

type Option[T any] struct {
	value T
	given bool
}

func None[T any]() Option[T] {
	return Option[T]{given: false}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, given: true}
}

func (opt Option[T]) IsNone() bool {
	return !opt.given
}

func (opt Option[T]) IsSome() bool {
	return opt.given
}
