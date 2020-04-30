/*
@Time : 2020/4/27 15:49
@Author : zxr
@File : factory_test
@Software: GoLand
*/
package pattern

import (
	"fmt"
	"jdbook/test/offer/pattern/factory"
	"testing"
)

//工厂模式
func TestFactory(t *testing.T) {
	fmt.Println(factory.GetAnimalName("dog").GetName())
	fmt.Println(factory.GetAnimalName("cat").GetName())
}
