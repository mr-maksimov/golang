package connect

import (
	"fmt"
	"log"
	"net/http"
	"learning/golang/sort"
	"learning/golang/shape"
)

// ListenServ слушать сервер http://localhost:8080/say
func ListenServ() {
	http.HandleFunc("/say", sayWeb)          // Устанавливаем роутер 
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// sayWeb
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
