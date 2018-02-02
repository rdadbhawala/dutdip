package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct {
	f model.SuperFactory
}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	s.f.NewDependency().Operation()
	s.f.NewAnotherDep().AnotherOp()
}

type ServiceFactoryImpl struct{}

func (s ServiceFactoryImpl) NewService(f model.SuperFactory) model.Service {
	return &ServiceImpl{f}
}
