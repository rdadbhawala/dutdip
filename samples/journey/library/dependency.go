package library

import (
	"fmt"
	"sync"

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

type SingletonDependency struct {
	o sync.Once
	d *DependencyImpl
}

func (sd SingletonDependency) NewDependency() model.Dependency {
	sd.o.Do(func() {
		sd.d = &DependencyImpl{}
	})
	return sd.d
}
