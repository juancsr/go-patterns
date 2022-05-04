package main

import "fmt"

type builder interface {
	buildRoof() error
	buildWindow() error
	buildWall() error
	buildGate() error
}

type director struct {
	builder builder
}

func newDirector(b builder) *director {
	return &director{builder: b}
}

func (d *director) changeBuilder(b builder) {
	d.builder = b
}

func (d *director) makeCastle() error {
	var err error

	if err = d.builder.buildRoof(); err != nil {
		return err
	}

	if err = d.builder.buildGate(); err != nil {
		return err
	}

	if err = d.builder.buildWall(); err != nil {
		return err
	}

	fmt.Println("casttle is done!")

	return err
}

func (d *director) makeHouse() error {
	var err error

	if err = d.builder.buildRoof(); err != nil {
		return err
	}

	if err = d.builder.buildWindow(); err != nil {
		return err
	}

	if err = d.builder.buildWall(); err != nil {
		return err
	}

	fmt.Println("house is done!")

	return err
}

type houseBuilder struct{}

func (h *houseBuilder) buildRoof() error {
	return nil
}

func (h *houseBuilder) buildWindow() error {
	return nil
}

func (h *houseBuilder) buildWall() error {
	return nil
}

func (h *houseBuilder) buildGate() error {
	return nil
}

type casttleBuilder struct{}

func (c *casttleBuilder) buildRoof() error {
	return nil
}

func (c *casttleBuilder) buildWindow() error {
	return nil
}

func (c *casttleBuilder) buildWall() error {
	return nil
}

func (c *casttleBuilder) buildGate() error {
	return nil
}

func main() {
	houseBulder := houseBuilder{}

	director := newDirector(&houseBulder)

	director.makeCastle()
}
