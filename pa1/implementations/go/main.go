package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Graph : main data structure for holding tasks and dependencies
type Graph map[string][]string

func main() {
	tasks, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("error at file read: %v", err)
	}
	fmt.Printf("from file: %s", tasks)
}
