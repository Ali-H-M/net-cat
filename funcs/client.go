package funcs

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandleClient(conn net.Conn) { // net.Conn = the network connection between the server and the client
	defer conn.Close()

	fmt.Fprintln(conn, LinuxAscii)
	fmt.Fprint(conn, "[ENTER YOUR NAME]")

	nameScanner := bufio.NewScanner(conn)
	if !nameScanner.Scan() { // Waits for user input
		return
	}

	clientName := nameScanner.Text()

	client := &Client{
		Name: clientName,
		Conn: conn,
	}

	// Add client to map
	clientsMutex.Lock()
	clients[conn] = client
	clientsMutex.Unlock()

	welcomeMsg := fmt.Sprintf("%s has joined our chat...", client.Name)
	fmt.Println(welcomeMsg)      // Server log
	chanMessage <- welcomeMsg // For sending to all clients

	// Send chat history to the new client
	clientsMutex.Lock()
	for _, msg := range serverLogHistory {
		fmt.Fprintln(conn, msg)
	}
	clientsMutex.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // Runs as long as the client keeps sending data

		if scanner.Text() == "" {
			continue
		}

		client.message = scanner.Text()
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		clientMessage := fmt.Sprintf("[%s] [%s]: %s", timestamp, client.Name, client.message)

		fmt.Println(clientMessage)      // Server log
		chanMessage <- clientMessage // For sending to all clients
	}

	// Client disconnected
	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()

	leaveMsg := fmt.Sprintf("%s has left the chat", client.Name)
	fmt.Println(leaveMsg)
	chanMessage <- leaveMsg

}
