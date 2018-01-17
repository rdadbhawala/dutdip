package main

import (
	"github.com/karlkfi/inject"
	"github.com/rdadbhawala/dutdip/samples/inject/library"
	"github.com/rdadbhawala/dutdip/samples/inject/service"
)

func main() {
	graph := inject.NewGraph()

	si := inject.NewAutoProvider(service.NewService)
	di := inject.NewProvider(library.NewDependency)

	// resolve a and all its (transitive) dependencies
	var s service.Service
	var d library.Dependency
	graph.Define(&s, si)
	graph.Define(&d, di)
	graph.Resolve(&s)
	s.Feature()
}
