/*
@Time : 2020/4/27 16:04
@Author : zxr
@File : mac
@Software: GoLand
*/
package factory

//工厂模式
type Animal interface {
	GetName() string
}

type Dog struct {
}
type Cat struct {
}

func (d *Dog) GetName() string {
	return "my is dog"
}

func (c *Cat) GetName() string {
	return "my is cat"
}

//工厂模式
func GetAnimalName(animal string) Animal {
	switch animal {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	}
	return nil
}
