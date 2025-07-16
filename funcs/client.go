package funcs

import (
	"bufio"
	"fmt"
	"net"
)

func HandleClient(conn net.Conn) { // net.Conn = open TCP
	defer conn.Close()

	fmt.Println("New client joined:")

	fmt.Fprintln(conn, "Welcome to the chat!")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // runs as long as the client keeps sending data
		message := scanner.Text()
		fmt.Printf("New Messge: %s\n", message)
	}

	fmt.Println("Client disconnected.")
}
