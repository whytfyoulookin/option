package option_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whytfyoulookin/option"
)

func TestAnd(t *testing.T) {
	tests := []struct {
		name      string
		a         option.Option[int]
		b         option.Option[string]
		wantValue string
		wantNone  bool
	}{
		{"some none", option.Some(2), option.None[string](), "", true},
		{"none some", option.None[int](), option.Some("foo"), "", true},
		{"some some", option.Some(2), option.Some("foo"), "foo", false},
		{"none none", option.None[int](), option.None[string](), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.And(tt.a, tt.b)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name      string
		a, b      option.Option[int]
		wantValue int
		wantNone  bool
	}{
		{"some none", option.Some(2), option.None[int](), 2, false},
		{"none some", option.None[int](), option.Some(67), 67, false},
		{"some some", option.Some(2), option.Some(67), 2, false},
		{"none none", option.None[int](), option.None[int](), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.Or(tt.a, tt.b)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		name      string
		a, b      option.Option[int]
		wantValue int
		wantNone  bool
	}{
		{"some none", option.Some(2), option.None[int](), 2, false},
		{"none some", option.None[int](), option.Some(67), 67, false},
		{"some some", option.Some(2), option.Some(67), 0, true},
		{"none none", option.None[int](), option.None[int](), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.Xor(tt.a, tt.b)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestAndThen(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue string
		wantNone  bool
		wantCalls int
	}{
		{"some", option.Some(2), "4", false, 1},
		{"none", option.None[int](), "", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.AndThen(tt.give, func(i int) option.Option[string] {
				calls++
				return option.Some(strconv.Itoa(i * i))
			})

			assert.Equal(t, tt.wantCalls, calls)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestOrElse(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"some", option.Some(2), 2, false, 0},
		{"none", option.None[int](), 67, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.OrElse(tt.give, func() option.Option[int] {
				calls++
				return option.Some(67)
			})

			assert.Equal(t, tt.wantCalls, calls)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func ExampleAnd() {
	x := option.Some(2)
	y := option.None[string]()
	fmt.Println(option.And(x, y).UnwrapOr("<none>"))

	x = option.None[int]()
	y = option.Some("foo")
	fmt.Println(option.And(x, y).UnwrapOr("<none>"))

	x = option.Some(2)
	y = option.Some("foo")
	fmt.Println(option.And(x, y).UnwrapOr("<none>"))

	x = option.None[int]()
	y = option.None[string]()
	fmt.Println(option.And(x, y).UnwrapOr("<none>"))

	// Output:
	// <none>
	// <none>
	// foo
	// <none>
}

func ExampleOr() {
	x := option.Some(2)
	y := option.None[int]()
	fmt.Println(option.Or(x, y).UnwrapOr(-1))

	x = option.None[int]()
	y = option.Some(100)
	fmt.Println(option.Or(x, y).UnwrapOr(-1))

	x = option.Some(2)
	y = option.Some(100)
	fmt.Println(option.Or(x, y).UnwrapOr(-1))

	x = option.None[int]()
	y = option.None[int]()
	fmt.Println(option.Or(x, y).UnwrapOr(-1))

	// Output:
	// 2
	// 100
	// 2
	// -1
}

func ExampleXor() {
	x := option.Some(2)
	y := option.None[int]()
	fmt.Println(option.Xor(x, y).UnwrapOr(-1))

	x = option.None[int]()
	y = option.Some(2)
	fmt.Println(option.Xor(x, y).UnwrapOr(-1))

	x = option.Some(2)
	y = option.Some(22)
	fmt.Println(option.Xor(x, y).UnwrapOr(-1))

	x = option.None[int]()
	y = option.None[int]()
	fmt.Println(option.Xor(x, y).UnwrapOr(-1))

	// Output:
	// 2
	// 2
	// -1
	// -1
}

func ExampleAndThen() {
	sqThenToString := func(x uint32) option.Option[string] {
		if x != 0 && x > math.MaxUint32/x {
			return option.None[string]()
		}

		return option.Some(strconv.FormatUint(uint64(x*x), 10))
	}

	fmt.Println(
		option.AndThen(option.Some[uint32](2), sqThenToString).UnwrapOr("<none>"),
	)

	fmt.Println(
		option.AndThen(option.Some[uint32](1_000_000), sqThenToString).
			UnwrapOr("<none>"),
	)

	fmt.Println(
		option.AndThen(option.None[uint32](), sqThenToString).UnwrapOr("<none>"),
	)

	// Output:
	// 4
	// <none>
	// <none>
}

func ExampleOrElse() {
	nobody := func() option.Option[string] { return option.None[string]() }
	vikings := func() option.Option[string] { return option.Some("vikings") }

	fmt.Println(
		option.OrElse(option.Some("barbarians"), vikings).UnwrapOr("<none>"),
	)

	fmt.Println(
		option.OrElse(option.None[string](), vikings).UnwrapOr("<none>"),
	)

	fmt.Println(option.OrElse(option.None[string](), nobody).UnwrapOr("<none>"))

	// Output:
	// barbarians
	// vikings
	// <none>
}
