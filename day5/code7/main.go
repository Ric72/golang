// code7 project main.go
package main

import (
	"fmt"
)

//接口的运用

type USB interface {
	GetName() string
	Connect()
}

type PhoneConnecter struct {
	Name string
}

func (pc PhoneConnecter) GetName() string {
	return pc.Name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.Name)
}

func main() {
	a := PhoneConnecter{"Iphone"}
	var usb USB
	a.Name = "aaaa"
	usb = PhoneConnecter(a)

	usb.Connect()
	//a.Connect()
	//DisConnected(a)
}

//此时不能确定是否a就是USB类型，所以还添加一个方法
//要加入判断方法，运用ok的方法

func DisConnected(u USB) {
	if a, ok := u.(PhoneConnecter); ok {
		fmt.Println("DisConnected:", a.GetName())
		return
	}
	fmt.Println("Unknown Device")

}
