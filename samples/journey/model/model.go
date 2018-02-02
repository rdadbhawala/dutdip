package model

type Dependency interface {
	Operation()
}

type NewDependency func() Dependency

type Service interface {
	Function()
}

type AnotherDep interface {
	AnotherOp()
}

type NewAnotherDep func() AnotherDep

type FuncFact struct {
	nd NewDependency
	na NewAnotherDep
}

func (f *FuncFact) NewDependency() Dependency {
	return f.nd()
}

func (f *FuncFact) NewAnotherDep() AnotherDep {
	return f.na()
}
