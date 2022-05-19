package main

import "fmt"

/**
Exercício 1 - Letras de uma palavra
A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
	1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
		letras que ela contém.
	2. Em seguida, imprimi cada uma das letras.

**/

func ex1(palavra string) {
	qtdLetras := len(palavra)
	fmt.Printf("A palavra tem %v letras\n", qtdLetras)
	fmt.Println(string([]rune(palavra)))
}

/**

Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

**/

func ex2(idade int, empregado bool, tempoEmpregado float32, salario float32) {
	type Cliente struct {
		Idade          int
		Empregado      bool
		TempoEmpregado float32
		Salario        float32
	}

	c := Cliente{idade, empregado, tempoEmpregado, salario}
	if c.Empregado == true && c.Idade > 22 && c.TempoEmpregado > 1 && c.Salario < 100000 {
		fmt.Println("O cliente inserido pode solicitar um empréstimo")
	} else if c.Empregado == true && c.Idade > 22 && c.TempoEmpregado > 1 && c.Salario > 100000 {
		fmt.Println("O cliente inserido pode solicitar um empréstimo SEM JUROS")
	} else {
		fmt.Println("O cliente inserido NÃO pode solicitar um empréstimo")
	}
}

/**
Exercício 3 - A que mês corresponde?

Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.

**/

func ex3(mes int) {
	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	}
	/*
		Porque foi utilazado o SWITCH ?
		Em sua definição o switch é utilizado quando queremos comparar a mesma variável ou expressão com várias opções. Sendo assim como estamos comparando a mesma variável
		e verificando o mês representado sem comparação com nenhuma outra variável ou expressão foi utilizado o switch.
	*/
}

/**
Exercício 4 - Quantos anos tem...
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.
**/

func ex4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var maior int
	fmt.Printf("A idade de Benjamin é %v anos\n", employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			maior++
		}
	}
	fmt.Printf("\nExistem %v funcionários com mais de 21 anos\n", maior)

	employees["Federico"] = 25

	fmt.Printf("\nO funcionário Federico foi adicionadom , confira a nova tabela :\n")
	for funcionario, _ := range employees {
		fmt.Println(funcionario)

	}
	delete(employees, "Pedro")
	fmt.Println("\nO funcionário Pedro foi deletado do quadro de funcionários, confira a nova tabela : \n")
	for funcionario, _ := range employees {
		fmt.Println(funcionario)

	}

}

// Main

func main() {

	//Exericio 1
	fmt.Println("EXERCÍCIO 1")
	ex1("Douglas")
	fmt.Println(" -----------------------")
	//Exercicio 2
	fmt.Println("EXERCÍCIO 2")
	ex2(29, true, 4, 80000)
	fmt.Println(" -----------------------")
	//Exercicio 3
	fmt.Println("EXERCÍCIO 3")
	ex3(2)
	fmt.Println(" -----------------------")
	fmt.Println("EXERCÍCIO 4")
	ex4()
	fmt.Println(" -----------------------")
}
