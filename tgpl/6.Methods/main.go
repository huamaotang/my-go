package main

import (
	"fmt"
)

type A map[string]string

func (a A) F1() {
	fmt.Println("a")
}

type PP struct {
	X, Y int
}

func (p PP) Add() int {
	return p.X + p.Y
}

func (p *PP) Double() {
	p.X *= 2
	p.Y *= 2
}

func (p PP) AddByParam(x int) int {
	return p.X + p.Y + x
}

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	a := &PP{1,2}
	a.Double()
	fmt.Println(a.Add())

	b := &PP{
		X: 1,
		Y: 2,
	}
	add := b.Add // 方法值
	addBy1 := b.AddByParam
	double := b.Double
	double()
	double()
	fmt.Println(add(), addBy1(1), b)

	c := &PP{
		X: 1,
		Y: 2,
	}
	addExpression := PP.Add // 方法表达式
	addBy := PP.AddByParam
	doubleExpression := (*PP).Double
	fmt.Println(addExpression(*c), addBy(*c, 1))
	doubleExpression(c)
	fmt.Println(c)
}
