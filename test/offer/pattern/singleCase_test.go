/*
@Time : 2020/4/27 14:46
@Author : zxr
@File : singleCase_test
@Software: GoLand
*/
package pattern

import (
	"fmt"
	"jdbook/test/offer/pattern/singlecase"
	"sync"
	"testing"
)

var wg sync.WaitGroup

//单例模式
func TestSingleCase(t *testing.T) {
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()
			singlecase.IncrAge1()
		}()

		go func() {
			defer wg.Done()
			singlecase.IncrAge2()
		}()
	}

	wg.Wait()
	s := singlecase.GetInstance()
	fmt.Println("age:", s.GetAge())
}
