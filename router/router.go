package router

import (
	"encoding/json"
	"fmt"
	"learning/golang/params"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type responseSort struct {
	Sort []int `json:"sort"`
}

// Router Отправка формата json
func Router() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", genGet)
	//http.Serve(autocert.NewListener("example.com"), r)
	http.ListenAndServe(":3000", r)
}

func genGet(w http.ResponseWriter, r *http.Request) {
	responseSort := responseSort{params.Array(10)}
	text, err := json.Marshal(responseSort)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(text)
}
