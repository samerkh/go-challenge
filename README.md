# Users api using go

### Client and server

This is a simple client and server implementation using go, the server is a simple api that runs CRUD operations on users, the client is a simple client that uses CRUD operations on users from the server.

For the server i used chi router and logrus for logging, for the client i used the standard go http client.

### Data storage

The data is stored in memory using a map, so when the server is restarted, the data is lost.

### Goroutines and Channels

For demonstartion purposes, I used goroutines and channels inside `./server/internal/handlers/get_user.go`, I used `go fetchUser` to run the function asynchrously, and I used a channel to send the response back to the main thread.

### Project structure

the structure of the project is based on the standard go project layout, mentioned in this repo [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### Running the project

you need to run the server first before running the client

### Run the server

```bash
cd server
go run cmd/api/server.go
```

### Run the client

```bash
cd client
go run client.go
```
