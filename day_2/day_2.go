package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
  file, ferr := os.Open("input.txt")
  if (ferr) != nil {
    panic(ferr)
  }

  scanner := bufio.NewScanner(file)
  
  var horizontal int
  var depth int
  var aim int
  for scanner.Scan() {
	command := scanner.Text()
	items := strings.Split(command, " ")
	direction := items[0]
	unit, err := strconv.Atoi(items[1])
    if err != nil {
        log.Println(err)
        continue
    }

	if direction == "forward" {
		horizontal += unit 
		depth += aim * unit
	} else if direction == "down" {
		aim += unit
	} else if direction == "up" {
		aim -= unit
	}
  }

  fmt.Println(horizontal * depth)
}

func part1() {
  file, ferr := os.Open("input.txt")
  if (ferr) != nil {
    panic(ferr)
  }

  scanner := bufio.NewScanner(file)
  
  var horizontal int
  var depth int
  for scanner.Scan() {
	command := scanner.Text()
	items := strings.Split(command, " ")
	direction := items[0]
	unit, err := strconv.Atoi(items[1])
    if err != nil {
        log.Println(err)
        continue
    }

	if direction == "forward" {
		horizontal += unit 
	} else if direction == "down" {
		depth += unit
	} else if direction == "up" {
		depth -= unit
	}
  }

  fmt.Println(horizontal * depth)

}
