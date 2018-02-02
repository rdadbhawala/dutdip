package library

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

func NewDependency() model.Dependency {
	return &DependencyImpl{}
}

type DependencyImpl struct{}

func (d *DependencyImpl) Operation() {
	fmt.Println("\t\tDependency.Operation")
}
