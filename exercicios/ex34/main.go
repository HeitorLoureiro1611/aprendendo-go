// Utilize a declaração defer de maneira que demonstre que
// sua execução só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {
	x := 6
	defer fmt.Println("\nvalor inicial de x:", x)

	for i := range 10 {
		x += i
		fmt.Println(x)
	}
}
