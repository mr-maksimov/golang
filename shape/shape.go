package shape

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Circle - круг
type Circle struct {
	x, y, r float64
}

//Rectangle - прямоугольник
type Rectangle struct {
	x1, y1, x2, y2 float64
}

//Shape - фигура
type Shape interface {
	area() float64
	perimetr() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimetr(shapes ...Shape) float64 {
	var perimetr float64
	for _, s := range shapes {
		perimetr += s.perimetr()
	}
	return perimetr
}

func (r *Rectangle) area() float64 {
	lenght, width := distance(r.x1, r.y1, r.x2, r.y2)
	return lenght * width
}

func (c *Circle) area() float64 {
	c.r = 5
	return math.Pi * c.r * c.r
}

func (r *Rectangle) perimetr() float64 {
	lenght, width := distance(r.x1, r.y1, r.x2, r.y2)
	return (lenght + width) * 2
}

func (c *Circle) perimetr() float64 {
	c.r = 5
	return 2 * math.Pi * c.r
}

func distance(x1, y1, x2, y2 float64) (float64, float64) {
	a := x2 - x1
	b := y2 - y1
	return a, b
}

func zero(x *int) {
	*x = 0
}

func swap(x *int, y *int) {
	var temp = *x
	*x = *y
	*y = temp
}

func fromFaringateToСelsius(Faringate float64) float64 {
	return (Faringate - 32) * 5 / 9
}

func fromFutToMeter(fut float64) float64 {
	return fut * 0.3048
}

func numberDeleteThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// MinElement Получить минимальный элемент массива
func MinElement(arr []int) int {
	minVal := 0
	for _, val := range arr {
		if minVal == 0 || minVal > val {
			minVal = val
		}
	}
	return minVal
}
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
