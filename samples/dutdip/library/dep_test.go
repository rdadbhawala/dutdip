package library

import (
	"testing"

	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
)

func BenchmarkDepFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDependency()
	}
}

func BenchmarkDepFac(b *testing.B) {
	d := &DependencyFactoryImpl{}
	for i := 0; i < b.N; i++ {
		d.NewDependency()
	}
}

func BenchmarkDepInf(b *testing.B) {
	var d model.DepFactory
	d = &DependencyFactoryImpl{}
	for i := 0; i < b.N; i++ {
		d.NewDependency()
	}
}
