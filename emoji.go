package main

import "fmt"

const hrule = "----------------------------------------------------------------"

func main() {
	fmt.Println("😀")
	fmt.Println(hrule)

	for r := 0x1F600; r <= 0x1F64F; r++ {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	fmt.Println(hrule)

}
