package handlers

import (
	"fmt"
	"os"
	"reinder/models"
)

type DocumentHandler struct{}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{}
}

func (documentHandler *DocumentHandler) EditFile(document models.Document) {
	clearScreen()
	data, err := os.ReadFile(document.Name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
