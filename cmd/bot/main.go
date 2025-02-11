package main

import (
	"fmt"

	. "github.com/ujooju/telegram_bot/internal/starter"
)

func main() {
	go Start()
	var command string
	for fmt.Scan(&command); ; fmt.Scan(&command) {
		if command == "exit" {
			return
		}
	}
}
