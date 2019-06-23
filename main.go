package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateCode(length int) string {
	var code []string

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < length; i++ {
		n := r.Intn(9)
		code = append(code, strconv.Itoa(n))
	}

	return strings.Join(code, "")
}

func getAnswer(length int, step int) string {
	var input string
	fmt.Printf("Answer %d: ", step)
	fmt.Scan(&input)

	for len(input) > length || len(input) < length {
		fmt.Printf("answer must be %d characters!\n", length)
		fmt.Printf("Answer %d: ", step)
		fmt.Scan(&input)
	}

	return input
}

func count(arr []string) map[string]int {
	counts := make(map[string]int)

	for _, item := range arr {
		if _, exists := counts[item]; exists {
			counts[item] += 1
		} else {
			counts[item] = 1
		}
	}

	return counts
}

func checkAnswer(code string, answer string) (int, int) {
	placement := 0
	color := 0

	codeArr := strings.Split(code, "")
	answerArr := strings.Split(answer, "")

	codeCounts := count(codeArr)
	answerCounts := count(answerArr)

	for index, codeCount := range codeCounts {
		if answerCounts[index] >= codeCount {
			color += codeCount
		} else if answerCounts[index] < codeCount {
			color += answerCounts[index]
		}
	}

	for index, item := range codeArr {
		if answerArr[index] == item {
			placement += 1
		}
	}

	color -= placement

	return placement, color
}

func main() {
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("can't parse length: %v\n", err)
	}

	step := 0
	solved := false

	code := generateCode(length)

	for !solved {
		answer := getAnswer(length, step)
		p, c := checkAnswer(code, answer)

		fmt.Printf("placement correct: %d, color correct: %d\n\n", p, c)

		step += 1

		if p == length {
			fmt.Println("Congratulations!")
			solved = true
		}
	}
}
