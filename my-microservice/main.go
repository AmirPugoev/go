package main

import (
	"log"
	"my-microservice/database"
	"my-microservice/handlers"
	"my-microservice/kafka"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database.InitDB("user=username dbname=mydb sslmode=disable")

	// Инициализация Kafka
	kafka.InitProducer("localhost:9092")

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/messages", handlers.PostMessage).Methods("POST")
	r.HandleFunc("/statistics", handlers.GetStatistics).Methods("GET")

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
