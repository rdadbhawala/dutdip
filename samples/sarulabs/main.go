package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/sarulabs/library"
	"github.com/rdadbhawala/dutdip/samples/sarulabs/service"
	"github.com/sarulabs/di"
)

func main() {
	// service.NewService(library.NewDependency()).Feature()

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
	sctx, _ := ctx.SubContext()
	srv := sctx.Get("srvImpl").(service.Service)
	srv.Feature()

	srv2 := sctx.Get("srvImpl").(service.Service)
	fmt.Println("Singleton: ", srv == srv2)

	sctx, _ = ctx.SubContext()
	srv3 := sctx.Get("srvImpl").(service.Service)
	fmt.Println("Singleton: ", srv == srv3)
	fmt.Println("Singleton: ", srv2 == srv3)

	srv4 := sctx.Get("srvImpl").(service.Service)
	fmt.Println("Singleton: ", srv3 == srv4)
}
