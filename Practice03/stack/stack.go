package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

func (e Element) String() string {
	return strconv.Itoa(e.elementValue)
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (s *Stack) New() {
	s.elements = make([]*Element, 0)
}

func (s *Stack) Push(element *Element) {
	s.elements = append(s.elements[:s.elementCount], element)
	s.elementCount = s.elementCount + 1
}

func (s *Stack) Pop() *Element {
	if s.elementCount == 0 {
		return nil
	}

	length := len(s.elements)
	element := s.elements[length-1]

	if length > 1 {
		s.elements = s.elements[:length-1]
	} else {
		s.elements = s.elements[0:]
	}
	s.elementCount = len(s.elements)
	return element
}

func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{5}
	var element3 *Element = &Element{7}
	var element4 *Element = &Element{9}
	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
