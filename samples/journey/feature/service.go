package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct{}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	dep := model.SF.NewDependency()
	defer dep.Release()
	dep.Operation()
	model.SF.NewAnotherDep().AnotherOp()
}

type ServiceFactoryImpl struct{}

func (s ServiceFactoryImpl) NewService() model.Service {
	return &ServiceImpl{}
}
