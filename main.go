package main

import (
    "github.com/MilkyC/Brainium/challenges"
    "fmt"
    "time"
)

type countdown struct {
    count int
    ticker *time.Ticker
}

func main() {
    countdowner := countdown{
        count: 0,
        ticker: time.NewTicker(time.Millisecond * 1000),
    }
    seed := int64(time.Now().Unix())
    mc := challenges.NewMathChallenge(seed)

    for true {
        user_done := make(chan bool, 1)
        countdown_done := make(chan bool, 1)
        go func(countdown_done chan<- bool) {
            for range countdowner.ticker.C {
                countdowner.count += 1
                fmt.Printf("%d seconds\n", countdowner.count)
                if countdowner.count > 10 {
                    countdown_done <- true
                }
            }
        }(countdown_done)

        go func(user_done chan<- bool) {
            mp := mc.GetProblem()
            fmt.Printf("%d %s %d\n", mp.Value1, mp.Symbol, mp.Value2)
            var input int64
            _, err := fmt.Scanln(&input)
            if err != nil {
                 fmt.Println("Error: ", err)
            }
            fmt.Println("Input: ", input)
            solution := mp.Solution()
            if input == solution{
                fmt.Println("Correct!")
            } else {
                fmt.Println("Wrong! Expected Answer:", solution)
            }
            user_done <- true
        }(user_done)

        finished := false

        for !finished {
            select {
            case <-countdown_done:
                fmt.Println("countdown finished first")
                finished = true
            case <-user_done:
                fmt.Println("user finished first")
                finished = true
            }
        }
        countdowner.count = 0
    }
}
