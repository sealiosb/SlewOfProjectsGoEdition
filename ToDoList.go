package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tasks []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To-Do List Application")
	fmt.Println("----------------------")

	for {
		fmt.Print("Enter a task (or type 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		task := scanner.Text()

		if task == "exit" {
			break
		}

		tasks = append(tasks, task)
		fmt.Println("Current To-Do List:")
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t)
		}
	}

	fmt.Println("Thank you for using the To-Do List Application!")
}
