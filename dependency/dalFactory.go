package dependency

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/model"
)

// DataAccessLayerFactoryImpl ...
type DataAccessLayerFactoryImpl struct{}

// CreateDataAccessLayer returns a new instance of DAL
func (df DataAccessLayerFactoryImpl) CreateDal() model.DataAccessLayer {
	fmt.Println("\tDataAccessLayer Initialization")
	return &dalImpl{}
}
