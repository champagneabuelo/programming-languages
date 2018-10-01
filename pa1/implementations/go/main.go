package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Graph : main data structure for holding tasks and dependencies
type Graph map[string][]string

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("error at file read: %v", err)
	}

	tasks := strings.Split(string(data), "\n")
	// removes the newline from EOF
	tasks = tasks[:len(tasks)-1]

	if len(tasks)%2 != 0 {
		fmt.Println("incomplete task list. exiting.")
		os.Exit(1)
	}

	graph := make(Graph)
	for len(tasks) > 0 {
		exists := graph[tasks[0]]
		if exists == nil {
			graph[tasks[0]] = []string{tasks[1]}
		} else {
			if !contains(graph[tasks[0]], tasks[1]) {
				graph[tasks[0]] = append(graph[tasks[0]], tasks[1])
			}
		}
		// remove the tasks from list
		tasks = append(tasks[:0], tasks[1:]...)
		tasks = append(tasks[:0], tasks[1:]...)
	}

	sorted := make([]string, 0)
	for index := range graph {
		traverseGraph(graph, sorted, index)
	}

	fmt.Println(strings.Join(sorted, ","))
}

func traverseGraph(graph Graph, sorted []string, task string) {
	if len(graph) == 0 {
		return
	}

	addToList(sorted, graph[task])
	delete(graph, task)

	if next := checkTasks(graph, task); next != "" {
		traverseGraph(graph, sorted, next)
	}

	return
}

func addToList(sorted []string, tasks []string) {
	// sort corresponding list alphabetically
	sort.Slice(tasks, func(i, j int) bool { return tasks[j] > tasks[i] })
	for _, task := range tasks {
		sorted = append(sorted, task)
	}
}

func checkTasks(graph Graph, task string) string {
	for index, list := range graph {
		if contains(list, task) {
			return index
		}
	}
	return ""
}

func contains(list []string, element string) bool {
	for _, current := range list {
		if current == element {
			return true
		}
	}
	return false
}
