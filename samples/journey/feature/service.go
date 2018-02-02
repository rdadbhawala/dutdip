package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct {
	d model.NewDependency
	a model.NewAnotherDep
}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	s.d().Operation()
	s.a().AnotherOp()
}

func NewService(dep model.NewDependency, ano model.NewAnotherDep) model.Service {
	return &ServiceImpl{dep, ano}
}
