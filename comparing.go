package option

import "cmp"

// Compare returns -1 if a is less than b, 0 if a equals b, or +1 if a is
// greater than b.
//
// [None] is less than any [Some] value. Two [None] values are equal. If both
// contain values, they are compared with [cmp.Compare].
//
// For floating-point values, [cmp.Compare] treats NaN as equal to NaN. [Equal]
// uses Go ==, so two [Some] NaN values compare equal here and unequal there.
func Compare[T cmp.Ordered](a, b Option[T]) int {
	switch {
	case a.IsNone() && b.IsNone():
		return 0
	case a.IsNone():
		return -1
	case b.IsNone():
		return 1
	default:
		return cmp.Compare(a.value, b.value)
	}
}

// Equal reports whether a and b are equal.
//
// Two [None] values are equal. A [None] is not equal to any [Some]. Two [Some]
// values are equal if their contained values are equal with Go ==.
//
// For floating-point values, two [Some] NaN values are not equal. [Compare]
// uses [cmp.Compare], so those same values compare equal there.
func Equal[T comparable](a, b Option[T]) bool {
	if a.IsNone() || b.IsNone() {
		return a.IsNone() && b.IsNone()
	}

	return a.value == b.value
}

// Ge reports whether a is greater than or equal to b.
//
// See [Compare] for the ordering of [None] and [Some].
func Ge[T cmp.Ordered](a, b Option[T]) bool {
	return Compare(a, b) >= 0
}

// Gt reports whether a is greater than b.
//
// See [Compare] for the ordering of [None] and [Some].
func Gt[T cmp.Ordered](a, b Option[T]) bool {
	return Compare(a, b) > 0
}

// Le reports whether a is less than or equal to b.
//
// See [Compare] for the ordering of [None] and [Some].
func Le[T cmp.Ordered](a, b Option[T]) bool {
	return Compare(a, b) <= 0
}

// Lt reports whether a is less than b.
//
// See [Compare] for the ordering of [None] and [Some].
func Lt[T cmp.Ordered](a, b Option[T]) bool {
	return Compare(a, b) < 0
}
