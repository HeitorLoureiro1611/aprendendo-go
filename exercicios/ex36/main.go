// Crie um tipo "quadrado"
// Crie um tipo "círculo"
// Crie um método "área" para cada tipo que calcule e retorne a área da figura
// Área do círculo: π * (raio * raio)
// Área do quadrado: lado * lado
// Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// // Crie um valor de tipo "quadrado"
// Crie um valor de tipo "círculo"
// Use a função "info" para demonstrar a área do "quadrado"
// Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
)

const PI float32 = 3.1415

type figura interface {
	area() float32
}

type quadrado struct {
	lado   float32
	altura float32
}

type circulo struct {
	raio float32
}

func (q quadrado) area() float32 {
	return q.lado * q.altura
}

func (c circulo) area() float32 {
	return PI * (c.raio * c.raio)
}
func info(f figura) {
	fmt.Println("A área é:", f.area())
}
func main() {

	circ := circulo{raio: 7.0}
	quad := quadrado{lado: 4.0, altura: 4.0}

	fmt.Println("Informações do circulo:\nRaio:", circ.raio)
	info(circ)
	fmt.Println()
	fmt.Println("Informações do quadrado:")
	fmt.Println("Lado:", quad.lado, "\nAltura:", quad.altura)
	info(quad)
}
