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

	st := make(chan rpio.State)
	go waiting_pin(pin, st)
	for {
		select {
		case res := <-st:
			fmt.Println(res)
		default:
		}
	}
	close(st)
}

func waiting_pin(pin rpio.Pin, st chan rpio.State) {
	old := rpio.High
	for {
		res := pin.Read()
		if res != old {
			st <- res
			old = res
		}

		time.Sleep(1 * time.Second)
	}
}
