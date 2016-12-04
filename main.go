package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("Hello, world!")

	rpio.Open()
	to_go := rpio.Pin(9)
	to_go.Input()
	do_not_go := rpio.Pin(10)
	do_not_go.Input()

	cg := make(chan rpio.State)
	cn := make(chan rpio.State)
	go waiting_pin(to_go, cg)
	go waiting_pin(do_not_go, cn)
	for {
		select {
		case res := <-cg:
			if res == rpio.High {
				fmt.Println("I'm going to izakaya.")
			}
		case res := <-cn:
			if res == rpio.High {
				fmt.Println("I will not go to izakaya.")
			}
		default:
		}
	}
	close(cg)
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
