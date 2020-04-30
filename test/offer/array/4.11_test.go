/*
@Time : 2020/4/7 18:30
@Author : zxr
@File : 4.11_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"net"
	"testing"
)

//如何找出数组中出现一次的数 [一个数组里，除了3个数是出现一次的，其它都是出现偶数次，找出三个数中任意一个]
func TestFindOneNum(t *testing.T) {
	arr := []int{1, 2, 4, 5, 6, 4, 2}
	num := array.FindOneNum(arr)
	fmt.Println("数组中出现一次的数:", num)

	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
