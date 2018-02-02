package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/feature"
	"github.com/rdadbhawala/dutdip/samples/journey/library"
)

func main() {
	fmt.Println("Main.Main")
	f := &mainFactory{}
	feature.NewService(f).Function()
}

type mainFactory struct {
	library.DependencyFactoryImpl
	library.AnotherDepFactoryImpl
}
