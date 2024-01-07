package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func rec(key int, rng int, matches map[string]int, record map[int]int) {
	for number := key + 1; number <= key+rng; number++ {
		record[number]++
		numberStr := strconv.Itoa(number)
		rec(number, matches[numberStr], matches, record)
	}
}

func main() {
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
	}

	record := make(map[string]int)
	matches := make(map[string]int)
	var str string

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		cardSeq := strings.Split(strings.Split(line, ":")[0], " ")
		cardNum := cardSeq[len(cardSeq)-1]
		fmt.Printf("Card #%s\n", cardNum)

		matches[cardNum] = 0
		record[cardNum] = 0

		tmp := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), "|")
		winNums := strings.Split(strings.TrimSpace(tmp[0]), " ")
		nums := strings.Split(strings.TrimSpace(tmp[1]), " ")
		fmt.Printf("Win nums: %q\n", winNums)
		fmt.Printf("My nums: %q\n", nums)

		winNumsMap := make(map[string]string)
		for _, winNum := range winNums {
			if winNum == "" {
				continue
			}

			winNumsMap[winNum] = ""
		}

		for _, num := range nums {
			_, ok := winNumsMap[num]
			if ok {
				fmt.Printf("%s ", num)
				matches[cardNum]++
			}
		}

		str += cardNum + ":" + strconv.Itoa(matches[cardNum]) + ";"
		fmt.Printf("\n---------------------------------------------------------------\n")
	}
	seq := strings.Split(str, ";")
	originals := make(map[int]int)

	for _, item := range seq {
		if item == "" {
			continue
		}

		orgStr := strings.Split(item, ":")[0]
		org, _ := strconv.Atoi(orgStr)

		rep := matches[strconv.Itoa(org)]
		originals[org]++
		rec(org, rep, matches, originals)
	}

	fmt.Printf("%v\n", originals)

  var result int
  for _, value := range originals {
    result += value
  }

  fmt.Println(result)
}
