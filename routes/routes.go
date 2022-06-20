package routes

import (
	"Api_go/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) //resebe o caminho e função q será executada.
	http.HandleFunc("/new", controllers.New)
}
