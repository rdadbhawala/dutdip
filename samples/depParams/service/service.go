package service

import (
	"sync"

	"github.com/rdadbhawala/dutdip/samples/depParams/library"
)

// Service interface is a service
type Service interface {
	Feature()
}

// NewService ...
func NewService(deps library.Dependency) Service {
	return &serviceImpl{deps}
}

type serviceImpl struct {
	d library.Dependency
}

func (s *serviceImpl) Feature() {
	s.d.Operation()
}

var singletonService Service
var sso sync.Once

// SingletonSetup ...
func SingletonSetup(deps library.Dependency) {
	sso.Do(func() {
		singletonService = &serviceImpl{deps}
	})
}

// NewSingletonService ...
func NewSingletonService() Service {
	return singletonService
}
