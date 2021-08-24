package controllers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

func FindController(w http.ResponseWriter, word string) {
	file, err := os.Open("word-list/en.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//array for every word in the english alphabet
	wordAsArrayNums := [26]int{}

	//count letters and fill them in the array (a=0, b=1, c=2...)
	for i := 0; i < len(word); i++ {

		//if the row has two words ex. fairy tales
		if word[i] == ' ' {
			continue
		}
		wordAsArrayNums[int(word[i])-int('a')]++
	}

	hasAnagram := false
	scanner := bufio.NewScanner(file)

	//iterate file
	for scanner.Scan() {
		if len(word) == len(scanner.Text()) {
			dictionaryWordAsArrayNum := [26]int{}

			//count letters and fill them in the array (a=0, b=1, c=2...)
			for i := 0; i < len(scanner.Text()); i++ {
				if scanner.Text()[i] == ' ' {
					continue
				}
				dictionaryWordAsArrayNum[int(scanner.Text()[i])-int('a')]++
			}
			if reflect.DeepEqual(wordAsArrayNums, dictionaryWordAsArrayNum) && !reflect.DeepEqual(word, scanner.Text()) {
				io.WriteString(w, scanner.Text()+"\n")
				hasAnagram = true
			}

		}
	}
	if !hasAnagram {
		io.WriteString(w, "There are no anagrams\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
