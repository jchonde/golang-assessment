package controllers

import (
	"fmt"
	"os"
)

func AddController(word string) {
	file, err := os.OpenFile("word-list/en.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err = file.WriteString("\n" + word); err != nil {
		fmt.Println(err)
	}
}
