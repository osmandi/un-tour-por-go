package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // Tiene tipo Vertex
	q = &Vertex{1, 4} // Tiene tipo *Vertex
	r = Vertex{X: 1}  // Y:0 es es implícito. Este es la estructura literal
	s = Vertex{}      // X:0 e Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
