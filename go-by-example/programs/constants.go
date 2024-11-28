package programs

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func printFloat(f float64) {
	fmt.Println(f)
}

func ConstantMainFunction() {
	fmt.Println(s)

	const n = 500000000 // A numeric constant has no type until it’s given one

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(n))

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	printFloat(d)
}
