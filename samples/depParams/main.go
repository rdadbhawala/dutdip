package main

import (
	"github.com/rdadbhawala/dutdip/samples/depParams/library"
	"github.com/rdadbhawala/dutdip/samples/depParams/service"
)

func main() {
	s := service.NewService(library.NewDependency())
	s.Feature()
}
