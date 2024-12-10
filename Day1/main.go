package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair[T, U int] struct {
	First  T
	Second U
}

func Zip[T, U int](ts []T, us []U) []Pair[T, U] {
	if len(ts) != len(us) {
		panic("slices have different length")
	}
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}

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
	rows = rows[:len(rows)-1]
	var left_col []int
	var right_col []int

	for _, row := range rows {
		nums := strings.Split(row, "   ")
		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])
		left_col = append(left_col, left_num)
		right_col = append(right_col, right_num)
	}
	sort.Ints(left_col)
	sort.Ints(right_col)

	pairs := Zip(left_col, right_col)
	sum_diffs := 0

	for _, pair := range pairs {

		diff := pair.First - pair.Second

		if diff < 0 {
			sum_diffs += -diff
		} else {
			sum_diffs += diff
		}
	}

	fmt.Print(sum_diffs)

}
