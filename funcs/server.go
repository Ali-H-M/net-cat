package funcs

import (
	"fmt"
	"net"
)

func HanlerServer(port string) {

	// Accepts incoming TCP connections
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server is listening on port", port)

	for {
		conn, err := server.Accept() // Waits for a new clients
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go HandleClient(conn) // To make the server can handle multiple client at the same time
	}
}
