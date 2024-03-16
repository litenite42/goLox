package main

import (
	"fmt"
	"os"
	"bufio"
	//"golox/loxerr"
	"golox/core"
)

type lox struct {
	had_error bool
}

func (l *lox) run_file(s string) {
	b, err := os.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b)

	str := string(b)

	fmt.Println(str)

	l.run(str)

	if l.had_error {
		os.Exit(65)
	}
}

func (l *lox) run_prompt() {
	fmt.Print(">> ")
	scnr := bufio.NewScanner(os.Stdin)

	for scnr.Scan() {
		inText := scnr.Text()

		fmt.Printf("In: %s\n", inText)

		l.run(inText)
	}

	if err := scnr.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)	
	}
}

func (l *lox) run(src string) {
	scnr := core.NewScanner(src)
	tokens := scnr.ScanTokens()

	for token := range tokens {
		fmt.Println(tokens[token])
	}
}

func main() {
	args := os.Args[1:]

	l := lox{had_error: false}

	if len(args) > 1 {
		fmt.Println("Usage: goLox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		l.run_file(args[0])
	} else {
		l.run_prompt()
	}
}
