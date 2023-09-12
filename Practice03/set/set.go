package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

func (s *Set) AddElement(element int) {
	if !s.ContainsElement(element) {
		s.integerMap[element] = true
	}
}

func (s *Set) DeleteElement(element int) {
	delete(s.integerMap, element)
}

func (s *Set) ContainsElement(element int) bool {
	_, exists := s.integerMap[element]
	return exists
}

func (s *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()
	for value, _ := range s.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (s *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()
	for value, _ := range s.integerMap {
		unionSet.AddElement(value)
	}
	for value, _ := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))
	var anotherSet *Set
	anotherSet = &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(set.Intersect(anotherSet))
	fmt.Println(set.Union(anotherSet))
}
