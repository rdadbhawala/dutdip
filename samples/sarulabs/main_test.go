package main

import (
	"fmt"
	"testing"

	"github.com/rdadbhawala/dutdip/samples/sarulabs/library"
	"github.com/rdadbhawala/dutdip/samples/sarulabs/service"
	"github.com/sarulabs/di"
)

func BenchmarkLibrary(tb *testing.B) {
	b, err := di.NewBuilder()
	if err != nil {
		fmt.Println("Error in di.NewBuilder: ", err)
		return
	}
	b.AddDefinition(di.Definition{
		Name:  "libDep",
		Scope: di.Request,
		Build: func(ctx di.Context) (interface{}, error) {
			return library.NewDependency(), nil
		},
	})
	b.AddDefinition(di.Definition{
		Name:  "srvImpl",
		Scope: di.Request,
		Build: func(ctx di.Context) (interface{}, error) {
			return service.NewService(ctx.Get("libDep").(library.Dependency)), nil
		},
	})
	ctx := b.Build()
	for i := 0; i < tb.N; i++ {
		sctx, _ := ctx.SubContext()
		_, ok := sctx.Get("srvImpl").(service.Service)
		if !ok {
			tb.Error("Error Casting")
		}
	}
}

func BenchmarkSingleton(tb *testing.B) {
	b, err := di.NewBuilder()
	if err != nil {
		fmt.Println("Error in di.NewBuilder: ", err)
		return
	}
	b.AddDefinition(di.Definition{
		Name: "libDep",
		Build: func(ctx di.Context) (interface{}, error) {
			return library.NewDependency(), nil
		},
	})
	b.AddDefinition(di.Definition{
		Name: "srvImpl",
		Build: func(ctx di.Context) (interface{}, error) {
			return service.NewService(ctx.Get("libDep").(library.Dependency)), nil
		},
	})
	ctx := b.Build()
	for i := 0; i < tb.N; i++ {
		_, ok := ctx.Get("srvImpl").(service.Service)
		if !ok {
			tb.Error("Error Casting")
		}
	}
}
