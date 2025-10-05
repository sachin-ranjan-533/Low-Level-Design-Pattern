package main

import "fmt"

type file struct {
	name string
}

func (f *file) ls() {
	fmt.Println(f.name)
}
