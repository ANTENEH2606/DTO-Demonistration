# Todo List Command-Line Application in Go

A simple command-line Todo List manager written in Go. It allows you to add, list, and mark tasks as completed. Todos are persisted in a JSON file (`todos.json`) stored locally.

## Features

- Add new tasks to your todo list
- List all tasks with their completion status
- Mark tasks as completed by their number
- Persistent storage of todos in a JSON file

## Todo Storage

- Todos are saved and loaded from the `todos.json` file in the program’s directory.
- Tasks have two fields: `task` (string) and `completed` (boolean).

## Error Handling

- The program checks for invalid task numbers when marking tasks completed.
- If the `todos.json` file does not exist, the app initializes with an empty todo list.
- Any file read/write errors display an error message and exit the program.

## Code Overview

- `main()` initializes the todo list and processes command-line arguments.
- Methods on `TodoList` manage loading, saving, adding, listing, and completing tasks.
- The program uses Go’s `flag` package for CLI argument parsing.
- JSON encoding/decoding is used for persisting data in `todos.json`.

## Usage

Run the program by writing "go run main.go" and add the first -add command-line flags:

- `-add "TASK"`: Add a new task to the list
- `-list`: List all current tasks with their status
- `-complete NUMBER`: Mark the task at the given number as completed


