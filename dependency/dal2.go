package dependency

import "fmt"

type anotherDalImpl struct{}

func (ad *anotherDalImpl) DalMethod1() {
	defer fmt.Println("\t\tDalMethod1 End")
	fmt.Println("\t\tDalMethod1 Start")
}
