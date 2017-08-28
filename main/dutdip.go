package main

import "github.com/rdadbhawala/dutdip/dependency"
import "github.com/rdadbhawala/dutdip/service"

type superFactoryImpl struct {
	dependency.DataAccessLayerFactoryImpl
	dependency.AnotherDalFactoryImpl
	service.BusinessServiceFactoryImpl
}
