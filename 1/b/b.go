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
	secondMaxCaloriesSum := 0
	thirdMaxCaloriesSum := 0

	// Loop in the array of elfs calories and calculate the sum 
	// Then check if the elf calories sum is greater then the previous max
	for i := 0; i < len(elfsCalories); i++ {
		fmt.Printf("summing calories for elf %d \n",i)
		elfCaloriesSum := sumOfElfCalories(elfsCalories[i])
		fmt.Printf("calories sum for elf %d is %d \n",i,elfCaloriesSum)
	
		if maxCaloriesSum < elfCaloriesSum {
			thirdMaxCaloriesSum = secondMaxCaloriesSum
			secondMaxCaloriesSum = maxCaloriesSum
			maxCaloriesSum = elfCaloriesSum
		} else if secondMaxCaloriesSum < elfCaloriesSum {
			thirdMaxCaloriesSum = secondMaxCaloriesSum
			secondMaxCaloriesSum = elfCaloriesSum
		} else if thirdMaxCaloriesSum < elfCaloriesSum {
			thirdMaxCaloriesSum = elfCaloriesSum
		}
	}

	topThreeCaloriesSum := maxCaloriesSum + secondMaxCaloriesSum + thirdMaxCaloriesSum

	fmt.Printf("the top 3 most calories carrying elves are having %d calories", topThreeCaloriesSum)
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