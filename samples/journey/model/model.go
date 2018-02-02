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

type Functions struct {
	nd NewDependency
	na NewAnotherDep
}

func (f *Functions) NewDependency() Dependency {
	return f.nd()
}

func (f *Functions) NewAnotherDep() AnotherDep {
	return f.na()
}

func NewFunctions(nd NewDependency, na NewAnotherDep) *Functions {
	return &Functions{nd, na}
}
