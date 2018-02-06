package library

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type DependencyImpl struct{}

func (d *DependencyImpl) Operation() {
	fmt.Println("\t\tDependency.Operation")
}

func (d *DependencyImpl) Release() {
	fmt.Println("\t\tDependency.Release")
}

type DependencyFactoryImpl struct{}

func (d DependencyFactoryImpl) NewDependency() model.Dependency {
	return &DependencyImpl{}
}
