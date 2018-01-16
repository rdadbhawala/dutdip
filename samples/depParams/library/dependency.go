package library

import (
	"time"
)

// Dependency interface
type Dependency interface {
	Operation()
}

// NewDependency ...
func NewDependency() Dependency {
	return &dependencyImpl{}
}

type dependencyImpl struct{}

func (d *dependencyImpl) Operation() {
	//fmt.Println("dependencyImpl.Operation")
	time.Sleep(10 * time.Millisecond)
}
