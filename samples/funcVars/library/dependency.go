package library

import "fmt"

// Dependency interface
type Dependency interface {
	Operation()
}

// NewDependency ...
var NewDependency = func() Dependency {
	return &dependencyImpl{}
}

type dependencyImpl struct{}

func (d *dependencyImpl) Operation() {
	fmt.Println("dependencyImpl.Operation")
	//time.Sleep(10 * time.Millisecond)
}
