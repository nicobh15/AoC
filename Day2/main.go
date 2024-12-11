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

	// var reports [][]int
	safe := 0
	unsafe := 0

	for _, row := range rows {
		nums := strings.Split(row, " ")
		flag := 0
		diff := 0
		for i := range nums {
			if i == 0 {
				continue
			} else if i == 1 {
				prev, _ := strconv.Atoi(nums[i-1])
				curr, _ := strconv.Atoi(nums[i])
				diff = curr - prev
				if diff == 0 || diff < -3 || diff > 3 {
					flag = 1
					break
				}
			}

			prev, _ := strconv.Atoi(nums[i-1])
			curr, _ := strconv.Atoi(nums[i])

			curr_diff := curr - prev
			if (diff > 0 && curr_diff < 0) || (diff < 0 && curr_diff > 0) {
				flag = 1
				break
			}

			if !((curr_diff > 0 && curr_diff < 4) || (curr_diff < 0 && curr_diff > -4)) {
				flag = 1
				break
			}
		}
		if flag != 0 {
			flag = 0
			unsafe += 1
		} else {
			safe += 1
		}
	}
	fmt.Println(safe)
}
