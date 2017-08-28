package model

// AnotherDalFactory ...
type AnotherDalFactory interface {
	CreateAnotherDal() AnotherDal
}

// DataAccessLayerFactory ...
type DataAccessLayerFactory interface {
	CreateDal() DataAccessLayer
}

// BusinessServiceFactory ...
type BusinessServiceFactory interface {
	NewBusinessService(ff SuperFactory) BusinessService
}

// SuperFactory ...
type SuperFactory interface {
	DataAccessLayerFactory
	AnotherDalFactory
}
