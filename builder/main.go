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

type englishBuilder struct{}

func (h *englishBuilder) buildRoof() error {
	return nil
}

func (h *englishBuilder) buildWindow() error {
	return nil
}

func (h *englishBuilder) buildWall() error {
	return nil
}

func (h *englishBuilder) buildGate() error {
	return nil
}

type russianBuilder struct{}

func (c *russianBuilder) buildRoof() error {
	return nil
}

func (c *russianBuilder) buildWindow() error {
	return nil
}

func (c *russianBuilder) buildWall() error {
	return nil
}

func (c *russianBuilder) buildGate() error {
	return nil
}

func main() {
	concreteBuilder := &englishBuilder{}
	director := newDirector(concreteBuilder)
	director.makeCastle()
}
