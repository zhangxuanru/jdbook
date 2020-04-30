/*
@Time : 2020/4/22 19:03
@Author : zxr
@File : queue
@Software: GoLand
*/
package t

import (
	"fmt"
	"testing"
	"time"
)

//实现消息队列（多生产者，多消费者）
func TestQueue(t *testing.T) {
	ch := make(chan int)
	go producer("生产者一", ch)
	go producer("生产者二", ch)
	go consumer("消费者一", ch)
	go consumer("消费者二", ch)
	time.Sleep(10 * time.Second)
	close(ch)
}

func consumer(cname string, ch <-chan int) {
	for v := range ch {
		fmt.Println(cname, "----", v)
	}
	fmt.Println(cname, " close ch")
}

func producer(cname string, ch chan<- int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}
}
