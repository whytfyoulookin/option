// Package option provides a generic Option type for representing values that
// may be absent.
//
// The zero value of [Option] contains no value and is equivalent to [None].
// Use [Some] to construct an [Option] containing a value.
//
// Combining operations are package-level functions. [And] and [AndThen] may
// change the contained type, and Go methods cannot declare additional type
// parameters. [Or], [Xor], and [OrElse] use the same form to keep the combining
// API consistent.
package option
