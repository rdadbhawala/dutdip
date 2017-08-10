package model

// BusinessService ...
type BusinessService interface {
	BusinessMethod1(callDal bool)
}

// DataAccessLayer is interface for DAL
type DataAccessLayer interface {
	DataMethod1()
}

type AnotherDal interface {
	DalMethod1()
}
