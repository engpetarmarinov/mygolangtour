package concurrency

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan<- int) {
	if t == nil {
		return
	}

	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan<- int) {
	if t == nil {
		return
	}

	walkRecursive(t.Right, ch)
	ch <- t.Value
	walkRecursive(t.Left, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chT1 := make(chan int)
	chT2 := make(chan int)
	go Walk(t1, chT1)
	go Walk(t2, chT2)
	for {
		v1, ok1 := <-chT1
		v2, ok2 := <-chT2
		if (ok1 != ok2) || (v1 != v2) {
			return false
		}

		if ok1 == false {
			break
		}
	}

	return true
}

func CompareEquivalentBinaryTreesTest() {
	fmt.Printf("Testing equivalent randomly generated binary trees are equal: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Testing non-equivalent randomly generated binary trees are equal: %v\n", Same(tree.New(1), tree.New(2)))
	fmt.Printf("Testing swapped non-equivalent randomly generated binary trees are equal: %v\n", Same(tree.New(2), tree.New(1)))
}
