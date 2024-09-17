package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TodoList Structure for storing tasks
type TodoList struct {
	tasks []string
}

//Load task list
func (t *TodoList) loadTasks(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No task file found, create a new task list")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t.tasks = append(t.tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading task", err)
	}
}

//Save task list to file
func (t *TodoList) saveTasks(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving task", err)
	}
	defer file.Close()

	for _, task := range t.tasks {
		file.WriteString(task + "\n")
	}
}

//add task
func (t *TodoList) addTask(task string) {
	t.tasks = append(t.tasks, task)
	fmt.Println("Task added", task)
}

//delete task
func (t *TodoList) deleteTask(idx int) {
	if idx < 0 || idx >= len(t.tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	fmt.Println("delete task: ", t.tasks[idx])
	t.tasks = append(t.tasks[:idx], t.tasks[idx+1:]...)
}

//display task list
func (t *TodoList) showTasks() {
	if len(t.tasks) == 0 {
		fmt.Println("There are currently no tasks available.")
		return
	}
	fmt.Println("To-do list:")
	for i, task := range t.tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
func main() {
	todoList := &TodoList{}
	todoList.loadTasks("tasks.txt")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nOptional command: add, delete, show, exit")
		fmt.Print("input command: ")
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "add":
			fmt.Print("please input task: ")
			scanner.Scan()
			task := scanner.Text()
			todoList.addTask(task)
		case "delete":
			todoList.showTasks()
			fmt.Print("Please enter the task number to be deleted: ")
			var idx int
			fmt.Scan(&idx)
			todoList.deleteTask(idx - 1)
		case "show":
			todoList.showTasks()
		case "exit":
			todoList.saveTasks("tasks.txt")
			fmt.Println("Task saved, exit program")
			return
		default:
			fmt.Println("invalid command. please re-enter")
		}
	}
}
