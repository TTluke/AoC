package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Lists struct {
	Left  []int
	Right []int
}

func ParseInput() (*Lists, error) {
	Lists := &Lists{}

	reader, err := os.Open("./input")
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ids := strings.Fields(scanner.Text())
		left_id, _ := strconv.Atoi(ids[0])
		right_id, _ := strconv.Atoi(ids[1])
		Lists.Left = append(Lists.Left, left_id)
		Lists.Right = append(Lists.Right, right_id)
	}

	sort.Ints(Lists.Left)
	sort.Ints(Lists.Right)

	return Lists, nil

}

func Solve1(input *Lists) int {
	answer := 0
	for idx := range len(input.Left) {
		distance := input.Left[idx] - input.Right[idx]
		if distance < 0 {
			distance = distance * -1
		}
		answer += distance
	}
	return answer
}

func Solve2(input *Lists) int {
	answer := 0
	rightMap := make(map[int]int)
	for _, num := range input.Right {
		rightMap[num]++
	}
	for _, num := range input.Left {
		answer += num * rightMap[num]
	}
	return answer
}

func main() {
	parse_start := time.Now()
	input, _ := ParseInput()
	parse_end := time.Since(parse_start)

	solve1_start := time.Now()
	result1 := Solve1(input)
	solve1_end := time.Since(solve1_start)

	solve2_start := time.Now()
	result2 := Solve2(input)
	solve2_end := time.Since(solve2_start)

	fmt.Printf("Parsing took: %v\n", parse_end)
	fmt.Printf("Solution 1: %v finished in: %v\n", result1, solve1_end)
	fmt.Printf("Solution 2: %v finished in: %v\n", result2, solve2_end)
}
