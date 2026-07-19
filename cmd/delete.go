package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "删除指定的待办事项",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		todos, _ := loadTodos()
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		saveTodos(todos)
		fmt.Printf("已删除 (ID: %d)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
