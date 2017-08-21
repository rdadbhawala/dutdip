package main

import "github.com/rdadbhawala/dutdip/dependency"

type superFactoryImpl struct {
	dependency.DataAccessLayerFactoryImpl
	dependency.AnotherDalFactoryImpl
}
