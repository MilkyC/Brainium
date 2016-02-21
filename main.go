package main

import (
	"fmt"
	"github.com/MilkyC/Brainium/challenges"
	"github.com/MilkyC/Brainium/db"
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
	mc := challenges.MakeMathChallenge(seed)
	mr := db.NewMathResult()
	for true {
		mp := mc.GetProblem()
		userDone := make(chan int64, 1)
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

		go func(userDone chan<- int64) {
			fmt.Printf("%d %s %d\n", mp.Value1, mp.Symbol, mp.Value2)
			var input int64
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println("Input: ", input)
			userDone <- input
		}(userDone)

		finished := false
		for !finished {
			select {
			case <-countdownDone:
				fmt.Println("countdown finished first")
				finished = true
			case answer := <-userDone:
				fmt.Println("user finished first")
				solution := mp.Solution()
				if answer == solution {
					fmt.Println("Correct!")
				} else {
					fmt.Println("Wrong! Expected Answer:", solution)
				}
				finished = true
				mr.Save((answer == solution), countdowner.count, answer)
			}
		}
		countdowner.count = 0
	}
}
