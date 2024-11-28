package programs

import "fmt"

func VariablesMainFunction() {
	var a = "initial"
	fmt.Println(a)

	b, c := 1, "sumit"
	fmt.Println(b, c)

	var d bool
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var f string
	fmt.Println(f == "")
	// var p *int
	// fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference

	var pointer *int = new(int) // pointer to an int no error in this case
	fmt.Println(*pointer)
}
