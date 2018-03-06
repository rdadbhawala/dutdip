package service

import (
	"sync"

	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
)

// // NewService ...
// func NewService(deps model.Dependency) model.Service {
// 	return &serviceImpl{deps}
// }

type serviceImpl struct {
	d model.Dependency
}

func (s *serviceImpl) Feature() {
	s.d.Operation()
}

// FactoryImpl ...
type FactoryImpl struct{}

// NewService ...
func (s *FactoryImpl) NewService() model.Service {
	return &serviceImpl{model.F.NewDependency()}
}

// // FactoryImplDep ...
// type FactoryImplDep struct {
// 	D model.DepFactory
// }

// // NewService ...
// func (f *FactoryImplDep) NewService() model.Service {
// 	return &serviceImpl{f.D.NewDependency()}
// }

// // FactoryImplSep ...
// type FactoryImplSep struct{}

// // NewService ...
// func (fs *FactoryImplSep) NewService() model.Service {
// 	return &serviceImpl{model.D.NewDependency()}
// }

// FactoryImplSingleton ...
type FactoryImplSingleton struct {
	s *serviceImpl
	o sync.Once
}

// NewService ...
func (si *FactoryImplSingleton) NewService() model.Service {
	si.o.Do(func() {
		si.s = &serviceImpl{model.F.NewDependency()}
	})
	return si.s
}
