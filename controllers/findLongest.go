package controllers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

func FindLongestController(w http.ResponseWriter) {

	longestWordInFile, longestWordSize := findLongest("word-list/en.txt")
	isAnagram := false
	updatedWord := false

	if longestWordSize > 0 {
		isAnagram = wordHasAnagram(longestWordInFile, "word-list/en.txt")
		if !isAnagram {
			for longestWordSize > 0 {
				longestWordInFile, longestWordSize, updatedWord = findNextLongest("word-list/en.txt", longestWordSize)
				if !updatedWord {
					io.WriteString(w, "There are NO Anagrams in file\n")
					break
				}
				isAnagram = wordHasAnagram(longestWordInFile, "word-list/en.txt")
				if isAnagram {
					printAnagrams(w, longestWordInFile, "word-list/en.txt")
					break
				}
			}
		}
	} else {
		io.WriteString(w, "There are no words in file\n")
	}

}

//find the longest word in the file
func findLongest(fileName string) (word string, longestSize int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	longest := 0
	longestWord := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > longest {
			longest = len(scanner.Text())
			longestWord = scanner.Text()
		}
	}

	return longestWord, longest
}

func findNextLongest(fileName string, currentLongestSize int) (word string, longestSize int, updatedWord bool) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	longestSize = 0
	longestWord := ""
	updatedWord = false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > longestSize && len(scanner.Text()) < currentLongestSize {
			longestSize = len(scanner.Text())
			longestWord = scanner.Text()
			updatedWord = true
		}
	}

	fmt.Println(longestSize)
	return longestWord, longestSize, updatedWord
}

func wordHasAnagram(word string, fileName string) (isAnagram bool) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	wordAsArrayNums := [26]int{}
	for i := 0; i < len(word); i++ {
		if word[i] == ' ' {
			continue
		}
		wordAsArrayNums[int(word[i])-int('a')]++
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(word) == len(scanner.Text()) {
			dictionaryWordAsArrayNum := [26]int{}
			for i := 0; i < len(scanner.Text()); i++ {
				if scanner.Text()[i] == ' ' {
					continue
				}
				dictionaryWordAsArrayNum[int(scanner.Text()[i])-int('a')]++
			}
			if reflect.DeepEqual(wordAsArrayNums, dictionaryWordAsArrayNum) && !reflect.DeepEqual(word, scanner.Text()) {
				isAnagram = true
				break
			}

		}
	}
	return isAnagram
}

func printAnagrams(w http.ResponseWriter, longestWordWithAnagrams string, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	wordAsArrayNums := [26]int{}

	for i := 0; i < len(longestWordWithAnagrams); i++ {
		if longestWordWithAnagrams[i] == ' ' {
			continue
		}
		wordAsArrayNums[int(longestWordWithAnagrams[i])-int('a')]++
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(longestWordWithAnagrams) == len(scanner.Text()) {
			dictionaryWordAsArrayNum := [26]int{}
			for i := 0; i < len(scanner.Text()); i++ {
				if scanner.Text()[i] == ' ' {
					continue
				}
				dictionaryWordAsArrayNum[int(scanner.Text()[i])-int('a')]++
			}
			if reflect.DeepEqual(wordAsArrayNums, dictionaryWordAsArrayNum) {
				io.WriteString(w, scanner.Text()+"\n")
			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
