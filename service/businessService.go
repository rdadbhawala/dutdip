package service

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// NewBusinessService returns an instance of BusinessService
func NewBusinessService(dal model.DataAccessLayer) model.BusinessService {
	return &businessServiceImpl{
		Dal: dal,
	}
}

type businessServiceImpl struct {
	Dal model.DataAccessLayer
}

func (b *businessServiceImpl) BusinessMethod1() {
	defer fmt.Println("BusinessMethod1 End")
	fmt.Println("BusinessMethod1 Start")
	b.Dal.DataMethod1()
}
