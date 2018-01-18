package main

import (
	"github.com/rdadbhawala/dutdip/samples/funcVars/library"
	"github.com/rdadbhawala/dutdip/samples/funcVars/service"
)

func main() {
	service.NewService(library.NewDependency()).Feature()
}
