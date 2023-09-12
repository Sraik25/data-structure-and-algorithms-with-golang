package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	if l.headNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = node
}
func (l *LinkedList) AddToEnd(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	lastNode := l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (l *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node

	for node = l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (l *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)

	linkedList.AddToEnd(5)

	fmt.Println(linkedList.NodeWithValue(3))

	linkedList.AddAfter(1, 7)

	lastNode := linkedList.LastNode()
	fmt.Println("lastNode", lastNode)

	linkedList.IterateList()

	//fmt.Printf("%+v", linkedList.headNode.property)
}
