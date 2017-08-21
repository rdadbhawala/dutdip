package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	af := GetAllFactory()
	ff := &af
	b := service.NewBusinessService(ff)
	b.BusinessMethod1(true)
	fmt.Println()
	b = service.NewBusinessService(ff)
	b.BusinessMethod1(false)
}
