package main

import "fmt"

const ebulicaoC = 100

func main() {

	tempC := ebulicaoC
	tempK := (tempC + 273)
	fmt.Printf("A temperatura de ebulição da água em °C = %v , temperatura de ebulição da água em °K =%v.", tempC, tempK)
}
