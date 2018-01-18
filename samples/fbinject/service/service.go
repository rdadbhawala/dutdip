package service

import "github.com/rdadbhawala/dutdip/samples/fbinject/library"

// Service interface is a service
type Service interface {
	Feature()
}

// NewService ...
func NewService(deps library.Dependency) Service {
	return &Impl{deps}
}

// Impl ...
type Impl struct {
	D library.Dependency `inject:""`
}

// Feature ...
func (s *Impl) Feature() {
	s.D.Operation()
}
