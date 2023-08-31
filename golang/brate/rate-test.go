package brate

import (
	"fmt"
	"time"
)

func RateTest() {
	rateLimiter := NewRateLimiter(5, 10*time.Second)
	go rateLimiter.ReSetCount()

	for i := 0; i < 40; i++ {
		for {
			if rateLimiter.Allow() {
				fmt.Println("allow..............")
				break
			}
			fmt.Println("not allow.................")
		}
		fmt.Println("time: ", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
	}
	rateLimiter.finish = true
}
