package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/dependency"
	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	b := service.NewBusinessService(dependency.NewDataAccessLayer(), dependency.NewAnotherDal())
	b.BusinessMethod1(true)
	fmt.Println()
	b = service.NewBusinessService(dependency.NewDataAccessLayer(), dependency.NewAnotherDal())
	b.BusinessMethod1(false)
}
