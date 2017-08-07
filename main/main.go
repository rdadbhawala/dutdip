package main

import (
	"github.com/rdadbhawala/dutdip/dependency"
	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	b := service.NewBusinessService(dependency.NewDataAccessLayer())
	b.BusinessMethod1()
}
