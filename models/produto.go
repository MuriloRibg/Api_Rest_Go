package models

import "Api_go/dataBase"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosProdutos() []Produto {
	db := dataBase.ConectaMySQL()
	selectProduto, err := db.Query("SELECT * FROM PRODUTOS;")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProduto.Next() {
		//mapeando oq retorna do banco de dados
		var ID, QUANTIDADE int
		var NOME, DESCRICAO string
		var PRECO float64

		//a ordem é importante
		err = selectProduto.Scan(&ID, &NOME, &DESCRICAO, &PRECO, &QUANTIDADE)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = NOME
		p.Preco = PRECO
		p.Quantidade = QUANTIDADE
		p.Descricao = DESCRICAO

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := dataBase.ConectaMySQL()
}
