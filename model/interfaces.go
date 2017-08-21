package model

// BusinessService ...
type BusinessService interface {
	BusinessMethod1(callDal bool)
}

// DataAccessLayer is interface for DAL
type DataAccessLayer interface {
	DataMethod1()
}

// DataAccessLayerFactory ...
type DataAccessLayerFactory interface {
	CreateDal() DataAccessLayer
}

// SuperFactory ...
type SuperFactory interface {
	DataAccessLayerFactory
}

// // AnotherDal is another dependency
// type AnotherDal interface {
// 	DalMethod1()
// }
