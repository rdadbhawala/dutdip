package library

import (
	"fmt"

	"github.com/rdadbhawala/dutdip/samples/journey/model"
)

type AnotherDepImpl struct{}

func (a *AnotherDepImpl) AnotherOp() {
	fmt.Println("\t\tAnotherDep.AnotherOp")
}

func NewAnotherDep() model.AnotherDep {
	return &AnotherDepImpl{}
}
