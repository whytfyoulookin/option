package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestOption_Insert(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		val  int
		want int
	}{
		{"none", option.None[int](), 1, 1},
		{"some", option.Some(2), 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			got := opt.Insert(tt.val)

			assert.Equal(t, tt.want, *got)
			assert.Equal(t, tt.want, opt.Unwrap())

			*got = 99
			assert.Equal(t, 99, opt.Unwrap())
		})
	}
}

func TestOption_GetOrInsert(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		val  int
		want int
	}{
		{"none", option.None[int](), 5, 5},
		{"some", option.Some(2), 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			got := opt.GetOrInsert(tt.val)

			assert.Equal(t, tt.want, *got)
			assert.Equal(t, tt.want, opt.Unwrap())

			*got = 99
			assert.Equal(t, 99, opt.Unwrap())
		})
	}
}

func TestOption_GetOrInsertDefault(t *testing.T) {
	tests := []struct {
		name string
		give option.Option[int]
		want int
	}{
		{"none", option.None[int](), 0},
		{"some", option.Some(2), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			got := opt.GetOrInsertDefault()

			assert.Equal(t, tt.want, *got)
			assert.Equal(t, tt.want, opt.Unwrap())

			*got = 99
			assert.Equal(t, 99, opt.Unwrap())
		})
	}
}

func TestOption_GetOrInsertWith(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		want      int
		wantCalls int
	}{
		{"none", option.None[int](), 5, 1},
		{"some", option.Some(2), 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			calls := 0
			got := opt.GetOrInsertWith(func() int {
				calls++
				return 5
			})

			assert.Equal(t, tt.wantCalls, calls)
			assert.Equal(t, tt.want, *got)
			assert.Equal(t, tt.want, opt.Unwrap())

			*got = 99
			assert.Equal(t, 99, opt.Unwrap())
		})
	}
}

func TestOption_Take(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue int
		wantNone  bool
	}{
		{"none", option.None[int](), 0, true},
		{"some", option.Some(2), 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			got := opt.Take()

			assert.True(t, opt.IsNone())

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestOption_TakeIf(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue int
		wantNone  bool
		leftValue int
		leftNone  bool
		wantCalls int
	}{
		{"none", option.None[int](), 0, true, 0, true, 0},
		{"some taken", option.Some(4), 4, false, 0, true, 1},
		{"some kept", option.Some(3), 0, true, 3, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			calls := 0
			got := opt.TakeIf(func(n int) bool {
				calls++
				return n%2 == 0
			})

			assert.Equal(t, tt.wantCalls, calls)

			if tt.wantNone {
				assert.True(t, got.IsNone())
			} else {
				assert.Equal(t, tt.wantValue, got.Unwrap())
			}

			if tt.leftNone {
				assert.True(t, opt.IsNone())
				return
			}

			assert.Equal(t, tt.leftValue, opt.Unwrap())
		})
	}
}

func TestOption_Replace(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		val       int
		wantValue int
		wantNone  bool
	}{
		{"none", option.None[int](), 3, 0, true},
		{"some", option.Some(2), 5, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.give
			got := opt.Replace(tt.val)

			assert.Equal(t, tt.val, opt.Unwrap())

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func ExampleOption_Insert() {
	opt := option.None[int]()
	p := opt.Insert(1)
	fmt.Println(*p)
	fmt.Println(opt.Unwrap())

	*p = 2
	fmt.Println(opt.Unwrap())

	// Output:
	// 1
	// 1
	// 2
}

func ExampleOption_GetOrInsert() {
	opt := option.None[int]()
	p := opt.GetOrInsert(5)
	fmt.Println(*p)

	*p = 7
	fmt.Println(opt.Unwrap())
	fmt.Println(*opt.GetOrInsert(9))

	// Output:
	// 5
	// 7
	// 7
}

func ExampleOption_GetOrInsertDefault() {
	opt := option.None[int]()
	p := opt.GetOrInsertDefault()
	fmt.Println(*p)

	*p = 7
	fmt.Println(opt.Unwrap())

	// Output:
	// 0
	// 7
}

func ExampleOption_GetOrInsertWith() {
	opt := option.None[int]()
	p := opt.GetOrInsertWith(func() int { return 5 })
	fmt.Println(*p)

	*p = 7
	fmt.Println(opt.Unwrap())

	// Output:
	// 5
	// 7
}

func ExampleOption_Take() {
	opt := option.Some(2)
	fmt.Println(opt.Take().UnwrapOr(-1))
	fmt.Println(opt.UnwrapOr(-1))

	opt = option.None[int]()
	fmt.Println(opt.Take().UnwrapOr(-1))
	fmt.Println(opt.UnwrapOr(-1))

	// Output:
	// 2
	// -1
	// -1
	// -1
}

func ExampleOption_TakeIf() {
	isEven := func(n int) bool { return n%2 == 0 }

	opt := option.Some(4)
	fmt.Println(opt.TakeIf(isEven).UnwrapOr(-1))
	fmt.Println(opt.UnwrapOr(-1))

	opt = option.Some(3)
	fmt.Println(opt.TakeIf(isEven).UnwrapOr(-1))
	fmt.Println(opt.UnwrapOr(-1))

	opt = option.None[int]()
	fmt.Println(opt.TakeIf(isEven).UnwrapOr(-1))

	// Output:
	// 4
	// -1
	// -1
	// 3
	// -1
}

func ExampleOption_Replace() {
	opt := option.Some(2)
	fmt.Println(opt.Replace(5).UnwrapOr(-1))
	fmt.Println(opt.Unwrap())

	opt = option.None[int]()
	fmt.Println(opt.Replace(3).UnwrapOr(-1))
	fmt.Println(opt.Unwrap())

	// Output:
	// 2
	// 5
	// -1
	// 3
}
