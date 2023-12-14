package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Hello() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna olá + nome passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n ", name)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "Mundo", "flag para concatenar com Olá")

	return cmd
}

func GreetTimeOfDay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greet-time",
		Short: "Saudação com base no hórario.",
		Run: func(cmd *cobra.Command, args []string) {
			currentHour := time.Now().Hour()

			var greeting string
			if currentHour < 12 {
				greeting = "Bom Dia!"
			} else {
				greeting = "Boa Tarde!"
			}
			fmt.Println(greeting)
		},
	}
	return cmd
}
