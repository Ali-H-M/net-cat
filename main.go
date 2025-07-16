package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}

	funcs.HanlerServer(os.Args[1])
}
