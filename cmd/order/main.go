package main

import (
	"database/sql"
	"fmt"

	"github.com/alfredodiani/gointensivo-2023/internal/infra/database"
	"github.com/alfredodiani/gointensivo-2023/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil{
		panic(err)
	}

	defer db.Close() //espera tudo rodar e depois fecha a conexão
	orderRepository := database.NewOrderRepository(db)
	usCase := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID: "129",
		Price: 15.0,
		Tax: 2.0,
	}
	
	output, err := usCase.Execute(input)
	if err != nil{
		panic(err)
	}
	
	fmt.Println(output)
}
