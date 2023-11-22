package main

import (
	"fmt"
	"time"
)

func jogador(s string, c chan string) {
	for i := 0; ; i++ {
		c <- s
	}
}

func main() {
	c := make(chan string)
	go jogador("ping", c)
	go jogador("pong", c)
	go func() {
		for {
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
	}()
	var entrada string
	fmt.Scanln(&entrada)
}
