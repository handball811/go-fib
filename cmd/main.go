package main

import (
	"fmt"
	"time"

	"github.com/handball811/go-fib/internal"
)

func main() {
	startTime := time.Now()
	defer func() {
		d := time.Since(startTime)
		fmt.Println("Take Time", d.Milliseconds(), "[ms]")
	}()
	r, err := internal.Run(1000)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("answer", r)
}
