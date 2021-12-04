package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    input := readInput("input.txt")
    // fmt.Println(getPowerConsumption(input))
    // fmt.Println(getOxygenGeneratorRating(input))
    fmt.Println(getScrubberRating(input) * getOxygenGeneratorRating(input))
}

func getScrubberRating(report []string) int {
    var valid []string
    var bit_1, bit_0 []string
    bitPos := 0
    for len(valid) != 1 {
        for _, number := range report {
            if bitPos == 0 || contains(number, valid) {
                if number[bitPos] == 48 {
                    bit_0 = append(bit_0, number)
                } else {
                    bit_1 = append(bit_1, number)
                }
            }
        }

        if len(bit_0) > len(bit_1) {
            valid = bit_1
        } else {
            valid = bit_0
        }
        bit_0, bit_1 = nil, nil
        bitPos++
    }
    rating, _ := strconv.ParseInt(valid[0], 2, 0)

    return int(rating)
}

func getOxygenGeneratorRating(report []string) int {
    // create 1 main array that has the valid running values
    // iterate through the diagnostic report
    // iterate on each bit
        // if index is not on mainArr, skip
        // create 2 new arrays, one for '1' bit another for '0' bit, add only indexes
        // mainArr = array with more elements
    
    var valid []string
    var bit_1, bit_0 []string
    bitPos := 0
    for len(valid) != 1 {
        for _, number := range report {
            if bitPos == 0 || contains(number, valid) {
                if number[bitPos] == 48 {
                    bit_0 = append(bit_0, number)
                } else {
                    bit_1 = append(bit_1, number)
                }
            }
        }

        if len(bit_0) > len(bit_1) {
            valid = bit_0
        } else {
            valid = bit_1
        }
        bit_0, bit_1 = nil, nil
        bitPos++
    }
    rating, _ := strconv.ParseInt(valid[0], 2, 0)

    return int(rating)
}

func getPowerConsumption(diagnosticReport []string) int {
    // traverse every bit 
    // keep track of 1s and 0s in every pos
    // if bit == 1 add 1 to that bit otherwise subtract 1
    bits := make(map[int]int)
    for _, number := range diagnosticReport {
        for i, char := range number {
            if (char == 48) {
                bits[i+1]--
            } else {
                bits[i+1]++
            }
        }
    }

    var gammaRateBinary string
    var epsilonRateBinary string
    for i := 1; i <= len(bits); i++ {
        if bits[i] > 0 {
            gammaRateBinary += "1"
            epsilonRateBinary += "0"
        } else {
            gammaRateBinary += "0"
            epsilonRateBinary += "1"
        }
    }

    gammaRate, _ := strconv.ParseInt(gammaRateBinary, 2, 0)
    epsilonRate, _ := strconv.ParseInt(epsilonRateBinary, 2, 0)

    return int(gammaRate)* int(epsilonRate)
}

func readInput(fname string) []string {
    file, ferr := os.Open(fname)
    if (ferr) != nil {
        panic(ferr)
    }
    scanner := bufio.NewScanner(file)
    var input []string

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    return input
}

func contains(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}