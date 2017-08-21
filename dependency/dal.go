package dependency

import "fmt"

type dalImpl struct{}

func (d *dalImpl) DataMethod1() {
	defer fmt.Println("\t\tDataMethod1 End")
	fmt.Println("\t\tDataMethod1 Start")
}
