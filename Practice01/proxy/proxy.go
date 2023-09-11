package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct{}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

type VirtualProxy struct {
	realObject *RealObject
}

func (v *VirtualProxy) performAction() {
	if v.realObject == nil {
		v.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	v.realObject.performAction()
}

func main() {
	object := VirtualProxy{}
	object.performAction()
}
