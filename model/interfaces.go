package model

// BusinessService ...
type BusinessService interface {
	BusinessMethod1(callDal bool)
}

// DataAccessLayer is interface for DAL
type DataAccessLayer interface {
	DataMethod1()
}

// FuncNewDataAccessLayer is factory function of DataAccessLayer
type FuncNewDataAccessLayer func() DataAccessLayer

// AnotherDal is another dependency
type AnotherDal interface {
	DalMethod1()
}
