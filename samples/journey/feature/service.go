package feature

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type ServiceImpl struct {
	d model.Dependency
}

func (s *ServiceImpl) Function() {
	fmt.Println("\tService.Function")
	s.d.Operation()
}

func NewService(dep model.Dependency) model.Service {
	return &ServiceImpl{dep}
}
