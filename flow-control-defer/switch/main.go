package main

import (
	"fmt"
	"runtime"
	"time"
)

// Note: In Go, switch does **not** fall through by default.
// It stops at the first matching case — no need for `break`.

func osSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // Short statement before switch
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func weekdaySwitch() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func timeSwitch() {
	t := time.Now()
	// switch with no condition = switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	osSwitch()
	weekdaySwitch()
	timeSwitch()
}
