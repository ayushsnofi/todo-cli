package cmd

import (
	"fmt"
	"os"
	"todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"todo",
	Short: "CLI todo App",
}

func Execute(){
	if err:= rootCmd.Execute(); err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		service.AddTask(title)
		fmt.Println("Task added")
	},
}