/*


Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.

Ajude o professor com as seguintes questões:
	1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
	2. Corrigir as incorrectas.

	var 1nome string
	var sobrenome string
	var int idade
	1sobrenome := 6
	var licenca_para_dirigir = true
	var estatura da pessoa int
	quantidadeDeFilhos := 2



	var 1nome string ---- Nome da variavel não deve começar com número
		CORREÇÃO : var nome string

	var sobrenome string ---- Declarada corretamente

	var int idade ---- o tipo da variável deve vir depois do nome
		CORREÇÃO : var idade int

	1sobrenome := 6 ---- Nome da variável não deve começar com número
		CORREÇÃO: sobrenome:= 6

	var licenca_para_dirigir = true ---- Padrão de variaveis com mais de uma palavra deve ser camelCase
		e para declaração de variável é necessário que o tipo seja passado explicitamente ou usado o sinal de ":=""
		CORREÇÃO : licencaParaDigir := true ou var licencaParaDirigir boolean

	var estatura da pessoa int ---- Nome de variável não pode ter espaço
		CORREÇÃO : var estaturaDaPessoa int

	quantidadeDeFilhos := 2 ---- Declarada corretamente

*/