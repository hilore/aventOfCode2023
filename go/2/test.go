package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var R int = 12
var G int = 13
var B int = 14

func main() {
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
	}

	fmt.Printf("Red: %d\n", R)
	fmt.Printf("Green: %d\n", G)
	fmt.Printf("Blue: %d\n\n", B)

	var result int

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		arr := strings.Split(line, ":")
		game := arr[0]
		other := strings.TrimSpace(arr[1])
		gameNumber, _ := strconv.Atoi(strings.Split(game, " ")[1])

		find := false
		cubes := strings.Split(other, ";")

		for _, item := range cubes {
			seq := strings.Split(strings.TrimSpace(item), ",")

			rErr := false
			gErr := false
			bErr := false

			for _, cubeItem := range seq {
				tmp := strings.TrimSpace(cubeItem)
				num, _ := strconv.Atoi(strings.Split(tmp, " ")[0])
				color := strings.Split(tmp, " ")[1]
				if color == "red" {
					if num > R {
						rErr = true
						break
					}
				} else if color == "green" {
					if num > G {
						gErr = true
						break
					}
				} else if color == "blue" {
					if num > B {
						bErr = true
						break
					}
				}
			}

			if rErr == true || gErr == true || bErr == true {
				find = false
				break
			} else {
				find = true
			}
		}

		if find == true {
			fmt.Printf("Game %d\n", gameNumber)
			result += gameNumber
			fmt.Printf("---------------------------------\n")
		}
	}

	fmt.Println(result)
}
