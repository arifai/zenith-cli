package printer

import (
	"fmt"
	"log"
)

const (
	green  = "\033[1;32m"
	red    = "\033[1;31m"
	yellow = "\033[1;33m"
	reset  = "\033[0m"
)

func Green(message string, a ...any) {
	formatted := fmt.Sprintf(message, a...)
	log.Printf(green + formatted + reset)
}

func Red(message string, a ...any) {
	formatted := fmt.Sprintf(message, a...)
	log.Printf(red + formatted + reset)
}

func Yellow(message string, a ...any) {
	formatted := fmt.Sprintf(message, a...)
	log.Printf(yellow + formatted + reset)
}
