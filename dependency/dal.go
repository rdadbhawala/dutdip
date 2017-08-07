package dependency

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

func NewDataAccessLayer() model.DataAccessLayer {
	return &dalImpl{}
}

type dalImpl struct{}

func (d *dalImpl) DataMethod1() {
	defer fmt.Println("\tDataMethod1 End")
	fmt.Println("\tDataMethod1 Start")
}
