package dependency

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// NewDataAccessLayer returns a new instance of DAL
func NewDataAccessLayer() model.DataAccessLayer {
	fmt.Println("\tDataAccessLayer Initialization")
	return &dalImpl{}
}

type dalImpl struct{}

func (d *dalImpl) DataMethod1() {
	defer fmt.Println("\t\tDataMethod1 End")
	fmt.Println("\t\tDataMethod1 Start")
}
