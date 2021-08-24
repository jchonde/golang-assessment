package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func CompareController(w http.ResponseWriter, word1 string, word2 string) {
	file, err := os.Open("word-list/en.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//two arrays for both words
	word1AsArrayNums := [26]int{}
	word2AsArrayNums := [26]int{}

	//fill arrays
	for i := 0; i < len(word1); i++ {
		word1AsArrayNums[int(word1[i])-int('a')]++
	}

	for i := 0; i < len(word2); i++ {
		word2AsArrayNums[int(word2[i])-int('a')]++
	}
	if strings.Compare(word1, word2) == 0 {
		io.WriteString(w, "They are the same\n")

	} else if len(word1) == len(word2) && reflect.DeepEqual(word1AsArrayNums, word2AsArrayNums) && strings.Compare(word1, word2) != 0 {
		io.WriteString(w, "They are anagrams\n")
	} else {
		io.WriteString(w, "They are NOT anagrams\n")
	}

}
