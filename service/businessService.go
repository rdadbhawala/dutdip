package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// NewBusinessService returns an instance of BusinessService
func NewBusinessService(ff model.SuperFactory) model.BusinessService {
	fmt.Println("BusinessService Initialization")
	return &businessServiceImpl{
		SuperFac: ff,
		// AnotherDalFunc: ff.NewAnotherDal,
	}
}

type businessServiceImpl struct {
	SuperFac model.SuperFactory
	// AnotherDalFunc model.FuncNewAnotherDal
}

func (b *businessServiceImpl) BusinessMethod1(callDal bool) {
	defer fmt.Println("BusinessMethod1 End")
	fmt.Println("BusinessMethod1 Start")
	if callDal {
		dal := b.SuperFac.CreateDal()
		dal.DataMethod1()
		// } else {
		// 	adal := b.AnotherDalFunc()
		// 	adal.DalMethod1()
	}
}
