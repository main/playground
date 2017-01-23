package main

import (
	"fmt"
)

func main() {
	var msn MakeSomeNoiser
	msn = new(Dog)
	msn.MakeSomeNoise(5)
	msn = Man{Name: "Petr"}
	//	msn.Name = "Petr"
	msn.MakeSomeNoise(5)
	fmt.Println(msn.Name)
}

type MakeSomeNoiser interface {
	MakeSomeNoise(volumeLevel int)
}

var _ MakeSomeNoiser = new(Man)

type Man struct {
	Name string
}

func (man Man) MakeSomeNoise(volumeLevel int) {
	fmt.Println("Hello, my name is", man.Name)
}

var _ MakeSomeNoiser = new(Dog)

type Dog struct{}

func (dog Dog) MakeSomeNoise(volumeLevel int) {
	fmt.Println("Gav Gav gav!!!111")
}
