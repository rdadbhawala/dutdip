package main

import (
	"github.com/rdadbhawala/dutdip/dependency"
	"github.com/rdadbhawala/dutdip/service"
)

func main() {
	b := &service.BusinessService{
		Dal: &dependency.DataAccessLayer{},
	}
	b.BusinessMethod1()
}
