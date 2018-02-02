package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/feature"
	"github.com/rdadbhawala/dutdip/samples/journey/library"
)

func main() {
	fmt.Println("Main.Main")
	feature.NewService(library.NewDependency).Function()
}
