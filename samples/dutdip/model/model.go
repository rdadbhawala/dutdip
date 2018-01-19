package model

// Dependency interface
type Dependency interface {
	Operation()
}

// DepFactory ...
type DepFactory interface {
	NewDependency() Dependency
}

// Service interface is a service
type Service interface {
	Feature()
}

// ServiceFactory ...
type ServiceFactory interface {
	NewService() Service
}

// AllFactories ...
type AllFactories interface {
	ServiceFactory
	DepFactory
}

// F ...
var F AllFactories

var (
	// S ...
	S ServiceFactory
	// D ...
	D DepFactory
)
