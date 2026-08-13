package option_test

import (
	"fmt"

	"github.com/whytfyoulookin/option"
)

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
