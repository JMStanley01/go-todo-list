package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todoFile = "tasks.json"

// Load tasks from the JSON file
func loadTasks() []Task {
	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		return []Task{} // Return an empty list if file doesn't exist
	}

	data, err := ioutil.ReadFile(todoFile)
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

// Save tasks to the JSON file
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	err = ioutil.WriteFile(todoFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// Add a new task to the list
func addTask(tasks []Task, taskName string) []Task {
	newTask := Task{Name: taskName, Done: false}
	tasks = append(tasks, newTask)
	fmt.Printf("Added task: %s\n", taskName)
	return tasks
}

// List all tasks with their status
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("[%d] %s - [%s]\n", i, task.Name, status)
	}
}

// Mark a task as complete by its index
func completeTask(tasks []Task, index int) []Task {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number!")
		return tasks
	}

	tasks[index].Done = true
	fmt.Printf("Marked task #%d as complete: %s\n", index, tasks[index].Name)
	return tasks
}

func main() {
	// Define command-line flags
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", -1, "Mark a task as complete")

	flag.Parse()

	// Load existing tasks from file
	tasks := loadTasks()

	// Handle commands
	if *add != "" {
		tasks = addTask(tasks, *add)
	} else if *list {
		listTasks(tasks)
	} else if *complete >= 0 {
		tasks = completeTask(tasks, *complete)
	}

	// Save updated tasks back to the file
	saveTasks(tasks)
}
