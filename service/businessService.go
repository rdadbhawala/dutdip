package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// NewBusinessService returns an instance of BusinessService
func NewBusinessService(ff *model.FunctionFactory) model.BusinessService {
	fmt.Println("BusinessService Initialization")
	return &businessServiceImpl{
		DalFunc: ff.NewDataAccessLayer,
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
