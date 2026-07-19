package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [待办内容]",
	Short: "添加新的待办事项",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		todos, _ := loadTodos()
		todo := Todo{
			ID:        nextID(todos),
			Text:      text,
			Done:      false,
			CreatedAt: time.Now(),
		}
		todos = append(todos, todo)
		saveTodos(todos)
		fmt.Printf("已添加: %s (ID: %d)\n", text, todo.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
