package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var tasks = []string{"go", "python"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
		printTasks(tasks, w)
	})

	log.Printf("Server Start on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func printTasks(taskItems []string, w http.ResponseWriter) {
	for index, task := range taskItems {
		fmt.Fprintln(w, index+1, task)
	}
}
