package main

import (
	"testing"

	"github.com/rdadbhawala/dutdip/samples/depParams/library"
	"github.com/rdadbhawala/dutdip/samples/depParams/service"
)

func BenchmarkService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.NewService(library.NewDependency())
	}
	// s.Feature()
}
