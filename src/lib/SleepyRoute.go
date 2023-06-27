package lib

import (
	"fmt"
	"time"
)

func SleepyRoute() {
	for i := 3; i > 0; i-- {
		fmt.Printf("Countdown: %d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Countdown: Go!")
}