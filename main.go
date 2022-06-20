package main

import (
	"Api_go/routes"
	"log"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	log.Printf("Servidor rodando na porta %d", 8000)
	http.ListenAndServe(":8000", nil)
}
