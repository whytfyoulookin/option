// Package option provides a generic Option type for representing values that
// may be absent.
//
// This package is experimental. It follows the Rust Option API as closely as
// Go allows, and also provides idiomatic Go helpers such as [FromPtr],
// [Option.Get], and [Option.OkOr].
//
// The zero value of [Option] contains no value and is equivalent to [None].
// Use [Some] to construct an [Option] containing a value.
//
// Operations that work with a single [Option] and do not need extra type
// parameters are methods. Package-level functions are used when an operation
// must introduce another type, such as [And], [AndThen], [Map], and
// [Flatten], or an additional constraint, such as [Compare] and [Equal].
package option
