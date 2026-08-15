package main

import (
	"fmt"
)

func main() {

	var tasks = []string{"go", "python"}

	printTasks(tasks)
}

func printTasks(taskItems []string) {
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}
