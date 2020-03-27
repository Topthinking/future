package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
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
	r, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer r.Close()

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		res := test()
		r.Do("incr", "hits")
		number, _ := redis.String(r.Do("GET", "hits"))
		c.String(200, res+number)
		// c.JSON(200, gin.H{
		// 	"message": "hello world",
		// })
	})

	route.Run(":8000")
}
