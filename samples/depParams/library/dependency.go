package library

import "fmt"

// Dependency interface
type Dependency interface {
	Operation()
}

func NewDependency() Dependency {
	return &dependencyImpl{}
}

type dependencyImpl struct{}

func (d *dependencyImpl) Operation() {
	fmt.Println("dependencyImpl.Operation")
}
