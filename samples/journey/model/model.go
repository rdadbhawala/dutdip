package model

type Dependency interface {
	Operation()
}

type NewDependency func() Dependency

type Service interface {
	Function()
}
