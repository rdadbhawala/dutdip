package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"

	"github.com/rdadbhawala/dutdip/samples/journey/feature"
	"github.com/rdadbhawala/dutdip/samples/journey/library"
)

func main() {
	fmt.Println("Main.Main")
	f := model.NewFunctions(library.NewDependency, library.NewAnotherDep)
	feature.NewService(f).Function()
}
