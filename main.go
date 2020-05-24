package main

import (
	"fmt"
	"learning/golang/connect"
	"learning/golang/shape"
	"learning/golang/sort"
	"net/http"
)

func main() {
	connect.ListenServ(sayWeb)
}

func sayWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Привет!
`)
	arr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Fprintf(w, "Минимальный элемент: %d", shape.MinElement(arr))
	fmt.Fprintf(w, "Сортированный массив %d", sort.Sort(arr))
}
