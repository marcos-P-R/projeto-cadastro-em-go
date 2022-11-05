package model

import "projeto-atelie-di/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaPorTodosProdutos() []Produto {
	db := db.ConexaoBD()

	selectAllProduct, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectAllProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoBD()
	inserirDados := "INSERT INTO produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)"
	inserirNewProduct, err := db.Prepare(inserirDados)
	if err != nil {
		panic(err.Error())
	}

	inserirNewProduct.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConexaoBD()

	deletarQuery := "DELETE FROM produtos WHERE id=$1"

	deletarOProduto, err := db.Prepare(deletarQuery)
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConexaoBD()
	selectProduto := "SELECT * FROM produtos WHERE id=$1"

	produtoDoBanco, err := db.Query(selectProduto, id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConexaoBD()

	atualizar := "UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5"

	atualizarProduto, err := db.Prepare(atualizar)
	if err != nil {
		panic(err.Error())
	}

	atualizarProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()

}
