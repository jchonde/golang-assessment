package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func RemoveController(word string) (string, error) {

	file, err := os.Open("word-list/en.txt")
	dstFilePath := "word-list/en_copy"
	if err != nil {
		log.Fatal(err)
		return "error", err
	}
	defer file.Close()

	dst, err := os.Create(dstFilePath)
	if err != nil {
		log.Fatal(err)
		return "error", err
	}
	defer dst.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if contains(word, scanner.Text()) {
			time.Sleep(1 * time.Second)
			continue
		}

		_, err = fmt.Fprintln(dst, scanner.Text())
		if err != nil {
			log.Fatal(err)
			return "error", err
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//delete the original file last, so that we can rollback ??
	file.Close()
	err = os.Remove("word-list/en.txt")
	if err != nil {
		log.Fatal(err)
		return "error", err
	}

	dst.Close()
	err = os.Rename(dstFilePath, "word-list/en.txt")
	if err != nil {
		log.Fatal(err)
		return "error", err
	}

	return "completed", err

}

func contains(word string, str string) bool {
	if reflect.DeepEqual(word, str) {
		return true
	}

	return false
}
