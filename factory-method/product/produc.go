package product

import "fmt"

// We create the interface that
type product interface {
	create() product
}

func buildProduct(p product) product {
	return p.create()
}

type laptopFactory interface {
	deliver(location string) product
}

func sendProduct(location string) {
	fmt.Println(p.deliver(location))
}

