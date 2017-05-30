package dependency

import "fmt"

type DataAccessLayer struct{}

func (d *DataAccessLayer) DataMethod1() {
	defer fmt.Println("\tDataMethod1 End")
	fmt.Println("\tDataMethod1 Start")
}
