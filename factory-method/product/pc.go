package product

import "fmt"

type pc struct {
	brand string
}

func (p *pc) build() string {
	fmt.Sprintf(`Building a %v PC`, p.brand)
	return 	fmt.Sprintf("Delivering the PC to %v", location)
}

type PCFactory struct {}

func (factory *PCFactory) deliver(location string) pc {

}

