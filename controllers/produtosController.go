package controllers

import (
	"Api_go/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*html")) //pegando as páginas.

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarTodosProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("input-nome")
		descricao := r.FormValue("input-descricao")
		preco := r.FormValue("input-preco")
		quantidade := r.FormValue("input-quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err.Error())
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err.Error())
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/", 201)
}
