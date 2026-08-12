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
