package routes

import (
	"Api_go/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) //resebe o caminho e função q será executada.
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/update", controllers.Update)
}
