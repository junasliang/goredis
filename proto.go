package main

import "fmt"

type Command struct {
	//???
}

func parseCommand(msg string) (Command, error) {
	t := msg[0]
	switch t {
	case '*':
		fmt.Println(msg)
	}
	return Command{}, nil
}
