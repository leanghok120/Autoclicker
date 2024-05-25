package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Autoclicker (Go version)")

	// fmt.Print("Delay: ")
	// var delay int
	// fmt.Scanln(&delay)

	time.Sleep(5 * time.Second)

	for {
		robotgo.Click()
		time.Sleep(1 * time.Millisecond)
	}
}
