package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello!")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello() // Goroutine
	fmt.Println("Concurrent World!")
	time.Sleep(1 * time.Second) // Donne du temps à la goroutine
}
