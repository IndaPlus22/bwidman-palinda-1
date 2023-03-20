package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		time.Sleep(delay)
		time := time.Now()
		fmt.Printf("The time is %d.%d.%d: %s\n", time.Hour(), time.Minute(), time.Second(), text)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 60*time.Second)

	select {}
}
