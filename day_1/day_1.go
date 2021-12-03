package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  part1()
  part2()
}

func readInput(fname string) []int64 {
  file, ferr := os.Open(fname)
  if (ferr) != nil {
    panic(ferr)
  }
  scanner := bufio.NewScanner(file)
  input := []int64{}

  for scanner.Scan() {
  depth, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    input = append(input, depth)
  }

  return input
}

func countLargerSums(arr []int64) int {
  first := true
  var prevDepth int64 = 0
  var res int
  for _, depth := range arr {
    if first {
      first = false
      prevDepth = depth
        continue
    }
    if depth > prevDepth {
      res++
    }
    prevDepth = depth
  }

  return res
}

func part1() {
  input := readInput("test.txt")
  fmt.Println(countLargerSums(input))
}

func part2() {
  input := readInput("test.txt")

  threeMeasurementWindow := []int64{}
  for i, value := range input {
    if (i == len(input) - 2) {
      break
    }
    var sumOfThree int64 = value + input[i+1] + input[i+2]
    threeMeasurementWindow = append(threeMeasurementWindow, sumOfThree)
  }

  fmt.Println(countLargerSums(threeMeasurementWindow))
}