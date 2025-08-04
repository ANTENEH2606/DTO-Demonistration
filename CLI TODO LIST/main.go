package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const todoFile = "todos.json"

type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type TodoList []Todo

func main() {
	todos, err := initializeTodoList()
	if err != nil {
		exitWithError("Error loading todos", err)
	}

	handleCommands(todos)
}

// Initialize and load todo list
func initializeTodoList() (TodoList, error) {
	var todos TodoList
	if err := todos.Load(); err != nil {
		return nil, err
	}
	return todos, nil
}

// Handle command-line arguments
func handleCommands(todos TodoList) {
	addFlag := flag.String("add", "", "Add a new task")
	listFlag := flag.Bool("list", false, "List all tasks")
	completeFlag := flag.Int("complete", 0, "Complete task by number")
	flag.Parse()

	switch {
	case *addFlag != "":
		handleAdd(todos, *addFlag)
	case *listFlag:
		todos.List()
	case *completeFlag > 0:
		handleComplete(todos, *completeFlag)
	default:
		printUsage()
	}
}

// Handle add operation
func handleAdd(todos TodoList, task string) {
	todos.Add(task)
	if err := todos.Save(); err != nil {
		exitWithError("Error saving todos", err)
	}
	fmt.Println("Task added:", task)
}

// Handle complete operation
func handleComplete(todos TodoList, index int) {
	if err := todos.Complete(index); err != nil {
		exitWithError("Error completing task", err)
	}
	if err := todos.Save(); err != nil {
		exitWithError("Error saving todos", err)
	}
	fmt.Println("Task completed:", index)
}

// Print usage instructions
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -add TASK        Add a new task")
	fmt.Println("  -list            List all tasks")
	fmt.Println("  -complete NUMBER Mark task NUMBER as completed")
}

// Error handling helper
func exitWithError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	os.Exit(1)
}

/* Existing TodoList methods remain the same */
func (t *TodoList) Load() error {
	data, err := os.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			*t = []Todo{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, t)
}

func (t *TodoList) Save() error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

func (t *TodoList) Add(task string) {
	*t = append(*t, Todo{Task: task, Completed: false})
}

func (t *TodoList) List() {
	if len(*t) == 0 {
		fmt.Println("Todo list is empty.")
		return
	}
	for i, todo := range *t {
		status := " "
		if todo.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Task)
	}
}

func (t *TodoList) Complete(index int) error {
	if index < 1 || index > len(*t) {
		return fmt.Errorf("invalid task number")
	}
	(*t)[index-1].Completed = true
	return nil
}
