package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

	case "list":
		if len(args) == 0 {
			fmt.Println("All tasks:")
			for _, t := range tasks {
				fmt.Printf("[%d] %s (%s)\n  Created: %s\n  Updated: %s\n", t.ID, t.Title, t.Status, t.CreatedAt.Format(time.RFC3339), t.UpdatedAt.Format(time.RFC3339))
			}
		} else if args[0] == "done" || args[0] == "in-progress" || args[0] == "todo" {
			fmt.Println("Filtered tasks:")
			for _, t := range tasks {
				if t.Status == args[0] {
					fmt.Printf("[%d] %s (%s)\n", t.ID, t.Title, t.Status)
				}
			}
		} else {
			fmt.Println("Invalid filter. Use: done, in-progress, or todo")
		}

	case "mark":
		if len(args) < 2 {
			fmt.Println("Usage: mark <id> <todo|in-progress|done>")
			return
		}
		id, _ := strconv.Atoi(args[0])
		status := args[1]
		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Status = status
				tasks[i].UpdatedAt = time.Now()
				break
			}
		}
		saveTasks(tasks)
		fmt.Println("Task status updated")

	case "update":
		if len(args) < 2 {
			fmt.Println("Usage: update <id> <new title>")
			return
		}
		id, _ := strconv.Atoi(args[0])
		title := strings.Join(args[1:], " ")
		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Title = title
				tasks[i].UpdatedAt = time.Now()
				break
			}
		}
		saveTasks(tasks)
		fmt.Println("Task updated")

	default:
		fmt.Println("Unknown command")
	}
}
