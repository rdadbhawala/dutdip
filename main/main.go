package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/dependency"
	"github.com/rdadbhawala/dutdip/model"
	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	ff := &model.FunctionFactory{
		NewDataAccessLayer: dependency.NewDataAccessLayer,
		NewAnotherDal:      dependency.NewAnotherDal,
	}
	b := service.NewBusinessService(ff)
	b.BusinessMethod1(true)
	fmt.Println()
	b = service.NewBusinessService(ff)
	b.BusinessMethod1(false)
}
