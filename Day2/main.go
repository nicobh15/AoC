package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	text := string(b[:])
	rows := strings.Split(text, "\n")
	safe := 0

	for _, row := range rows {
		parts := strings.Split(row, " ")
		nums := make([]int, len(parts))
		for i, p := range parts {
			val, _ := strconv.Atoi(p)
			nums[i] = val
		}

		if checkSafety(nums) {
			safe++
		} else {
			madeSafe := false
			for i := 0; i < len(nums); i++ {
				newNums := append([]int(nil), nums[:i]...)

				newNums = append(newNums, nums[i+1:]...)
				if checkSafety(newNums) {
					madeSafe = true
					break
				}
			}
			if madeSafe {
				safe++
			}
		}

	}
	fmt.Println(safe)
}

func checkSafety(nums []int) bool {
	var diff int
	for i := range nums {
		if i == 0 {
			continue
		} else if i == 1 {
			prev := nums[i-1]
			curr := nums[i]
			diff = curr - prev
			if diff == 0 || diff < -3 || diff > 3 {
				return false
			}
		}

		prev := nums[i-1]
		curr := nums[i]

		curr_diff := curr - prev
		if (diff > 0 && curr_diff < 0) || (diff < 0 && curr_diff > 0) {
			return false
		}

		if !((curr_diff > 0 && curr_diff < 4) || (curr_diff < 0 && curr_diff > -4)) {
			return false
		}
	}
	return true
}
