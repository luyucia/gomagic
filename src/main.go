package main

import (
	"bitset"
	"fmt"
)

func main() {

	a := [6]byte{}

	a[0] = 1

	bit := bitset.New(1440)

	bit.Set(0)
	bit.Set(2)
	bit.Set(7)
	bit.Set(8)
	bit.Set(9)
	bit.Set(999)
	bit.Set(1339)
	//bit.Set(1)
	//bit.Set(8)
	//bit.Set(9)
	//bit.Set(16)

	//t := 17
	//b := t>>3
	//c := t%8
	//fmt.Print(b)
	//fmt.Print(c)

	fmt.Println(bit.Get(0))
	fmt.Println(bit.Get(1))
	fmt.Println(bit.Get(2))
	fmt.Println(bit.Get(3))
	fmt.Println(bit.Get(4))
	fmt.Println(bit.Get(5))
	fmt.Println(bit.Get(6))
	fmt.Println(bit.Get(7))
	fmt.Println(bit.Get(8))
	fmt.Println(bit.Get(9))
	fmt.Println(bit.Get(10))
	fmt.Println(bit.Get(999))
	fmt.Println(bit.Get(1339))

}
