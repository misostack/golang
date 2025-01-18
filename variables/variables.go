package variables

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func ExampleDataTypes() {
	// boolean
	var isValid bool
	fmt.Println(isValid)
	// numbers : integer, float, complex
	var a, b int = 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Printf("sqrt(%v^2 + %v^2) = %v\n", a, b, c)
	var f float64 = 3.14
	fmt.Printf("f=%v\n", f)
	// go does not have built-in decimal
	// why decimal is important, let's check this example
	var x, y, z float64 = 0.1, 0.2, 0.3
	var xPlusY = x + y
	fmt.Printf("%v + %v = %v vs %v\n", x, y, xPlusY, z)
	// no worry we have a popular package can solve this problem
	var dx, dy = decimal.NewFromFloat(x), decimal.NewFromFloat(y)
	var dz = decimal.NewFromFloat(z)
	dXPlusY := dx.Add(dy)
	fmt.Printf("%v + %v = %v vs %v\n", dx, dy, dXPlusY, dz)

	// string
	var s string = "Đây là chuỗi utf8"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	for i, r := range s {
		fmt.Printf("%v: %c\n", i, r)
	}
}
