package sort_learn

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	count_arr := 10000
	var arr []int
	for i := 0; i < count_arr; i++ {
		arr = append(arr, rand.Int())
	}
	b.StartTimer()
	Sort(arr)
	b.StopTimer()
}

func BenchmarkBasicSort(b *testing.B) {
	count_arr := 10000
	var arr []int
	for i := 0; i < count_arr; i++ {
		arr = append(arr, rand.Int())
	}
	b.StartTimer()
	sort.Ints(arr)
	b.StopTimer()
}

func TestSort(t *testing.T) {
	cases := [...]struct {
		name   string
		input  []int
		result []int
	}{
		{
			"Массив уже отсортирован",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"Инверсия массива",
			[]int{6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"В массиве есть дубликаты",
			[]int{6, 5, 2, 3, 2, 1},
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			Sort(c.input)
			CompareArr(t, c.result, c.input)
		})
	}
}

func CompareArr(t *testing.T, expected, actual []int) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Errorf("У данных массивов разная длина %d против %d",
			len(expected), len(actual))
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Отличаются элементы на позиции %d. Значение %d против %d",
				i, expected[i], actual[i])
		}
	}
}
