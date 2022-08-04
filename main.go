package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	println("MAIN FUNCTION ENTERED")

	ann := boring("ann")
	jen := boring("jen")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-ann)
		fmt.Printf("You say: %q\n", <-jen)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}