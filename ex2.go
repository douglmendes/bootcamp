package main

import "fmt"

func main() {

	temp := 14
	umidade := 50
	pressAtmosferica := 1018

	fmt.Printf("===============Clima 19/05===============\nTemperatura :%v\nUmidade:%v\nPressão Atmosférica:%v mb\n\n", temp, umidade, pressAtmosferica)
	fmt.Printf("===============Tipos das variáveis===============\ntemp :%T\numidade:%T\npressAtmosferica:%T", temp, umidade, pressAtmosferica)

}
