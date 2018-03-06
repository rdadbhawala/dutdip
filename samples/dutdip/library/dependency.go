package library

import (
	"fmt"
	"sync"

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

type DependencySingleton struct {
	o sync.Once
	d *dependencyImpl
}

func (ds DependencySingleton) NewDependency() model.Dependency {
	ds.o.Do(func() {
		ds.d = &dependencyImpl{}
	})
	return ds.d
}
