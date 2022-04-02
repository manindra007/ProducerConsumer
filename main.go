package main

import (
	"fmt"
	ls "producerconsumer/listener"
	wr "producerconsumer/writer"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var pr, prod int
	fmt.Println("Enter the number of product")
	fmt.Scanln(&prod)
	fmt.Println("Enter Production Rate")
	fmt.Scanln(&pr)

	go wr.Writer(&wg, pr, prod)
	// time.Sleep(5 * time.Second)
	go ls.Listen(&wg, pr, prod)
	wg.Wait()
	fmt.Println("Process complete")
}
