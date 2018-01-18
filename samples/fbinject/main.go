package main

import (
	"fmt"

	"github.com/facebookgo/inject"
	"github.com/rdadbhawala/dutdip/samples/fbinject/library"
	"github.com/rdadbhawala/dutdip/samples/fbinject/service"
)

func main() {
	var g inject.Graph
	var s service.Impl
	if err := g.Provide(&inject.Object{Value: library.NewDependency()},
		&inject.Object{Value: &s}); err != nil {
		fmt.Println("err in g.Provide: ", err)
	}
	if err := g.Populate(); err != nil {
		fmt.Println("err in g.Populate: ", err)
	}
	s.Feature()

	var s2 service.Impl
}
