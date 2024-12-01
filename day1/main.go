package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)


/**
    This Open a file scans it
**/
func Open(fileName string) []string {
    var text []string
    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal(err)
    }


    defer file.Close()


    scanner := bufio.NewScanner(file)

    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        text = append(text, scanner.Text())
    }

    return text
}


/**
    In the first part we sort the location IDs.
**/
func firstPart(first []int, second []int) int {
    var sum int = 0
    if len(first) != len(second) {
        log.Fatal("the passed slices are not the same")
    }

    bubbleSort(first)
    bubbleSort(second)

    for i := range first {
        var difference float64 = float64(first[i]) - float64(second[i])
        sum += int(math.Abs(difference))
    }

    return sum
}

func bubbleSort(list []int) {
    n := len(list)

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if list[j] > list[j+1] {
                list[j], list[j+1] = list[j+1], list[j]
            }
        }
    }
}

func ReadData(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		for i, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
			if i%2 == 0 {
				left = append(left, val)
			} else {
				right = append(right, val)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}


/**
    In the second part we search for the similarity score
**/
func secondPart(first []int, second []int) int {
    var sum int

    for _, ID1 := range first {
        var foo int
        for _, ID2 := range second {
            if ID1 != ID2 {
                continue
            }
            foo += 1
        }
        sum += foo * ID1
    }


    return sum
}


/**
    Here we handle both.
**/
func main() {
    
    var List1, List2 []int
    var differenceSum int
    var similar int

    for _, line := range Open("./inputData.txt") {
        var numbers []string =  strings.Fields(line)

        ID1, err := strconv.Atoi(numbers[0]) 
        if err != nil {
            fmt.Printf("Could not convert List1 items %s to integer\n", numbers[0])
        }

        ID2, err := strconv.Atoi(numbers[1]) 
        if err != nil {
            fmt.Printf("Could not convert List2 itmes %s to integer\n", numbers[1])
        }

        List1 = append(List1, ID1)
        List2 = append(List2, ID2)
    }

    differenceSum = firstPart(List1, List2)
    fmt.Printf("Difference between two lists is %v\n", differenceSum)

    similar = secondPart(List1, List2)
    fmt.Printf("The similarity between the two list is %v\n", similar)
}
