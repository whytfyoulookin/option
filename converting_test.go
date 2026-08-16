package option_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestOption_Get(t *testing.T) {
	tests := []struct {
		name     string
		give     option.Option[int]
		want     int
		wantSome bool
	}{
		{"none", option.None[int](), 0, false},
		{"some", option.Some(7), 7, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.give.Get()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantSome, ok)
		})
	}
}

func TestOption_Ptr(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		assert.Nil(t, option.None[int]().Ptr())
	})

	t.Run("some", func(t *testing.T) {
		opt := option.Some(7)
		p := opt.Ptr()
		assert.Equal(t, 7, *p)

		*p = 9
		assert.Equal(t, 7, opt.Unwrap())
	})
}

func TestOption_OkOr(t *testing.T) {
	errMissing := errors.New("missing")

	tests := []struct {
		name    string
		give    option.Option[int]
		want    int
		wantErr error
	}{
		{"none", option.None[int](), 0, errMissing},
		{"some", option.Some(7), 7, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.give.OkOr(errMissing)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}

	t.Run("nil error", func(t *testing.T) {
		got, err := option.None[int]().OkOr(nil)
		assert.Equal(t, 0, got)
		assert.NoError(t, err)
	})
}

func TestOption_OkOrElse(t *testing.T) {
	errMissing := errors.New("missing")

	tests := []struct {
		name      string
		give      option.Option[int]
		want      int
		wantErr   error
		wantCalls int
	}{
		{"none", option.None[int](), 0, errMissing, 1},
		{"some", option.Some(7), 7, nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got, err := tt.give.OkOrElse(func() error {
				calls++
				return errMissing
			})
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}

func ExampleOption_Get() {
	if v, ok := option.Some(7).Get(); ok {
		fmt.Println(v)
	}

	if _, ok := option.None[int]().Get(); !ok {
		fmt.Println("none")
	}

	// Output:
	// 7
	// none
}

func ExampleOption_Ptr() {
	p := option.Some(7).Ptr()
	fmt.Println(*p)

	fmt.Println(option.None[int]().Ptr() == nil)

	// Output:
	// 7
	// true
}

func ExampleOption_OkOr() {
	fmt.Println(option.Some(7).OkOr(errors.New("missing")))
	fmt.Println(option.None[int]().OkOr(errors.New("missing")))

	// Output:
	// 7 <nil>
	// 0 missing
}

func ExampleOption_OkOrElse() {
	fmt.Println(option.Some(7).OkOrElse(func() error { return errors.New("missing") }))
	fmt.Println(option.None[int]().OkOrElse(func() error { return errors.New("missing") }))

	// Output:
	// 7 <nil>
	// 0 missing
}
