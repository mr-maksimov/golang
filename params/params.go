package params

import "math/rand"

// Array генерация рандомного массива
func Array(countArr int) []int {
	var arr []int
	for i := 0; i < countArr; i++ {
		arr = append(arr, rand.Int())
	}
	return arr
}
