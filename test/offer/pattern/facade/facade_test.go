/*
@Time : 2020/4/29 17:40
@Author : zxr
@File : facade_test
@Software: GoLand
*/
package facade

import "testing"

//门面模式
func TestCarFacade_CreateCompleteCar(t *testing.T) {
	car := NewCarFacade()
	car.CreateCompleteCar()
}
