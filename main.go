package main

import (
	"fmt"

	ac "practice/actions"
	dv "practice/devices"
)

func main(){
	var p dv.Pc = dv.Pc{
		Cpu: dv.Cpu{
			Prod: "Intel",
			Core: 8,
		},
	}
	p.Off()
	side := ac.Side{14813217673}
	fmt.Printf("Area of square: %d\n", side.Area())
	fmt.Printf("Perimeter of square: %d\n", side.Perimeter())
	fmt.Println("Faiza")
}