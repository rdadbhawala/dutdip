package main

import (
	"testing"

	"github.com/rdadbhawala/dutdip/samples/funcVars/library"
	"github.com/rdadbhawala/dutdip/samples/funcVars/service"
)

func BenchmarkLibrary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.NewService(library.NewDependency())
	}
}

func BenchmarkSingleton(b *testing.B) {
	service.SingletonSetup(library.NewDependency())
	for i := 0; i < b.N; i++ {
		service.NewSingletonService()
	}
}
