package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveController(word string) {

	file, err := os.Open("word-list/en.txt")
	dstFilePath := "word-list/en_copy"
	check(err)
	defer file.Close()

	dst, err := os.Create(dstFilePath)
	check(err)
	defer dst.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if contains(word, scanner.Text()) {
			fmt.Printf("word %s found, skipping write\n", scanner.Text())
			time.Sleep(1 * time.Second)

			continue
		}

		_, err = fmt.Fprintln(dst, scanner.Text())
		check(err)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//delete the original file last, so that we can rollback ??
	file.Close()
	err = os.Remove("word-list/en.txt")
	check(err)

	dst.Close()
	err = os.Rename(dstFilePath, "word-list/en.txt")
	check(err)

	fmt.Println("completed")

}

func contains(word string, str string) bool {
	if reflect.DeepEqual(word, str) {
		return true
	}

	return false
}
