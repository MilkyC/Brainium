package main

import (
    "fmt"
    "time"
)


func main() {
    count := 0
    ticker := time.NewTicker(time.Millisecond * 1000)
    user_done := make(chan bool, 1)
    countdown_done := make(chan bool, 1)

    go func(countdown_done chan<- bool) {
        for range ticker.C {
            count += 1
            fmt.Printf("%d seconds\n", count)
            if count > 10 {
                countdown_done <- true
            }
        }
    }(countdown_done)

    go func(user_done chan<- bool) {
        fmt.Scanln()
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
}
