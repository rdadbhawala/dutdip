package model

// AnotherDalFactory ...
type AnotherDalFactory interface {
	CreateAnotherDal() AnotherDal
}

// DataAccessLayerFactory ...
type DataAccessLayerFactory interface {
	CreateDal() DataAccessLayer
}

// SuperFactory ...
type SuperFactory interface {
	DataAccessLayerFactory
	AnotherDalFactory
}
