package dependency

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

type AnotherDalFactoryImpl struct{}

// CreateAnotherDal returns a new instance of AnotherDal
func (AnotherDalFactoryImpl) CreateAnotherDal() model.AnotherDal {
	fmt.Println("\tAnotherDal Initialization")
	return &anotherDalImpl{}
}
