package model

// FuncNewDataAccessLayer is factory function of DataAccessLayer
type FuncNewDataAccessLayer func() DataAccessLayer

type FunctionFactory struct {
	NewDataAccessLayer FuncNewDataAccessLayer
}
