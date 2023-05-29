/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello World\n")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
