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
	NewBusinessService() BusinessService
}

// BusinessServiceDependencies ...
type BusinessServiceDependencies interface {
	DataAccessLayerFactory
	AnotherDalFactory
}

// SuperFactory ...
type SuperFactory interface {
	BusinessServiceFactory
	DataAccessLayerFactory
	AnotherDalFactory
}

var sfSingleton SuperFactory

// GetSuperFactory gets the global SuperFactory
func GetSuperFactory() SuperFactory {
	return sfSingleton
}

// SetSuperFactory sets the global SuperFactory
func SetSuperFactory(sf SuperFactory) {
	sfSingleton = sf
}
