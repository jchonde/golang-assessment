package controllers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ListController(w http.ResponseWriter) {
	file, err := os.Open("word-list/en.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io.WriteString(w, scanner.Text()+"\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
