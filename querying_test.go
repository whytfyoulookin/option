package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
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

func ExampleOption_IsNone() {
	fmt.Println(option.None[int]().IsNone())
	fmt.Println(option.Some(2).IsNone())

	// Output:
	// true
	// false
}

func ExampleOption_IsSome() {
	fmt.Println(option.None[int]().IsSome())
	fmt.Println(option.Some(2).IsSome())

	// Output:
	// false
	// true
}

func ExampleOption_IsNoneOr() {
	positive := func(v int) bool { return v > 0 }

	fmt.Println(option.None[int]().IsNoneOr(positive))
	fmt.Println(option.Some(0).IsNoneOr(positive))

	// Output:
	// true
	// false
}

func ExampleOption_IsSomeAnd() {
	positive := func(v int) bool { return v > 0 }

	fmt.Println(option.None[int]().IsSomeAnd(positive))
	fmt.Println(option.Some(5).IsSomeAnd(positive))

	// Output:
	// false
	// true
}
