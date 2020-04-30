/*
@Time : 2020/4/8 14:46
@Author : zxr
@File : 4.14
@Software: GoLand
*/
package array

//如何求集合的所有子集
func GetAllSybsetSub(str string) (arr []string) {
	arr = make([]string, 0)
	if str == "" {
		return
	}
	arr = append(arr, string(str[0]))
	for i := 1; i < len(str); i++ {
		ll := len(arr)
		for j := 0; j < ll; j++ {
			arr = append(arr, arr[j]+string(str[i]))
		}
		arr = append(arr, string(str[i:i+1]))
	}
	return
}
