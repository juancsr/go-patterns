package main

import (
	"fmt"
	"prototype/container"
)

func main() {
	container := container.NewContainer("MySQL")
	fmt.Println(container)
	containerClone := container.Clone()
	fmt.Println(containerClone)
	fmt.Println(container == containerClone)
}
