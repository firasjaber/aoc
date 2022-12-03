package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read the input from the text file
	b, err := os.ReadFile("./1/input.txt") // just pass the file name
  if err != nil {
    fmt.Print(err)
  }

	// Convert the input to string then split it by empty new lines
  str := string(b)
	elfsCalories := splitByEmptyNewline(str)

	// Init the max calories sum to 0
	maxCaloriesSum := 0

	// Loop in the array of elfs calories and calculate the sum 
	// Then check if the elf calories sum is greater then the previous max
	for i := 0; i < len(elfsCalories); i++ {
		fmt.Printf("summing calories for elf %d \n",i)
		elfCaloriesSum := sumOfElfCalories(elfsCalories[i])
		fmt.Printf("calories sum for elf %d is %d \n",i,elfCaloriesSum)
	
		if maxCaloriesSum < elfCaloriesSum {
			maxCaloriesSum = elfCaloriesSum
		}
	}

	fmt.Printf("the most calories carrying elf is having %d calories", maxCaloriesSum)
}

func sumOfElfCalories(str string) int {
	arr := strings.Split(str, "\n")
	sum := 0
	for i := 0; i < len(arr); i++ {
		integer,err := strconv.Atoi(strings.TrimSpace(arr[i]))
		if err != nil {
			panic(err)
		}

		sum += int(integer)
	}

	return sum
}

func splitByEmptyNewline(str string) []string {
  strNormalized := regexp.
    MustCompile("\r\n").
    ReplaceAllString(str, "\n")

  return regexp.
    MustCompile(`\n\s*\n`).
    Split(strNormalized, -1)
}