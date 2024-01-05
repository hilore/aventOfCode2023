package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
  "strconv"
  "math"
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

  var result float64

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

    split := strings.Split(line, ":")
    seq := strings.Split(strings.TrimSpace(split[1]), ";")

    var red float64
    var green float64
    var blue float64
    var tmpRes float64

    fmt.Println(line)

    for _, item := range seq {
      cubes := strings.Split(item, ",")
      for _, cube := range cubes {
        tmp := strings.Split(strings.TrimSpace(cube), " ")
        tmpNum, _ := strconv.Atoi(tmp[0])
        num := float64(tmpNum)
        color := strings.TrimSpace(tmp[1])

        if color == "red" {
          red = math.Max(red, num)
        } else if color == "green" {
          green = math.Max(green, num)
        } else if color == "blue" {
          blue = math.Max(blue, num)
        }
      }
    }

    tmpRes = red * green * blue
    result += tmpRes

    fmt.Printf("%f %f %f\n-------------------\n", red, green, blue)
  }

  fmt.Printf("Result - %f\n", result)
}
