/*Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
- Crie um tipo para um struct chamado "pessoa"
- Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
- Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
- Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
- Demonstre no seu código:
    - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
    - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/

package main

import "fmt"

type person struct {
	Name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Pode ser? Eita Bichiga!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	fmt.Println("Diga alguma expressão sua...")
	h.speak()
}

func main() {
	John := person{
		"Moises",
		31,
	}
	saySomething(&John)

	//Para utilizar o metodo diretamente teria q ser feito utilzando a seguinte notação
	fmt.Println("\n\n\n\n------------------>")
	(&John).speak()
	fmt.Println("<------------------")
}
