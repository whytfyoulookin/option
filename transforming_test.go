package option_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wlrgo/option"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[option.Option[int]]
		wantValue int
		wantNone  bool
	}{
		{"some some", option.Some(option.Some(6)), 6, false},
		{"some none", option.Some(option.None[int]()), 0, true},
		{"none", option.None[option.Option[int]](), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := option.Flatten(tt.give)

			if tt.wantNone {
				assert.True(t, got.IsNone())
				return
			}

			assert.Equal(t, tt.wantValue, got.Unwrap())
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[string]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"some", option.Some("foo"), 3, false, 1},
		{"none", option.None[string](), 0, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.Map(tt.give, func(v string) int {
				calls++
				return len(v)
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

func TestMapOr(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[string]
		want      int
		wantCalls int
	}{
		{"some", option.Some("foo"), 3, 1},
		{"none", option.None[string](), 42, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.MapOr(tt.give, 42, func(v string) int {
				calls++
				return len(v)
			})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}

func TestMapOrDefault(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[string]
		want      int
		wantCalls int
	}{
		{"some", option.Some("hi"), 2, 1},
		{"none", option.None[string](), 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.MapOrDefault(tt.give, func(v string) int {
				calls++
				return len(v)
			})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantCalls, calls)
		})
	}
}

func TestMapOrElse(t *testing.T) {
	tests := []struct {
		name             string
		give             option.Option[string]
		want             int
		wantMapCalls     int
		wantDefaultCalls int
	}{
		{"some", option.Some("foo"), 3, 1, 0},
		{"none", option.None[string](), 42, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapCalls := 0
			defaultCalls := 0
			got := option.MapOrElse(
				tt.give,
				func() int {
					defaultCalls++
					return 42
				},
				func(v string) int {
					mapCalls++
					return len(v)
				},
			)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantMapCalls, mapCalls)
			assert.Equal(t, tt.wantDefaultCalls, defaultCalls)
		})
	}
}

func TestZipWith(t *testing.T) {
	tests := []struct {
		name      string
		opt       option.Option[int]
		other     option.Option[string]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"some some", option.Some(1), option.Some("hi"), 3, false, 1},
		{"some none", option.Some(1), option.None[string](), 0, true, 0},
		{"none some", option.None[int](), option.Some("hi"), 0, true, 0},
		{"none none", option.None[int](), option.None[string](), 0, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := option.ZipWith(tt.opt, tt.other, func(n int, s string) int {
				calls++
				return n + len(s)
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

func TestOption_Filter(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"none", option.None[int](), 0, true, 0},
		{"some rejected", option.Some(3), 0, true, 1},
		{"some kept", option.Some(4), 4, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := tt.give.Filter(func(n int) bool {
				calls++
				return n%2 == 0
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

func TestOption_Inspect(t *testing.T) {
	tests := []struct {
		name      string
		give      option.Option[int]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"some", option.Some(2), 2, false, 1},
		{"none", option.None[int](), 0, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := tt.give.Inspect(func(v int) {
				calls++
				assert.Equal(t, tt.wantValue, v)
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

func TestOption_Reduce(t *testing.T) {
	tests := []struct {
		name      string
		opt       option.Option[int]
		other     option.Option[int]
		wantValue int
		wantNone  bool
		wantCalls int
	}{
		{"some some", option.Some(12), option.Some(17), 29, false, 1},
		{"some none", option.Some(12), option.None[int](), 12, false, 0},
		{"none some", option.None[int](), option.Some(17), 17, false, 0},
		{"none none", option.None[int](), option.None[int](), 0, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			got := tt.opt.Reduce(tt.other, func(a, b int) int {
				calls++
				return a + b
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

func ExampleFlatten() {
	x := option.Some(option.Some(6))
	fmt.Println(option.Flatten(x).UnwrapOr(-1))

	x = option.Some(option.None[int]())
	fmt.Println(option.Flatten(x).UnwrapOr(-1))

	x = option.None[option.Option[int]]()
	fmt.Println(option.Flatten(x).UnwrapOr(-1))

	// Output:
	// 6
	// -1
	// -1
}

func ExampleMap() {
	x := option.Some("Hello, World!")
	fmt.Println(option.Map(x, func(v string) int { return len(v) }).UnwrapOr(-1))

	y := option.None[string]()
	fmt.Println(option.Map(y, func(v string) int { return len(v) }).UnwrapOr(-1))

	// Output:
	// 13
	// -1
}

func ExampleMapOr() {
	x := option.Some("foo")
	fmt.Println(option.MapOr(x, 42, func(v string) int { return len(v) }))

	x = option.None[string]()
	fmt.Println(option.MapOr(x, 42, func(v string) int { return len(v) }))

	// Output:
	// 3
	// 42
}

func ExampleMapOrDefault() {
	x := option.Some("hi")
	y := option.None[string]()

	fmt.Println(option.MapOrDefault(x, func(v string) int { return len(v) }))
	fmt.Println(option.MapOrDefault(y, func(v string) int { return len(v) }))

	// Output:
	// 2
	// 0
}

func ExampleMapOrElse() {
	i := 21

	x := option.Some("foo")
	fmt.Println(
		option.MapOrElse(
			x,
			func() int { return 2 * i },
			func(v string) int { return len(v) },
		),
	)

	x = option.None[string]()
	fmt.Println(
		option.MapOrElse(
			x,
			func() int { return 2 * i },
			func(v string) int { return len(v) },
		),
	)

	// Output:
	// 3
	// 42
}

func ExampleZipWith() {
	x := option.Some(1)
	y := option.Some("hi")
	fmt.Println(option.ZipWith(x, y, func(n int, s string) int { return n + len(s) }).UnwrapOr(-1))
	fmt.Println(option.ZipWith(x, option.None[string](), func(n int, s string) int { return n + len(s) }).UnwrapOr(-1))

	// Output:
	// 3
	// -1
}

func ExampleOption_Filter() {
	isEven := func(n int) bool {
		return n%2 == 0
	}

	fmt.Println(option.None[int]().Filter(isEven).UnwrapOr(-1))
	fmt.Println(option.Some(3).Filter(isEven).UnwrapOr(-1))
	fmt.Println(option.Some(4).Filter(isEven).UnwrapOr(-1))

	// Output:
	// -1
	// -1
	// 4
}

func ExampleOption_Inspect() {
	x := option.Some(2).Inspect(func(v int) { fmt.Println("got:", v) })

	fmt.Println(x.Unwrap())

	option.None[int]().Inspect(func(v int) { fmt.Println("got:", v) })

	// Output:
	// got: 2
	// 2
}

func ExampleOption_Reduce() {
	add := func(a, b int) int { return a + b }

	fmt.Println(option.Some(12).Reduce(option.Some(17), add).UnwrapOr(-1))
	fmt.Println(option.Some(12).Reduce(option.None[int](), add).UnwrapOr(-1))
	fmt.Println(option.None[int]().Reduce(option.Some(17), add).UnwrapOr(-1))
	fmt.Println(option.None[int]().Reduce(option.None[int](), add).UnwrapOr(-1))

	// Output:
	// 29
	// 12
	// 17
	// -1
}
