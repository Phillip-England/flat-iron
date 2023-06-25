package lib

import (
	"fmt"
	"os"
)

func Debug(message string) {
	if os.Getenv("DEBUG") == "TRUE" {
		fmt.Println(message)
	}
}