package brate

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	rate     int           // 每分钟的请求数限制
	interval time.Duration // 时间间隔
	count    int           // 当前计数
	lock     sync.Mutex    // 互斥锁
	finish   bool          //
}

func NewRateLimiter(rate int, d time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:     rate,
		interval: d,
		count:    0,
		lock:     sync.Mutex{},
	}
}

func (rl *RateLimiter) ReSetCount() {
	for {
		if rl.finish {
			fmt.Println("finish..............")
			break
		}
		time.Sleep(rl.interval)
		fmt.Println("--------------origin count--------------: ", rl.count)
		// rl.lock.Lock()
		// defer rl.lock.Unlock()
		rl.count = 0
	}
}

func (rl *RateLimiter) Allow() bool {
	fmt.Println("allow count...............: ", rl.count)
	rl.lock.Lock()
	fmt.Println("wait lock.................")
	defer rl.lock.Unlock()

	if rl.count >= rl.rate {
		fmt.Println("allow false.................")
		return false
	}
	rl.count++
	fmt.Println("allow true.....................")
	return true
}
