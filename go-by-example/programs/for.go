package programs

import "fmt"

func ForMainFunction() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// for i := range 3 { // This only works in latest go version >= 1.22
	// 	fmt.Println("range", i)
	// }

	for {
		fmt.Println("loop")
		break
	}

	// for n := range 6 { // This only works in latest go version >= 1.22
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	for i, v := range [5]int{2, 3, 4, 6, 6} {
		fmt.Println(i, v)
	}
}
