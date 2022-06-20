package controllers

import (
	"Api_go/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*html")) //pegando as páginas.

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarTodosProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}
