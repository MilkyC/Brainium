package main

import (
	"fmt"
	"github.com/MilkyC/Brainium/challenges"
	"time"
)

type countdown struct {
	count  int
	ticker *time.Ticker
}

func main() {
	countdowner := countdown{
		count:  0,
		ticker: time.NewTicker(time.Millisecond * 1000),
	}
	seed := int64(time.Now().Unix())
	mc := challenges.NewMathChallenge(seed)

	for true {
		userDone := make(chan bool, 1)
		countdownDone := make(chan bool, 1)
		go func(countdownDone chan<- bool) {
			for range countdowner.ticker.C {
				countdowner.count++
				fmt.Printf("%d seconds\n", countdowner.count)
				if countdowner.count > 10 {
					countdownDone <- true
				}
			}
		}(countdownDone)

		go func(userDone chan<- bool) {
			mp := mc.GetProblem()
			fmt.Printf("%d %s %d\n", mp.Value1, mp.Symbol, mp.Value2)
			var input int64
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println("Input: ", input)
			solution := mp.Solution()
			if input == solution {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong! Expected Answer:", solution)
			}
			userDone <- true
		}(userDone)

		finished := false

		for !finished {
			select {
			case <-countdownDone:
				fmt.Println("countdown finished first")
				finished = true
			case <-userDone:
				fmt.Println("user finished first")
				finished = true
			}
		}
		countdowner.count = 0
	}
}
