package container

import (
	"fmt"
	"math/rand"
	"time"
)

type Clonable interface {
	Clone() Clonable
}

type Container struct {
	ID       string
	Name     string
	metadata [3]string
}

func NewContainer(name string) *Container {
	rand.Seed(time.Now().UnixNano())
	containerID := fmt.Sprintf("%v", rand.Int())
	return &Container{
		ID:       containerID,
		Name:     fmt.Sprintf("%s-clone", name),
		metadata: [3]string{time.Now().UTC().String(), name, fmt.Sprintf("%s-%v", name, containerID)},
	}
}

func (c *Container) Clone() Clonable {
	return &Container{
		ID:       c.ID,
		Name:     c.Name,
		metadata: c.metadata,
	}
}
