package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whytfyoulookin/option"
)

func TestOption_IsNone(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want bool
	}{
		{"none", option.None[int](), true},
		{"some", option.Some(1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.IsNone()
			assert.Equal(t, tt.want, got, "want: %v, got: %v", tt.want, got)
		})
	}
}

func TestOption_IsSome(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want bool
	}{
		{"none", option.None[int](), false},
		{"some", option.Some(1), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.IsSome()
			assert.Equal(t, tt.want, got, "want: %v, got: %v", tt.want, got)
		})
	}
}

func TestOption_IsNoneOr(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		pred func(int) bool
		want bool
	}{
		{
			name: "none",
			give: option.None[int](),
			pred: nil, // не вызывается
			want: true,
		},
		{
			name: "some with true predicate",
			give: option.Some(5),
			pred: func(v int) bool { return v > 0 },
			want: true,
		},
		{
			name: "some with false predicate",
			give: option.Some(-1),
			pred: func(v int) bool { return v > 0 },
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.IsNoneOr(tt.pred)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOption_IsSomeAnd(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		pred func(int) bool
		want bool
	}{
		{
			name: "none",
			give: option.None[int](),
			pred: func(v int) bool { return true }, // не вызывается
			want: false,
		},
		{
			name: "some with true predicate",
			give: option.Some(5),
			pred: func(v int) bool { return v > 0 },
			want: true,
		},
		{
			name: "some with false predicate",
			give: option.Some(-1),
			pred: func(v int) bool { return v > 0 },
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.give.IsSomeAnd(tt.pred)
			assert.Equal(t, tt.want, got)
		})
	}
}

func ExampleOption_IsNone_none() {
	opt := option.None[int]()
	fmt.Println(opt.IsNone())
	// Output: true
}

func ExampleOption_IsNone_some() {
	opt := option.Some(2)
	fmt.Println(opt.IsNone())
	// Output: false
}

func ExampleOption_IsSome_none() {
	opt := option.None[int]()
	fmt.Println(opt.IsSome())
	// Output: false
}

func ExampleOption_IsSome_some() {
	opt := option.Some(2)
	fmt.Println(opt.IsSome())
	// Output: true
}

func ExampleOption_IsNoneOr_none() {
	opt := option.None[int]()
	fmt.Println(opt.IsNoneOr(func(v int) bool { return v > 0 }))
	// Output: true
}

func ExampleOption_IsNoneOr_some() {
	opt := option.Some(0)
	fmt.Println(opt.IsNoneOr(func(v int) bool { return v > 0 }))
	// Output: false
}

func ExampleOption_IsSomeAnd_none() {
	opt := option.None[int]()
	fmt.Println(opt.IsSomeAnd(func(v int) bool { return v > 0 }))
	// Output: false
}

func ExampleOption_IsSomeAnd_some() {
	opt := option.Some(5)
	fmt.Println(opt.IsSomeAnd(func(v int) bool { return v > 0 }))
	// Output: true
}
