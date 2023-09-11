package main

import "fmt"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (t *Tree) insert(m int) {
	if t != nil {
		if t.LeftNode == nil {
			t.LeftNode = &Tree{nil, m, nil}
		} else {
			if t.RightNode == nil {
				t.RightNode = &Tree{nil, m, nil}
			} else {
				if t.LeftNode != nil {
					t.LeftNode.insert(m)
				} else {
					t.RightNode.insert(m)
				}
			}
		}
	} else {
		t = &Tree{nil, m, nil}
	}
}

// print method for printing a Tree
func print(tree *Tree) {
	if tree != nil {

		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

func main() {
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.LeftNode.insert(7)
	print(tree)

}
