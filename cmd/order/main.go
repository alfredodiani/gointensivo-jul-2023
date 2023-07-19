package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/alfredodiani/gointensivo-2023/internal/infra/database"
	"github.com/alfredodiani/gointensivo-2023/internal/usecase"
	"github.com/alfredodiani/gointensivo-2023/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil{
		panic(err)
	}

	defer db.Close() //espera tudo rodar e depois fecha a conexão
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch,err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting RabbitMQ")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil{
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Message processed and saved on database:", output)
	}
}