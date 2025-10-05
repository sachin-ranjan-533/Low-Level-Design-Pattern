package main

import "fmt"

type directory struct {
	name     string
	children []FileInterface
}

func (d *directory) createDirectory(name string) *directory {
	return &directory{name: name}

}

func (d *directory) ls() {
	fmt.Println(d.name)
	for _, child := range d.children {
		child.ls()
	}
}

func (d *directory) addChild(child FileInterface) {
	d.children = append(d.children, child)
}
