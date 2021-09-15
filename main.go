package main

import "os"

var (
	message2 = "Morg" //quando eu declaro fora de uma func, eu preciso declarar e tipar
	//posso declarar quantas variáveis eu quiser aqui dentro
)

const (
	num0 = 255
	num2 = 0
)

func main() {
	//var message string = "Hello, world!" //não preciso usar var message string...
	//var message = "Hello, world!"
	//var message string
	//message = "Hello world"

	message := "Hello, world!" //declarando e atribuindo
	message += " Morgana"

	//para constantes:
	const num = 150 //pode ser declarada e nunca utilizada

	println(message)
	//posso declarar variáveis dentro do if, isso economiza espaço na memória, pois sai da pilha de declarações
	//a declaração e atribuição acaba dentro da condição, ela não retorna um valor, portanto
	if morgana := 200; morgana > 100 { //if sempre precisa ter um bloco
		println("Big")
	} else {
		println("Small")
	}

	switch num {
	//posso usar condições (num > x) nos cases
	case 100:
	case 150: //testa as duas condições e manda a mesma saída pra elas
		println("A hundred")
	case 200:
		println("Two hundred")
	default:
		println("Lalalal")
	}

	//não temos while em golang
	//caso quisermos um while true, usamos apenas um for {} ou for true{}
	for c := 200; c > 150; c-- {
		println("Morgana")
	}

	//os: pacote nativo
	//pego todos os argumentos do meu pacote e coloco em args
	args := os.Args
	/*for i := 0; i < len(args); i++ {
		println(args[i]) //vai printar todos os argumentos depois do go run main. go <argumentos>
	}*/
	// o primeiro argumento sempre será o nome do arquivo
	for i, args := range args { // for _, args := range args => ignora o primeiro
		println(i, " -> ", args)
	}

	//a, b := 1, 2
	//c, d, _, b = 1, 2, 5, 7

	//declaração de array
	//args := []int{1,2,3,4}

}
