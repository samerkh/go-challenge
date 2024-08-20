# Users api using go

### Goroutines and Channels

For demonstartion purposes, I used goroutines and channels inside `./server/internal/handlers/get_user.go`, I used `go fetchUser` to run the function asynchrously, and I used a channel to send the response back to the main thread.

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
