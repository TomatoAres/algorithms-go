package main

import "github.com/golang-collections/collections/queue"
import "fmt"

func main() {
	q := queue.New()
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}
