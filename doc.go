// Package option provides a generic Option type for representing values that
// may be absent.
//
// The zero value of [Option] contains no value and is equivalent to [None].
// Use [Some] to construct an [Option] containing a value.
//
// Combining and transforming operations are package-level functions. Some of
// them, such as [And], [AndThen], [Map], and [Flatten], may change the
// contained type, and Go methods cannot declare additional type parameters.
// Operations that could be methods, such as [Or], [Filter], and [Inspect],
// use the same form to keep these APIs consistent.
package option
