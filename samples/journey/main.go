package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/feature"
	"github.com/rdadbhawala/dutdip/samples/journey/library"
	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

func main() {
	fmt.Println("Main.Main")
	model.SF = &mainFactory{}
	model.SF.NewService().Function()
}

type mainFactory struct {
	// library.DependencyFactoryImpl
	library.SingletonDependency
	library.AnotherDepFactoryImpl
	feature.ServiceFactoryImpl
}
