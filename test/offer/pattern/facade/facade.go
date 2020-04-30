/*
@Time : 2020/4/29 17:28
@Author : zxr
@File : facade
@Software: GoLand
*/
package facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Printf(" CarModel - SetModel\n")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Printf(" CarEngine - SetEngine\n")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Printf(" CarBody - SetBody\n")
}

type CarAccessories struct {
}

func NewCarAccessories() *CarAccessories {
	return &CarAccessories{}
}
func (c *CarAccessories) SetAccessories() {
	fmt.Printf(" CarAccessories - SetAccessories\n")
}

type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
	c.accessories.SetAccessories()
}
