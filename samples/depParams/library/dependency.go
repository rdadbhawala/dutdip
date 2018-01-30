package library

import "fmt"

// Dependency interface
type Dependency interface {
	Operation()
	Silent()
}

// NewDependency ...
func NewDependency() Dependency {
	return &dependencyImpl{}
}

type dependencyImpl struct{}

func (d *dependencyImpl) Operation() {
	fmt.Println("dependencyImpl.Operation")
	//time.Sleep(10 * time.Millisecond)
}

func (d *dependencyImpl) Silent() {
}
