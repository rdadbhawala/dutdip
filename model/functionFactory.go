package model

// FuncNewDataAccessLayer is factory function of DataAccessLayer
type FuncNewDataAccessLayer func() DataAccessLayer

// FuncNewAnotherDal is Another Dependency
type FuncNewAnotherDal func() AnotherDal

// FunctionFactory is Factory of Functions
type FunctionFactory struct {
	NewDataAccessLayer FuncNewDataAccessLayer
	NewAnotherDal      FuncNewAnotherDal
}
