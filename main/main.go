package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	ff := &superFactoryImpl{}
	b := service.NewBusinessService(ff)
	b.BusinessMethod1(true)
	fmt.Println()
	b = service.NewBusinessService(ff)
	b.BusinessMethod1(false)
}
