package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// goroutine
func test() (res string) {
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// wg.Wait()

	// var a chan int
	// var wg sync.WaitGroup
	// a = make(chan int)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	x := <-a
	// 	fmt.Println("goroutine", x)
	// }()
	// a <- 10
	// fmt.Println("发送")
	// wg.Wait()
	// close(a)

	var wg sync.WaitGroup
	var a chan int
	a = make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i == 5 {
				rand.Seed(time.Now().UnixNano())
				x := rand.Intn(10)
				if x < 5 {
					time.Sleep(time.Second)
				}
				a <- i - 1
			}
		}(i)
	}
	select {
	case <-time.After(time.Microsecond):
		res = "timeout"
	case x := <-a:
		if x == 4 {
			res = "match"
		}
	}
	wg.Wait()
	return
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		res := test()
		c.String(200, res)
		// c.JSON(200, gin.H{
		// 	"message": "hello world",
		// })
	})

	r.Run(":8000")
}
