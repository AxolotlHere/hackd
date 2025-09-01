package main

import "fmt"

func main() {
	initDbController()
	fmt.Println("Hellow")
	//fmt.Println(createUser("axolotl", "Im him", "snaveenbharath2007@gmail.com", "Py2nb-0513"))
	fmt.Println(validateUser("snaveenbharath2007@gmail.com", "Py2nb-05131"))
}
