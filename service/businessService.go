package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// NewBusinessService returns an instance of BusinessService
func NewBusinessService(dalFunc model.FuncNewDataAccessLayer) model.BusinessService {
	fmt.Println("BusinessService Initialization")
	return &businessServiceImpl{
		DalFunc: dalFunc,
	}
}

type businessServiceImpl struct {
	DalFunc model.FuncNewDataAccessLayer
}

func (b *businessServiceImpl) BusinessMethod1(callDal bool) {
	defer fmt.Println("BusinessMethod1 End")
	fmt.Println("BusinessMethod1 Start")
	if callDal {
		dal := b.DalFunc()
		dal.DataMethod1()
	}
}
