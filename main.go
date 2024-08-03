package main

import "fmt"

func main() {
	fmt.Println()
}

type laptopSize float64

func (l laptopSize) getSizeOfLaptop() laptopSize {
	return l
}
