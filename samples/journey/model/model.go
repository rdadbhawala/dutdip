package model

type Dependency interface {
	Operation()
}

type DependencyFactory interface {
	NewDependency() Dependency
}

type Service interface {
	Function()
}

type AnotherDep interface {
	AnotherOp()
}

type AnotherDepFactory interface {
	NewAnotherDep() AnotherDep
}

type SuperFactory interface {
	DependencyFactory
	AnotherDepFactory
}
