package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
	}

	var result int
	var tmp string
	digits := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, word := range strings.Split(string(data), "\n") {
		if word == "" {
			continue
		}

    findFirst := false
    for _, ch := range word {
			for digit := range digits {
				if string(ch) == strconv.Itoa(digit) {
					tmp += string(ch)
          findFirst = true
					break
				}
			}
      if findFirst == true {
        break
      }
		}

		for i := len(word) - 1; i > -1; i-- {
			if string(word[i]) == "" {
				continue
			}

      findSecond := false
      for _, ch2 := range string(word[i]) {
				for digit2 := range digits {
					if string(ch2) == strconv.Itoa(digit2) {
						tmp += string(ch2)
            findSecond = true
						break
					}
				}
			}
      if findSecond == true {
        break
      }
		}
		fmt.Printf("Word - %s, number - %c\n", word, tmp)
    tmpResult, _ := strconv.Atoi(tmp)
    result += tmpResult
    tmp = ""
	}

  fmt.Println("=================================================")
  fmt.Printf("Result - %d\n", result)
}
