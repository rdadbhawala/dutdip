package main

import (
	"testing"

	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
)

func BenchmarkLibrary(b *testing.B) {
	model.F = &factory{}
	for i := 0; i < b.N; i++ {
		model.F.NewService()
	}
}

func BenchmarkSingleton(b *testing.B) {
	model.F = &factorySi{}
	for i := 0; i < b.N; i++ {
		model.F.NewService()
	}
}

// func BenchmarkServiceDep(b *testing.B) {
// 	s := &service.FactoryImplDep{D: &library.DependencyFactoryImpl{}}
// 	for i := 0; i < b.N; i++ {
// 		s.NewService()
// 	}
// }

// func BenchmarkSeparateFactories(b *testing.B) {
// 	model.S = &service.FactoryImplSep{}
// 	model.D = &library.DependencyFactoryImpl{}
// 	for i := 0; i < b.N; i++ {
// 		model.S.NewService()
// 	}
// }
