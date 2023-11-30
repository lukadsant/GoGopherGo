// experimento que utiliza arrays para gerar uma media
package main

import "fmt"

func mainslice() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

}
func mainappend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

}

func main() {

	mainappend()
}
