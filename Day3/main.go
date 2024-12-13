package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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
	if err != nil {
		log.Fatal(err)
	}
	text := string(b)
	ans := 0
	mulEnabled := true

	mulRe := regexp.MustCompile(`mul\(\s*\d{1,3}\s*,\s*\d{1,3}\s*\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	allRe := regexp.MustCompile(`mul\(\s*\d{1,3}\s*,\s*\d{1,3}\s*\)|do\(\)|don't\(\)`)

	queue := allRe.FindAllString(text, -1)

	for _, instruction := range queue {
		if doRe.MatchString(instruction) {
			mulEnabled = true
		} else if dontRe.MatchString(instruction) {
			mulEnabled = false
		} else if mulRe.MatchString(instruction) {
			if mulEnabled {
				ans += multiplySubstring(instruction)
			}
		}
	}

	fmt.Print(ans)
}

func multiplySubstring(substring string) int {
	sides := regexp.MustCompile(`\d+`).FindAllString(substring, -1)
	num1, err1 := strconv.Atoi(sides[0])
	num2, err2 := strconv.Atoi(sides[1])
	if err1 != nil || err2 != nil {
		return 0
	}
	return num1 * num2
}
