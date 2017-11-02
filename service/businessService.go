package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

type businessServiceImpl struct{}

func (b *businessServiceImpl) BusinessMethod1(callDal bool) {
	defer fmt.Println("BusinessMethod1 End")
	fmt.Println("BusinessMethod1 Start")
	if callDal {
		dal := model.GetSuperFactory().CreateDal()
		dal.DataMethod1()
	} else {
		adal := model.GetSuperFactory().CreateAnotherDal()
		adal.DalMethod1()
	}
}
