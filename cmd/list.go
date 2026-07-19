package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有待办事项",
	Run: func(cmd *cobra.Command, args []string) {
		todos, _ := loadTodos()
		printTodos(todos)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
