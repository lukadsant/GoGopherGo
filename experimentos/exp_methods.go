package main

import "fmt"

type triangle struct {
	size int
}

func (t *triangle) area() int {
	return t.size * t.size
}

func (t *triangle) perimeter() int {
	return t.size * 3
}

type square struct {
	size int
}

func (s *square) area() int {
	return s.size * s.size
}

func (s *square) perimeter() int {
	return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	t := triangle{3}
	s := square{4}
	t.doubleSize()
	fmt.Println("size:", t.size)
	fmt.Println("perimeter (triangle):", t.perimeter())
	fmt.Println("perimeter (square):", s.perimeter())
}
