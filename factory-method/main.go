package main

import "fmt"

func main() {
	fmt.Println("welcom!")
	pc := pc{ brand: "Dell" }
	sendProduct(&pc, "Tokyo")
}
