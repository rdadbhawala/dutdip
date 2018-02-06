package model

type Dependency interface {
	Operation()
	Release()
}

type DependencyFactory interface {
	NewDependency() Dependency
}

type Service interface {
	Function()
}

type ServiceFactory interface {
	NewService() Service
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
	ServiceFactory
}
