// este experimento utiliza a função de ponteiros
package main

import "fmt"

func main() {
	firstName := "jj"
	updateName(firstName)
	fmt.Println(firstName)
}

func updateName(name string) {
	name = "david"
}

/*
package main

import "fmt"

func main() {
	firstName := "jj"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "david"
}
*/
