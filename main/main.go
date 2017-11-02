package main

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

func main() {
	model.SetSuperFactory(&superFactoryImpl{})
	b := model.GetSuperFactory().NewBusinessService()
	b.BusinessMethod1(true)
	fmt.Println()
	b = model.GetSuperFactory().NewBusinessService()
	b.BusinessMethod1(false)
}
