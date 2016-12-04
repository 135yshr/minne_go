package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("Hello, world!")

	rpio.Open()
	pin := rpio.Pin(9)
	pin.Input()
	for {
		res := pin.Read()
		fmt.Println(res)

		time.Sleep(1 * time.Second)
	}
}
