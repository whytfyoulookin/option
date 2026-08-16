package option_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestFromOK(t *testing.T) {
	tests := []struct {
		name      string
		v         int
		err       error
		wantValue int
		wantNone  bool
	}{
		{"nil error", 7, nil, 7, false},
		{"error", 7, errors.New("missing"), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.FromOK(tt.v, tt.err)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestFromPtr(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := option.FromPtr[int](nil)
		assert.True(t, got.IsNone())
	})

	t.Run("non-nil", func(t *testing.T) {
		x := 7
		got := option.FromPtr(&x)
		assert.Equal(t, 7, got.Unwrap())

		x = 9
		assert.Equal(t, 7, got.Unwrap())
	})
}

func ExampleFromOK() {
	parse := func(s string) (int, error) {
		if s == "7" {
			return 7, nil
		}
		return 0, errors.New("invalid")
	}

	fmt.Println(option.FromOK(parse("7")).UnwrapOr(-1))
	fmt.Println(option.FromOK(parse("x")).UnwrapOr(-1))

	// Output:
	// 7
	// -1
}

func ExampleFromPtr() {
	x := 7
	fmt.Println(option.FromPtr(&x).UnwrapOr(-1))
	fmt.Println(option.FromPtr[int](nil).UnwrapOr(-1))

	// Output:
	// 7
	// -1
}

func ExampleOption() {
	var opt option.Option[int]

	fmt.Println(opt.IsNone())
	// Output: true
}

func ExampleNone() {
	opt := option.None[int]()

	fmt.Println(opt.IsNone())
	// Output: true
}

func ExampleSome() {
	opt := option.Some(10)

	fmt.Println(opt.IsSome())
	fmt.Println(opt.Unwrap())
	// Output:
	// true
	// 10
}
