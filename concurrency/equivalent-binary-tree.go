package concurrency

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var ans bool = true
	for i := 0; i < 10; i++ {
		x := <-ch1
		y := <-ch2
		fmt.Println(x, y)
		if x != y {
			ans = false
		}
	}
	return ans
}

func MainTreeProblem() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(Same(t1, t2))
}
