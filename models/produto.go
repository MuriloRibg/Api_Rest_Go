package models

import (
	"Api_go/dataBase"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosProdutos() []Produto {
	db := dataBase.ConectaMySQL()
	selectProduto, err := db.Query("SELECT * FROM PRODUTOS ORDER BY ID;")
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

		p.Id = ID
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

	//Verificando se o insert vai retornar algum erro
	insertDB, err := db.Prepare("INSERT INTO PRODUTOS(NOME, DESCRICAO, PRECO, QUANTIDADE) VALUES(?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}

	//Adicionando os valores no insert
	_, err = insertDB.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeletaProduto(id string) {
	db := dataBase.ConectaMySQL()

	deleteDb, err := db.Prepare("DELETE FROM PRODUTOS WHERE ID = ?;")
	if err != nil {
		panic(err.Error())
	}

	_, err = deleteDb.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := dataBase.ConectaMySQL()

	selectProduto, err := db.Query("SELECT * FROM PRODUTOS WHERE ID = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for selectProduto.Next() {
		var ID, QUANTIDADE int
		var NOME, DESCRICAO string
		var PRECO float64

		err = selectProduto.Scan(&ID, &NOME, &DESCRICAO, &PRECO, &QUANTIDADE)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = ID
		produtoParaAtualizar.Nome = NOME
		produtoParaAtualizar.Preco = PRECO
		produtoParaAtualizar.Quantidade = QUANTIDADE
		produtoParaAtualizar.Descricao = DESCRICAO
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizarProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := dataBase.ConectaMySQL()

	updateProduto, err := db.Prepare("UPDATE PRODUTOS SET NOME = ?, DESCRICAO = ?, PRECO = ?, QUANTIDADE = ? WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = updateProduto.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
