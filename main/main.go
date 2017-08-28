package main

import (
	"fmt"
)

func main() {
	ff := &superFactoryImpl{}
	b := ff.NewBusinessService(ff)
	b.BusinessMethod1(true)
	fmt.Println()
	b = ff.NewBusinessService(ff)
	b.BusinessMethod1(false)
}
