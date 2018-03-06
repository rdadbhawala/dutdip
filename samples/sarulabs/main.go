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
	fmt.Println("Singleton: 1 == 2", srv == srv2)

	sctx2, _ := ctx.SubContext()
	srv3 := sctx2.Get("srvImpl").(service.Service)
	fmt.Println("Singleton: 1 == 3", srv == srv3)
	fmt.Println("Singleton: 2 == 3", srv2 == srv3)

	srv4 := sctx2.Get("srvImpl").(service.Service)
	fmt.Println("Singleton: 3 == 4", srv3 == srv4)
}
