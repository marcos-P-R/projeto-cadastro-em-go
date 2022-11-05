package controller

import (
	"log"
	"net/http"
	"projeto-atelie-di/model"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := model.BuscaPorTodosProdutos()
	temp.ExecuteTemplate(w, "index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Panic()
	}
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro na converção do preço: ", err)
	}

	quantidadeConvertida, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro na converção da quantidade: ", err)
	}

	model.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	model.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := model.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Panic()
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	idConverter, err := strconv.Atoi(id)
	if err != nil {
		log.Println("erro na converção do ID para int: ", err)
	}
	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("erro na converção do preço para Float", err)
	}

	quantidadeConvertida, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("erro na converção do quantidade para int: ", err)
	}

	model.AtualizarProduto(idConverter, nome, descricao, precoConvertido, quantidadeConvertida)

	http.Redirect(w, r, "/", 301)
}
