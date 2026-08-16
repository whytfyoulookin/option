package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestCollect(t *testing.T) {
	tests := []struct {
		name      string
		give      []option.Option[int]
		wantValue []int
		wantNone  bool
	}{
		{"empty", nil, []int{}, false},
		{"all some", []option.Option[int]{option.Some(1), option.Some(2)}, []int{1, 2}, false},
		{"first none", []option.Option[int]{option.None[int](), option.Some(2)}, nil, true},
		{"middle none", []option.Option[int]{option.Some(1), option.None[int](), option.Some(3)}, nil, true},
		{"last none", []option.Option[int]{option.Some(1), option.None[int]()}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.Collect(tt.give)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestOption_Seq(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want []int
	}{
		{"none", option.None[int](), nil},
		{"some", option.Some(2), []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			for v := range tt.give.Seq() {
				got = append(got, v)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func ExampleCollect() {
	fmt.Println(
		option.Collect([]option.Option[int]{option.Some(1), option.Some(2)}).
			UnwrapOr([]int{-1}),
	)
	fmt.Println(
		option.Collect([]option.Option[int]{option.Some(1), option.None[int]()}).
			UnwrapOr([]int{-1}),
	)

	// Output:
	// [1 2]
	// [-1]
}

func ExampleOption_Seq() {
	for v := range option.Some(2).Seq() {
		fmt.Println(v)
	}

	for range option.None[int]().Seq() {
		fmt.Println("none")
	}

	// Output:
	// 2
}