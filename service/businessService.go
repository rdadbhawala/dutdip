package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/dependency"
)

type BusinessService struct {
	Dal *dependency.DataAccessLayer
}

func (b *BusinessService) BusinessMethod1() {
	defer fmt.Println("BusinessMethod1 End")
	fmt.Println("BusinessMethod1 Start")
	b.Dal.DataMethod1()
}
