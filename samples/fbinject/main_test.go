package main

import (
	"strconv"
	"testing"

	"github.com/facebookgo/inject"
	"github.com/rdadbhawala/dutdip/samples/fbinject/library"
	"github.com/rdadbhawala/dutdip/samples/fbinject/service"
)

func BenchmarkLibrary(b *testing.B) {
	var g inject.Graph
	if err := g.Provide(&inject.Object{Value: library.NewDependency()}); err != nil {
		b.Error("err in g.Provide: ", err)
		return
	}
	for i := 0; i < b.N; i++ {
		var s service.Impl
		if err := g.Provide(&inject.Object{Value: &s, Name: "s" + strconv.Itoa(i)}); err != nil {
			b.Error("err in g.Provide: ", err)
			break
		}
		if err := g.Populate(); err != nil {
			b.Error("err in g.Populate: ", err)
			break
		}
	}
}
