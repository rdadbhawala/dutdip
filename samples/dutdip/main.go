package main

import (
	"github.com/rdadbhawala/dutdip/samples/dutdip/library"
	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
	"github.com/rdadbhawala/dutdip/samples/dutdip/service"
)

func main() {
	// service.NewService(library.NewDependency()).Feature()

	model.F = &factory{}
	model.F.NewService().Feature()
}

type factory struct {
	library.DependencyFactoryImpl
	service.FactoryImpl
}

type factorySi struct {
	library.DependencyFactoryImpl
	service.FactoryImplSingleton
}
