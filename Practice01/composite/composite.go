package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (b *Branch) perform() {
	fmt.Println("Branch: " + b.name)
	for _, leaf := range b.leafs {
		leaf.perform()
	}

	for _, branch := range b.branches {
		branch.perform()
	}
}

// Branch class method add leaflet
func (b *Branch) add(leaf Leaflet) {
	b.leafs = append(b.leafs, leaf)
}

// Branch class method addBranch branch
func (b *Branch) addBranch(branch Branch) {
	b.branches = append(b.branches, branch)
}

// getLeaflets
func (b *Branch) getLeaflets() []Leaflet {
	return b.leafs
}

func main() {
	branch := &Branch{name: "branch 1"}
	leaflet1 := Leaflet{name: "leaf 1"}
	leaflet2 := Leaflet{name: "leaf 2"}
	branch2 := Branch{name: "branch 2"}
	branch.add(leaflet1)
	branch.add(leaflet2)
	branch.addBranch(branch2)
	branch.perform()
}
