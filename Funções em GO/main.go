package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

/*
Exercício 1 - Impostos de salário

Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.
*/
func salaryTax(salary float64) float64 {
	var tax float64
	switch salary {
	case 50000:
		tax = float64(math.Round(salary * 0.15))
	case 150000:
		tax = float64(math.Round(salary * 0.17))
	}
	return tax
}

/*
Exercício 2 - Calcular média
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/
func gradeAverage(grades []float32) (float32, error) {
	var sum float32
	for _, g := range grades {
		if g > 0 {
			sum += g
		} else {
			return 0, errors.New("Existem notas com valores negativos")
		}
	}
	average := sum / float32(len(grades))
	return average, nil
}

/*
Exercício 3 - Calcular salário

Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria do funcionário.
Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais
*/

func salaryCalculator(categoryWorked map[string]int) {
	var salary float32
	categoryTab := make(map[string]float32)
	categoryTab["c"] = 1000
	categoryTab["b"] = 1500
	categoryTab["a"] = 3000
	for category, workedHours := range categoryWorked {
		if category == "a" {
			if workedHours > 160 {
				salary = (categoryTab["a"] * float32(workedHours)) + ((categoryTab["a"] * float32(workedHours)) * 0.5)
				fmt.Printf("Esse funcionário é da categoria %s e trabalhou mais de 160h mensais, seu salário é de R$%v\n\n", category, salary)
			} else {
				salary = categoryTab["a"] * float32(workedHours)
				fmt.Printf("Esse funcionário é da categoria %s, seu salário é de R$%v\n\n", category, salary)
			}
		} else if category == "b" {
			if workedHours > 160 {
				salary = (categoryTab["b"] * float32(workedHours)) + ((categoryTab["b"] * float32(workedHours)) * 0.2)
				fmt.Printf("Esse funcionário é da categoria %s e trabalhou mais de 160h mensais, seu salário é de R$%v\n\n", category, salary)
			} else {
				salary = categoryTab["b"] * float32(workedHours)
				fmt.Printf("Esse funcionário é da categoria %s, seu salário é de R$%v\n\n", category, salary)
			}
		} else if category == "c" {
			salary = categoryTab["c"] * float32(workedHours)
			fmt.Printf("Esse funcionário é da categoria %s, seu salário é de R$%v\n\n", category, salary)
		}
	}
}

/*
Exercício 4 - Cálculo de estatísticas

Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior
Exemplo:

const (
minimum = "minimum"
average = "average"
maximum = "maximum"
)

...
minhaFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...
minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func opMinimum(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
func opAverage(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	average := sum / len(values)
	return average
}
func opMaximum(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func calcType(calc string, values ...int) int {
	switch calc {
	case Minimum:
		return calcReturn(values, opMinimum)
	case Average:
		return calcReturn(values, opAverage)
	case Maximum:
		return calcReturn(values, opMaximum)
	}
	return 0
}

func calcReturn(values []int, calc func(values ...int) int) int {
	var result int
	for i, value := range values {
		if i == 0 {
			result = value
		} else {
			result = calc(result, value)
		}
	}
	return result
}

/*
Exercício 5 - Cálculo da quantidade de alimento
Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
Exemplo:

const (
dog = "dog"
cat = "cat"
)

...
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animalDog(quantity int) float64 {
	return float64(quantity) * 10
}

func animalCat(quantity int) float64 {
	return float64(quantity) * 5
}

func animalHamster(quantity int) float64 {
	return float64(quantity) * 0.25
}

func animalTarantula(quantity int) float64 {
	return float64(quantity) * 0.15
}

func Animal(animal string) (func(quantity int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("Animal inválido.")
	}
}

func qtdAlimento() {
	dogFn, err := Animal(dog)
	catFn, err := Animal(cat)
	hamsterFn, err := Animal(hamster)
	tarantulaFn, err := Animal(tarantula)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		var amount float64
		amount += dogFn(5)
		fmt.Println(amount)
		amount += catFn(8)
		fmt.Println(amount)
		amount += hamsterFn(4)
		fmt.Println(amount)
		amount += tarantulaFn(10)
		fmt.Println(amount)
	}

}

func main() {

	//EXERCICIO 1
	fmt.Println("EXERCÍCIO 1")
	fmt.Printf("O salário de R$50000 tem um imposto de R$%v e o salário de R$150000 tem um imposto de R$%v", salaryTax(50000), salaryTax(150000))
	fmt.Println(" \n-----------------------")
	//EXERCICIO 2
	fmt.Println("EXERCÍCIO 2")
	grades := []float32{12.0, 13.5, 15}
	average, err := gradeAverage(grades)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A media das notas é %v", average)
	fmt.Println(" \n-----------------------")
	//EXERCICIO 3
	fmt.Println("EXERCÍCIO 3")
	categoryWorked := make(map[string]int)
	categoryWorked["c"] = 162
	categoryWorked["b"] = 176
	categoryWorked["a"] = 172
	salaryCalculator(categoryWorked)
	fmt.Println(" \n-----------------------")
	//EXERCICIO 4
	fmt.Println("EXERCÍCIO 4")
	fmt.Println(calcType(Average, 5, 5, 10, 15))
	fmt.Println(" \n-----------------------")
	fmt.Println("EXERCÍCIO 5")
	qtdAlimento()
	fmt.Println(" \n-----------------------")

}
