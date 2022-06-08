package main

import "errors"

var ErrNaoEncontrado = errors.New("Não foi possível encontrar a palavra que você procura")

type Dicionario map[string]string

func Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) {
	d[palavra] = definicao
}
