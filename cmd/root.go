package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

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

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command,args []string) {
		tasks,_:= service.ListTasks()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "ID\tTitle\tDone\tCreated")

		for _,t := range tasks{
			fmt.Fprintf(w, "%d\t%s\t%t\t%s\n", t.ID, t.Title, t.Completed, t.CreatedAt)
		}
	},
}

func timeAgo(t string )string {
	parsed,err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t
	}
	return time.Since(parsed).String()
}