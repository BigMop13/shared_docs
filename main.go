package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//reads a data from file
//goroutine to read values from channel
//goroutine to write values to file

const dataDir = "."

type User struct {
	Username string
	Email    string
	Created  time.Time
}

type Document struct {
	Name     string
	Created  time.Time
	Modified time.Time
}

type SharedDocsSystem struct {
	users     map[string]*User
	documents Document
	dataDir   string
}

func NewSharedDocsSystem() *SharedDocsSystem {

	return &SharedDocsSystem{
		users: make(map[string]*User),
		documents: Document{
			ID: "document.txt",
		},
		dataDir: dataDir,
	}
}

func main() {
	system := NewSharedDocsSystem()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Welcome to Shared Docs System ===")
	fmt.Println("A simplified collaborative document editing system")
	fmt.Println()

	for {
		//fmt.Println("Main Menu:")
		//fmt.Println("1. Create User")
		//fmt.Println("2. Login")
		//fmt.Println("3. List Users")
		//fmt.Println("4. Exit")
		fmt.Println("1. Edit file")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			system.createUser(reader)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
		fmt.Println()
	}
}
