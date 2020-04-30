/*
@Time : 2020/4/27 14:54
@Author : zxr
@File : case
@Software: GoLand
*/
package singlecase

import (
	"sync"
)

//单例模式
type Single struct {
	age  int
	name string
	mux  sync.RWMutex
}

var (
	P    *Single
	once sync.Once
)

//单例模式
func init() {
	once.Do(func() {
		P = &Single{}
	})
}

func GetInstance() *Single {
	return P
}

func (s *Single) setName(name string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.name = name
}

func (s *Single) IncrAge() {
	s.mux.RLock()
	defer s.mux.RUnlock()
	s.age++
}

func (s *Single) getName() string {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.name
}

func (s *Single) GetAge() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.age
}
