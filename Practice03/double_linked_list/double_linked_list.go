package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.previousNode = node
	}
	l.headNode = node
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

func (l *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (l *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (l *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
			break
		}
	}
	return lastNode
}

func (l *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node

	for node = l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}

	return nodeWith
}

func (l *LinkedList) IterateList() {
	for node := l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {

	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	//fmt.Println(linkedList.headNode.property)

	var node *Node
	node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println("NodeBetweenValues", node.property)

	linkedList.IterateList()

}
