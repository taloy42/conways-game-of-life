package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Enter board size: ")
	var size int
	fmt.Scanln(&size)
	fmt.Print("Enter probability for cell to be alive (at the start): ")
	var prob float64
	fmt.Scanln(&prob)
	b := RandBoard(size, size, prob)
	var interval int64
	fmt.Print("Enter number of miliseconds to elapse between board updates [1000 for 1 second]: ")
	fmt.Scanln(&interval)
	interval *= 1000000
	for {
		fmt.Println(b.String())
		b.NextPhase()
		time.Sleep(time.Duration(interval))
	}
}
