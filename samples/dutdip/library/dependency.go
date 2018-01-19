package library

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
)

// NewDependency ...
func NewDependency() model.Dependency {
	return &dependencyImpl{}
}

type dependencyImpl struct{}

func (d *dependencyImpl) Operation() {
	fmt.Println("dependencyImpl.Operation")
	//time.Sleep(10 * time.Millisecond)
}

// DependencyFactoryImpl ...
type DependencyFactoryImpl struct {
}

// NewDependency ...
func (d *DependencyFactoryImpl) NewDependency() model.Dependency {
	return &dependencyImpl{}
}
