package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)


// Point型とそのメソッド
type Point struct {
	X float64
	Y float64
}

func NewPoint() *Point {
	p := Point{rand.Float64(), rand.Float64()}
	return &p
}

func (p *Point) Abs() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2)
}


func main() {
	n, err := strconv.Atoi(os.Args[1])
	var cnt int = 0

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		p := NewPoint()
		if p.Abs() < 1.0 {
			cnt++
		}
	}

	var pi float64 = 4.0 * float64(cnt) / float64(n)
	fmt.Println("pi is: ", pi)
}
