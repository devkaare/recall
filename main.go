package main

import (
	"bufio"
	"fmt"
	"os"

	// "math"
	"math/rand"
)

type QA struct {
	Question, Awnser string
}

var questions = map[int]QA{
	0: {
		"Example question #0",
		"Example awnser #0\n",
	},
	1: {
		"Example question #1",
		"Example awnser #1\n",
	},
	2: {
		"Example question #2",
		"Example awnser #2\n",
	},
	3: {
		"Example question #3",
		"Example awnser #3\n",
	},
}

var solved map[int]QA

func main() {
	// TODO: Option mode instead?
	for {
		var q QA

		// Check if question has been solved
		randNum := rand.Intn(3)

		if _, ok := solved[randNum]; !ok {
			q = questions[randNum]
		}

		r := bufio.NewReader(os.Stdin)
		fmt.Println(q.Question)
		awnser, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Check awnser
		if awnser != q.Awnser {
			fmt.Printf("Wrong:\nExpected: %sAwnser: %s\n", q.Awnser, awnser)
			return
		}

		fmt.Println("Correct!")
	}
}
