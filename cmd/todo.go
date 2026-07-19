package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const dataFile = "todos.json"

type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func loadTodos() ([]Todo, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func nextID(todos []Todo) int {
	maxID := 0
	for _, t := range todos {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func printTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("暂无待办事项")
		return
	}
	for _, t := range todos {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %3d. %s\n", status, t.ID, t.Text)
	}
}
