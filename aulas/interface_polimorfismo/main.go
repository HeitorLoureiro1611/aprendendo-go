package main

import (
	"fmt"
)

// declaração de interface, predefinição de metodos
// que uma struct pode declarar
type Animal interface {
	som() string
}

// declaração da struct que implementa a interface animal que
// permite usar o metodo som() caso seja declarado
type Quokka struct {
	Animal
}

type Camelo struct {
	Animal
}

// declaração de uma função de uma interface que retorna um tipo string
// apenas animais conseguem usar este metodo
func comer(Animal) string {
	return "eu estou Comendo!\n"
}

// polimorfismo: usando a mesma chamada de metodo com retornos diferentes
func (q Quokka) som() string {
	return ":D"
}

func (c Camelo) som() string {
	return "Guiererere"
}

func main() {

	delo := Quokka{}
	danilo := Camelo{}

	fmt.Println(danilo.som())
	fmt.Println(comer(danilo))

	fmt.Println(delo.som())
	fmt.Println(comer(delo))

}
