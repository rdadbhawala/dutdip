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

	var s service.Service
	var d library.Dependency
	graph.Define(&s, si)
	graph.Define(&d, di)
	graph.Resolve(&s)
	s.Feature()

	// var s2 service.Service
	// graph.Define(&s2, si)
	// graph.Resolve(&s2)
	// fmt.Println("Singleton S: ", s == s2)
	// fmt.Println("Singleton S.D: ", s.GetDependency() == s2.GetDependency())
}
