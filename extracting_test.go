package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whytfyoulookin/option"
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

func ExampleOption_Expect_none() {
	defer func() { fmt.Println(recover()) }()

	opt := option.None[string]()
	opt.Expect("string should not be empty")
	// Output: string should not be empty
}

func ExampleOption_Expect_some() {
	opt := option.Some("value")
	fmt.Println(opt.Expect("fruits are healthy"))
	// Output: value
}

func ExampleOption_Unwrap_none() {
	defer func() { fmt.Println(recover()) }()

	opt := option.None[string]()
	opt.Unwrap()
	// Output: option: Unwrap called on None
}

func ExampleOption_Unwrap_some() {
	opt := option.Some("air")
	fmt.Println(opt.Unwrap())
	// Output: air
}

func ExampleOption_UnwrapOr_none() {
	opt := option.None[string]()
	fmt.Println(opt.UnwrapOr("bike"))
	// Output: bike
}

func ExampleOption_UnwrapOr_some() {
	opt := option.Some("car")
	fmt.Println(opt.UnwrapOr("bike"))
	// Output: car
}

func ExampleOption_UnwrapOrDefault_none() {
	opt := option.None[int]()
	fmt.Println(opt.UnwrapOrDefault())
	// Output: 0
}

func ExampleOption_UnwrapOrDefault_some() {
	opt := option.Some(123)
	fmt.Println(opt.UnwrapOrDefault())
	// Output: 123
}

func ExampleOption_UnwrapOrElse_none() {
	opt := option.None[int]()
	x := 5
	fmt.Println(opt.UnwrapOrElse(func() int { return 2 * x }))
	// Output: 10
}

func ExampleOption_UnwrapOrElse_some() {
	opt := option.Some(6)
	x := 5
	fmt.Println(opt.UnwrapOrElse(func() int { return 2 * x }))
	// Output: 6
}
