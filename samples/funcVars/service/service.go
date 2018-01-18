package service

import "github.com/rdadbhawala/dutdip/samples/funcVars/library"

// Service interface is a service
type Service interface {
	Feature()
}

// NewService ...
var NewService = func(deps library.Dependency) Service {
	return &serviceImpl{deps}
}

type serviceImpl struct {
	d library.Dependency
}

func (s *serviceImpl) Feature() {
	s.d.Operation()
}
