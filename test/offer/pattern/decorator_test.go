/*
@Time : 2020/4/23 15:51
@Author : zxr
@File : decorator
@Software: GoLand
*/
package pattern

import (
	"fmt"
	"log"
	"math"
	"os"
	"testing"
	"time"
)

type piFunc func(int) float64

//装饰器模式
func TestDecorator(t *testing.T) {
	pi := Pi(10000)
	fmt.Println(pi)

	logger := wrapLogger(Pi, log.New(os.Stdout, "test", log.LstdFlags))
	f := logger(10000)
	fmt.Println("piw:", f)

	pi2 := wrapLogger2(Pi, log.New(os.Stdout, "pre2:", log.LstdFlags))
	i := pi2(100000)
	fmt.Println("pi2:", i)
}

//装饰器模式
func wrapLogger(fun piFunc, logger *log.Logger) piFunc {
	return func(i int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v,n=%v,result=%v", time.Since(t), i, result)
			}(time.Now())
			return fun(n)
		}
		return fn(i)
	}
}

//装饰器模式
func wrapLogger2(fun piFunc, logger *log.Logger) piFunc {
	return func(i int) (result float64) {
		defer func(t time.Time) {
			logger.Printf("took=%v,n=%v,result=%v", time.Since(t), i, result)
		}(time.Now())
		return fun(i)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}
