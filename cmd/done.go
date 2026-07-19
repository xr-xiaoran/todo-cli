package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [ID]",
	Short: "标记待办事项为已完成",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		todos, _ := loadTodos()
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}
		saveTodos(todos)
		fmt.Printf("已标记为完成 (ID: %d)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
