package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var nums = map[string]int{
	"one": 1, "two": 2,
	"three": 3, "four": 4,
	"five": 5, "six": 6,
	"seven": 7, "eight": 8,
	"nine": 9, "zero": 0,
}

func digitFront(word string) (int, int) {
	for idx, ch := range word {
		if unicode.IsDigit(ch) {
			value, _ := strconv.Atoi(string(ch))
			return idx, value
		}
	}

	return -1, -1
}

func digitBack(word string) (int, int) {
	for i := len(word) - 1; i > -1; i-- {
		if unicode.IsDigit(rune(word[i])) {
			value, _ := strconv.Atoi(string(word[i]))
			return i, value
		}
	}

	return -1, -1
}

func stringFront(word string) (int, int) {
	find := false
	var tmp string

	for idx, ch := range word {
		tmp += string(ch)

		for i := idx + 1; i < len(word); i++ {
			tmp += string(word[i])
			if len(tmp) > 5 && find == false {
				tmp = ""
				break
			}

			if len(tmp) >= 3 {
				value, ok := nums[tmp]
				if ok {
					find = true
					resIdx := i - len(tmp) + 1
					return resIdx, value
				}
			}
		}

		if find == false {
			tmp = ""
		}
	}

	return -1, -1
}

func stringBack(word string) (int, int) {
	find := false
	var tmp string

	for i := len(word) - 1; i > -1; i-- {
		tmp += string(word[i])

		for j := i - 1; j > -1; j-- {
			tmp += string(word[j])

			if len(tmp) > 5 && find == false {
				tmp = ""
				break
			}

			if len(tmp) >= 3 {
				reversedTmp := reverseString(tmp)
				value, ok := nums[reversedTmp]
				if ok {
					find = true
					return j, value
				}
			}
		}

		if find == false {
			tmp = ""
		}
	}

	return -1, -1
}

func reverseString(str string) string {
	var result string
	for i := len(str) - 1; i > -1; i-- {
		result += string(str[i])
	}

	return result
}

func main() {
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
	}

	var result int

	for _, word := range strings.Split(string(data), "\n") {
		// fmt.Printf("Word %d - %s\n", wordIdx, word)
		// fmt.Printf("Word - %s\n", word)
		if word == "" {
			continue
		}

		fmt.Println(word)
		var front int
		var back int
		var tmpResult string

		// front
		strIdx, strVal := stringFront(word)
		digIdx, digVal := digitFront(word)
		fmt.Printf("String Front: index %d, value %d\n", strIdx, strVal)
		fmt.Printf("Digit Front: index %d, value %d\n", digIdx, digVal)

		if strIdx == -1 && digIdx == -1 {
			continue
		} else if strIdx == -1 && digIdx > -1 {
			front += digVal
		} else if strIdx > -1 && digIdx == -1 {
			front += strVal
		} else if strIdx > -1 && digIdx > -1 {
			if strIdx < digIdx {
				front += strVal
			} else if strIdx > digIdx {
				front += digVal
			}
		}

		// back
		strBackIdx, strBackValue := stringBack(word)
		digBackIdx, digBackValue := digitBack(word)
		fmt.Printf("String Back: index %d, value %d\n", strBackIdx, strBackValue)
		fmt.Printf("Digit Back: index %d, value %d\n", digBackIdx, digBackValue)

		if strBackIdx == -1 && digBackIdx == -1 {
			continue
		} else if strBackIdx == -1 && digBackIdx > -1 {
			back += digBackValue
		} else if strBackIdx > -1 && digBackIdx == -1 {
			back += strBackValue
		} else if strBackIdx > -1 && digBackIdx > -1 {
			if strBackIdx < digBackIdx {
				back += digBackValue
			} else if strBackIdx > digBackIdx {
				back += strBackValue
			}
		}

		tmpResult = strconv.Itoa(front) + strconv.Itoa(back)
		tmpVal, _ := strconv.Atoi(tmpResult)
		result += tmpVal

		fmt.Printf("Front: %d\n", front)
		fmt.Printf("Back: %d\n", back)
		fmt.Printf("Result: %s\n", tmpResult)
		// fmt.Println("---------------------------------------------")
	}

	fmt.Printf("Total - %d\n", result)
}
