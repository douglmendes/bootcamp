package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Cumprimentar(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimentar(w, "mundo")
}

func main() {

	err := http.ListenAndServe(":5001", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}
	Cumprimentar(os.Stdout, "Eloide")
}
