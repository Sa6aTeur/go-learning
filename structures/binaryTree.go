package binary_tree

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func printTree(t *Tree) {
	if t == nil {
		return
	}
	printTree(t.Left)
	fmt.Print(t.Value, " ")
	printTree(t.Right)
}

func createTree(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Right: nil, Left: nil, Value: v}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func TestTreeLogic() {
	tree := createTree(10)
	fmt.Println("The value of the root of the tree is ", tree.Value)
	fmt.Println("firstTree")
	printTree(tree)
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	fmt.Println()
	fmt.Println("secondTree")
	printTree(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is ", tree.Value)
}
