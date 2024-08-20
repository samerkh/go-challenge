package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

func main() {
	var user = User{1, "John Doe"}
	var user2 = User{2, "Jack Smith"}

	fmt.Println("Adding users")
	addUser(user)
	addUser(user2)

	fmt.Println("Getting users")
	getUsers()

	fmt.Println("Getting user 1")
	getUser(1)

	fmt.Println("Updating user 1")
	updateUser(1, User{1, "Jane Doe"})

	fmt.Println("Getting users")
	getUsers()

	fmt.Println("Deleting user 2")
	deleteUser(2)

	fmt.Println("Getting users")
	getUsers()
}

func addUser(user User) {

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling user:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Received response:", string(body))
}

func getUsers() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Received response:", string(body))
}

func getUser(id uint) {
	resp, err := http.Get("http://localhost:8080/users/" + fmt.Sprint(id))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Received response:", string(body))
}

func updateUser(id uint, user User) {
	userJSON, err := json.Marshal(user)

	req, err := http.NewRequest("PUT", "http://localhost:8080/users/"+fmt.Sprint(id), bytes.NewBuffer(userJSON))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Received response:", string(body))
}

func deleteUser(id uint) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/users/"+fmt.Sprint(id), nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Received response:", string(body))
}
