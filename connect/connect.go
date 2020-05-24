package connect

import (
	"log"
	"net/http"
)

// ListenServ слушать сервер
func ListenServ(sayWeb func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/say", sayWeb)          // Устанавливаем роутер http://localhost:8080/say
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
