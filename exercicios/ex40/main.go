// Callback: passe uma função como argumento a outra função.

package main

import (
	"fmt"
)

func oibondia() {
	fmt.Println("como vai você?")
}

// callback está aqui,
// recebe uma função como argumento e usa ela depois
func callback(f func()) {
	fmt.Println("Opa, tudo bem?")
	f()
}

func main() {

	callback(oibondia)

}
