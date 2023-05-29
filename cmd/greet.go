/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string
var count int

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets a user",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			fmt.Printf("Hello %s\n", name)
		}
	},
}

func init() {
	greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
	greetCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of times to print the greeting")
	rootCmd.AddCommand(greetCmd)
}
