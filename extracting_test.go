package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestOption_Expect(t *testing.T) {
	const msg = "expected value"
	tests := []struct {
		name      string
		give      option.Option[int]
		want      int
		wantPanic bool
	}{
		{"none", option.None[int](), 0, true},
		{"some", option.Some(10), 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.PanicsWithValue(t, msg, func() { tt.give.Expect(msg) })
				return
			}
			got := tt.give.Expect(msg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOption_Unwrap(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		want      int
		wantPanic bool
	}{
		{"none", option.None[int](), 0, true},
		{"some", option.Some(10), 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { tt.give.Unwrap() })
				return
			}
			got := tt.give.Unwrap()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOption_UnwrapOr(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want int
	}{
		{"none", option.None[int](), 67},
		{"some", option.Some(10), 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.UnwrapOr(67)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOption_UnwrapOrDefault(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want int
	}{
		{"none", option.None[int](), 0},
		{"some", option.Some(10), 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.UnwrapOrDefault()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		want      int
		wantCalls int
	}{
		{"none", option.None[int](), 67, 1},
		{"some", option.Some(10), 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := tt.give.UnwrapOrElse(func() int {
				calls++
				return 67
			})
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}

func ExampleOption_Expect() {
	defer func() { fmt.Println(recover()) }()

	fmt.Println(option.Some("value").Expect("fruits are healthy"))
	option.None[string]().Expect("string should not be empty")

	// Output:
	// value
	// string should not be empty
}

func ExampleOption_Unwrap() {
	defer func() { fmt.Println(recover()) }()

	fmt.Println(option.Some("air").Unwrap())
	option.None[string]().Unwrap()

	// Output:
	// air
	// option: Unwrap called on None
}

func ExampleOption_UnwrapOr() {
	fmt.Println(option.None[string]().UnwrapOr("bike"))
	fmt.Println(option.Some("car").UnwrapOr("bike"))

	// Output:
	// bike
	// car
}

func ExampleOption_UnwrapOrDefault() {
	fmt.Println(option.None[int]().UnwrapOrDefault())
	fmt.Println(option.Some(123).UnwrapOrDefault())

	// Output:
	// 0
	// 123
}

func ExampleOption_UnwrapOrElse() {
	x := 5
	fmt.Println(option.None[int]().UnwrapOrElse(func() int { return 2 * x }))
	fmt.Println(option.Some(6).UnwrapOrElse(func() int { return 2 * x }))

	// Output:
	// 10
	// 6
}
