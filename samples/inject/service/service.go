package service

import "github.com/rdadbhawala/dutdip/samples/inject/library"

// Service interface is a service
type Service interface {
	Feature()
	GetDependency() library.Dependency
}

// NewService ...
func NewService(deps library.Dependency) Service {
	return &serviceImpl{deps}
}

type serviceImpl struct {
	d library.Dependency
}

func (s *serviceImpl) GetDependency() library.Dependency {
	return s.d
}

func (s *serviceImpl) Feature() {
	s.d.Operation()
}
