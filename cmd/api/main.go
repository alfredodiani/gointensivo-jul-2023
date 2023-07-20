package main

import (
	"encoding/json"
	"net/http"

	"github.com/alfredodiani/gointensivo-2023/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	//usando chi como router - (poderia usar o echo no lugar do chi)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", Order) //trata a requisição usando a bibliotech go-chi
	http.ListenAndServe(":8888", r)
}

func Order(w http.ResponseWriter, r *http.Request){
	order, err := entity.NewOrder("158", 100, 15)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}