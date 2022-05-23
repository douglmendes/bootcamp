package main

import "fmt"

/*
Exercício 1 - Registro de estudantes

Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
Data e que tenha um método de detalhamento
*/

type aluno struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao string
}

func (a aluno) detalhamento() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("DataAdmissao:", a.DataDeAdmissao)
}

/*
Exercicio 2
Exercício 2 - Produtos de e-commerce
Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
produtos e devolver o valor do preço total.
As empresas têm 3 tipos de produtos:
- Pequeno, Médio e Grande.
Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
adicional pelo envio de $2500.

2

Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.
- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
custo adicional segundo o tipo de produto.

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
tenha)
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
e adicioná-lo a lista da loja
*/

type Produto interface {
	calcularCusto() float64
	getName() string
	getPreco() float64
	getTipoProduto() string
}

type produto struct {
	tipoProduto string
	nome        string
	preco       float64
}

func (p produto) calcularCusto() float64 {
	return p.preco
}

func novoProduto(p Produto) {
	fmt.Printf("Nome: %s\n", p.getName())
	fmt.Printf("Preço: R$%v\n", p.getPreco())
	fmt.Printf("Tipo do produto: %s\n", p.getTipoProduto())
	fmt.Println(p.calcularCusto())
}

func (p produto) getName() string {
	return p.nome
}

func (p produto) getPreco() float64 {
	return p.preco
}

func (p produto) getTipoProduto() string {
	return p.tipoProduto
}

func main() {

	//Exercicio 1
	fmt.Println("Exercicio 1\n")
	a1 := aluno{
		Nome:           "Douglas",
		Sobrenome:      "Mendes",
		RG:             "99999999",
		DataDeAdmissao: "01/01/1900",
	}
	a1.detalhamento()
	fmt.Println("\nFim exercicio  1")
	//fim exercicio 1
	p := produto{
		tipoProduto: "Pequeno",
		nome:        "mouse",
		preco:       10.5,
	}
	novoProduto(p)
	//Exercicio 2

}
