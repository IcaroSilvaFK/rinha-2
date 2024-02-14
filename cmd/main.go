package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/rinha-backend-02/application/controllers"
	"github.com/IcaroSilvaFK/rinha-backend-02/application/database"
)

func main() {

	mx := http.NewServeMux()
	db := database.NewDatabaseConnection()

	tController := controllers.NewTransactionController(db)

	mx.HandleFunc("GET /heathy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})
	mx.HandleFunc("POST /clientes/{userId}/transacoes", tController.CreateNewTransaction)
	mx.HandleFunc("GET /clientes/{userId}/extrato", tController.ListTransactions)

	log.Fatal(http.ListenAndServe(":8080", mx))
	fmt.Println("Server running on port 8080")
}
