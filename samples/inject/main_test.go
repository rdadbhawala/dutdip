package main

import (
	"testing"

	"github.com/karlkfi/inject"
	"github.com/rdadbhawala/dutdip/samples/inject/library"
	"github.com/rdadbhawala/dutdip/samples/inject/service"
)

func BenchmarkSingletonDependency(b *testing.B) {
	graph := inject.NewGraph()

	si := inject.NewAutoProvider(service.NewService)
	di := inject.NewProvider(library.NewDependency)
	var d library.Dependency
	graph.Define(&d, di)

	for i := 0; i < b.N; i++ {
		var s service.Service
		graph.Define(&s, si)
		graph.Resolve(&s)
	}
}

func BenchmarkLibrary(b *testing.B) {
	graph := inject.NewGraph()

	di := inject.NewProvider(library.NewDependency)
	for i := 0; i < b.N; i++ {
		var d library.Dependency
		graph.Define(&d, di)
		si := inject.NewProvider(service.NewService, &d)
		var s service.Service
		graph.Define(&s, si)
		graph.Resolve(&s)
	}
}