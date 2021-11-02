package product

import "fmt"

type mac struct {
	brand string
	model string
	year uint8
}

func (m *mac) deliver(location string) string {
	fmt.Printf("Builing a Mac\n")
	fmt.Sprintf(`Building a Mac.
	Props:
		* %v
		* %v `, m.model, m.year)
	return 	fmt.Sprintf("Delivering the Mac to %v", location)
}
