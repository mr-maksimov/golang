package main

import (
	"fmt"
	"math"
	"math/rand"
	"test_learning/sort_learn"
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

func main() {
	/*
		c := Circle{1, 3, 2}
		r := Rectangle{1, 1, 10, 5}

		fmt.Println("Площадь круга =", c.area())
		fmt.Printf("Площадь прямоугольника = %v\n", r.area())
		s := fmt.Sprintf("Общая площадь фигур = %v\n", totalArea(&c, &r))
		io.WriteString(os.Stdout, s)
		fmt.Println("Общий периметр фигур =", totalPerimetr(&c, &r))
		zero(&x)
		fmt.Println(x)
		fmt.Println(len("dsada"))
		var x = 1
		var y = 2
		swap(&x, &y)
		fmt.Println(x, y)
		fmt.Println(fromFaringateToСelsius(86))
		fmt.Println(fromFutToMeter(30))
		numberDeleteThree()
		minElement()
		xs := []int{1, 2, 3}
		fmt.Println(add(xs...))
		fmt.Println(factorial(2))
		fmt.Println(fib(7))
	*/
	/*for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)*/
	arr := []int{3, 1, 22, 2, 11, 15}
	sort_learn.Sort(arr)
	fmt.Println(arr)
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

func minElement() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minVal := 0
	for _, val := range x {
		if minVal == 0 || minVal > val {
			minVal = val
		}
	}
	fmt.Println(minVal)
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
