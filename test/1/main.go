package main

import "fmt"

type s interface {
}
type a struct{}

func (d *a) q() {
	fmt.Print("1")
}
func (d *a) good() {
	fmt.Print("2")
}
func main() {
	var g s
	g = *new(a)
	var jj a
	jj = g.(a)
	jj.q()
}
