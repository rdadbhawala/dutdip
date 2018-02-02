package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct {
	f *model.Functions
}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	s.f.NewDependency().Operation()
	s.f.NewAnotherDep().AnotherOp()
}

func NewService(f *model.Functions) model.Service {
	return &ServiceImpl{f}
}
