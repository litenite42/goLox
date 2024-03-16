package loxerr

import "fmt"

func Error(line int, msg string) {
	Report(line, "", msg)
}

func Report(line int, where, msg string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, msg)
}
