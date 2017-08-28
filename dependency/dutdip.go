package dependency

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// DataAccessLayerFactoryImpl ...
type DataAccessLayerFactoryImpl struct{}

// CreateDal returns a new instance of DAL
func (df DataAccessLayerFactoryImpl) CreateDal() model.DataAccessLayer {
	fmt.Println("\tDataAccessLayer Initialization")
	return &dalImpl{}
}

// AnotherDalFactoryImpl ...
type AnotherDalFactoryImpl struct{}

// CreateAnotherDal returns a new instance of AnotherDal
func (AnotherDalFactoryImpl) CreateAnotherDal() model.AnotherDal {
	fmt.Println("\tAnotherDal Initialization")
	return &anotherDalImpl{}
}
