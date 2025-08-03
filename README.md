# NetCat (TCP Chat)

This project is a simplified version of the classic NetCat (`nc`) tool. It creates a TCP-based server-client chat system using Go, supporting multiple clients and real-time message broadcasting.


## Features

- Start a TCP chat server on a given port
- Accept multiple client connections concurrently
- Broadcast messages from one client to all others
- Logs all messages with timestamps
- Displays previous chat history to new users


## How to Run

1. Run the Server

```bash
go run main.go <port>
```

Example:

 ```
 go run main.go 8080
 ```

2. Open new Terminal window:

```
nc localhost 8080
```

### Each connected client will:

- Be prompted to enter a name

- Receive all previous messages

- Broadcast new messages to all users

## Future Work:

* Unit Testing

* Save server log to external file
