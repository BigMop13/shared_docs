package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type User struct {
	Username string
	Email    string
	Created  time.Time
}

type Document struct {
	ID       string
	Title    string
	Content  string
	Owner    string
	Editors  []string
	Created  time.Time
	Modified time.Time
}

type SharedDocsSystem struct {
	users     map[string]*User
	documents map[string]*Document
	dataDir   string
}

func NewSharedDocsSystem() *SharedDocsSystem {
	dataDir := "./data"
	os.MkdirAll(dataDir, 0755)
	os.MkdirAll(filepath.Join(dataDir, "users"), 0755)
	os.MkdirAll(filepath.Join(dataDir, "documents"), 0755)

	return &SharedDocsSystem{
		users:     make(map[string]*User),
		documents: make(map[string]*Document),
		dataDir:   dataDir,
	}
}

func main() {
	system := NewSharedDocsSystem()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Welcome to Shared Docs System ===")
	fmt.Println("A simplified collaborative document editing system")
	fmt.Println()

	for {
		fmt.Println("Main Menu:")
		fmt.Println("1. Create User")
		fmt.Println("2. Login")
		fmt.Println("3. List Users")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			system.createUser(reader)
		case "2":
			system.login(reader)
		case "3":
			system.listUsers()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
		fmt.Println()
	}
}
