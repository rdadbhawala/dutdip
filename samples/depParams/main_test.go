package main

import (
	"testing"

	"github.com/rdadbhawala/dutdip/samples/depParams/library"
	"github.com/rdadbhawala/dutdip/samples/depParams/service"
)

func BenchmarkService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.NewService(library.NewDependency())
	}
	// s.Feature()
}

func BenchmarkCall(b *testing.B) {
	s := service.NewService(library.NewDependency())
	for i := 0; i < b.N; i++ {
		s.Shh()
	}
}

func BenchmarkSingleton(b *testing.B) {
	service.SingletonSetup(library.NewDependency())
	for i := 0; i < b.N; i++ {
		service.NewSingletonService()
	}
}
