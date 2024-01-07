package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

func main() {
  file := os.Args[1]
  data, err := ioutil.ReadFile(file)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
  }

  var pointsTotal int

  for _, line := range strings.Split(string(data), "\n") {
    if line == "" {
      continue
    }

    var pointsPerCard int

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
        if pointsPerCard == 0 {
          pointsPerCard = 1
        } else {
          pointsPerCard *= 2
        }
      }
    }

    fmt.Printf("\nPoints - %d\n", pointsPerCard)
    pointsTotal += pointsPerCard

    fmt.Printf("---------------------------------------------------------------\n")
  }

  fmt.Printf("Points total - %d\n", pointsTotal)
}
