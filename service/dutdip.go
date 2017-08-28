package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// BusinessServiceFactoryImpl ...
type BusinessServiceFactoryImpl struct{}

// NewBusinessService returns an instance of BusinessService
func (BusinessServiceFactoryImpl) NewBusinessService(ff model.BusinessServiceDependencies) model.BusinessService {
	fmt.Println("BusinessService Initialization")
	return &businessServiceImpl{
		SuperFac: ff,
	}
}
