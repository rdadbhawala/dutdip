package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct {
	d model.NewDependency
}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	s.d().Operation()
}

func NewService(dep model.NewDependency) model.Service {
	return &ServiceImpl{dep}
}
