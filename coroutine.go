package main

import (
	"fmt"
	"time"
)

func main(){
	
	canal := make(chan int)

	go func ()  {
		for i := 0; i < 10; i++ {
			fmt.Println("Jogou no canal - valor: ", i)
			canal <- i
		}
	}()

	go worker(canal, 1)
	worker(canal, 2)

}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal no worker ", workerID, " --- valor: ", <- canal, )
		time.Sleep(time.Second)
	}
}