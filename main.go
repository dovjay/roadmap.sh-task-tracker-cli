package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"` // "todo", "in-progress", "done"
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(fileName)
	if err == nil {
		_ = json.Unmarshal(file, &tasks)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	return os.WriteFile(fileName, data, 0644)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskcli <command> [args]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]
	tasks, _ := loadTasks()

	switch command {
	case "add":
		title := strings.Join(args, " ")
		now := time.Now()
		task := Task{
			ID:        nextID(tasks),
			Title:     title,
			Status:    "todo",
			CreatedAt: now,
			UpdatedAt: now,
		}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("Task added:", title)

	default:
		fmt.Println("Unknown command")
	}
}
