# dutdip
Don't Use This Dependency Injection Pattern (D-U-T-D-I-P)

In the quest for achieving design goals of maintainability, flexibility, extensibility and **_etcetrability_**, I boldly went where no one has gone before, and got kicked out! Nonetheless, I believe this pattern promotes a very high level of loose coupling, and you are free not to use it! 

Enter at your own peril. This pattern doesn't follow some of the recommended patterns of GoLang. You might feel offended. However, I don't write code with the assumption that well known patterns are the best solution for my needs. Nor do I assume that I can't do better.

## Conventions
* This repository will contain a series of commits that will take services and dependencies through varying mechanisms of orchestration in GoLang.
* Key stages in the evolution of this design will be tagged as a "release".
* Documentation will be added to this readme file explaining these releases.

## Terminology & Code Hints
* Service: an element in the 'business logic layer'.
* Dependency: an element consumed by the service.
* For our discussion, Service and Dependency reside in different packages. In reality, Dependencies may be in a different repository altogether. Package separation is adequate to represent that.
* At each release, we will evaluate Achievements, Issues and Plans.
* Run the code with the command to see simple Println output reflecting code behavior:
  `go run ./main/main.go`

# v0.1: [Compare View](https://github.com/rdadbhawala/dutdip/compare/v0.0...v0.1)
Achievements
* Created a [BusinessService](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/service/businessService.go) that depends on [DataAccessLayer](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/dependency/dal.go)
* Invoked from [Main](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/main/main.go)

Issues
* Tight Coupling: Main to BusinessService & DataAccessLayer, and BusinessService to DataAccessLayer

Plans
* Create interfaces in model package for loose coupling
* Create "New" functions to instantiate Service and Dependency in respective package.

# v0.2: [Compare View](https://github.com/rdadbhawala/dutdip/compare/v0.1...v0.2)
Achievements
* Created [Interfaces](https://github.com/rdadbhawala/dutdip/blob/474398b5f8c21c01d85cf970ed44bbbccbc2dbb6/model/interfaces.go) for BusinessService and DataAccessLayer
* Made implementations of interface private to the respective package: [Service](https://github.com/rdadbhawala/dutdip/blob/474398b5f8c21c01d85cf970ed44bbbccbc2dbb6/service/businessService.go#L17) and [Dependency](https://github.com/rdadbhawala/dutdip/blob/474398b5f8c21c01d85cf970ed44bbbccbc2dbb6/dependency/dal.go#L13)
* Added a constructor 'New' function to instantiate Service and Dependency: [Service](https://github.com/rdadbhawala/dutdip/blob/474398b5f8c21c01d85cf970ed44bbbccbc2dbb6/service/businessService.go#L11) and [Dependency](https://github.com/rdadbhawala/dutdip/blob/474398b5f8c21c01d85cf970ed44bbbccbc2dbb6/dependency/dal.go#L9)

Issues
* The Service package is still tightly coupled with Dependency. Though service.businessServiceImpl refers to the dependency by its interface, the "NewBusinessService" function still needs to know the actual dependency.
* The import statement reflects the tight-coupling of service package with the dependency package.
* As long as the Service package refers to the Dependency package, 'clean' loose coupling can't be achieved.

Plans
* Let us try to overcome this limitation by passing the dependency itself as a parameter to the New function

# v0.3: [Compare View](https://github.com/rdadbhawala/dutdip/compare/v0.2...v0.3)
Achievements
* Service package's tight coupling with Dependency package removed.
* Main method initializes the dependency and passes it as a parameter.
* Enables mocking of DataAccessLayer in the BusinessService.

Issues
* Every instantiation of BusinessService will require the Dependency to be passed as a parameter.
* If a new dependency is added, every invocation of NewBusinessService will also have to be updated. This implies breaking changes in code every time the dependencies change.
* Dependency is initialized even if it is not utilized. This challenge exists even in case of v0.2.
  * Underlying resources of the dependency get allocated (eg: eatabase connection is opened) even if it is not utilized. It is our responsibility to use the resources in an optimum manner. Such overheads should be avoided. Resources like database connections are expensive and shouldn't be misused.
  * This can be overcome with lazy initialization check on every method of the dependency. This, however, is challenging to implement and easy to forget for new methods. This pattern will have higher maintainance effort.

Plans
* Indicate the challenge of unutilized dependencies.

# v0.4: Objectives of Dependency Injection
This is probably a good time to take a step back and define the objectives for Dependency Injection.
* Dependencies be injected
  
  This means that when a particular element (or a component) makes calls to a different subsystem to achieve its results, then such a dependency should be injected into the element, and not instantiated/ created by the element itself. This obviously is the most basic and principle expectation from a Dependency Injection framework.
* Prevent breaking changes

  Should addition of more dependencies or new relationships in an element/ component lead to breaking channges in the code already written? I think not. While I acknowledge that requirements change, I believe in the principle that good software design, by definition, shouldn't require you to look backwards. Rather, it allows you to look forward freely, unconstrained by the work already done. In that sense, if there are changes in the dependencies of an element, the DI framework should absorb the impact in such a manner that such changes can be incorporated easily.
* Dependencies are initialized only if required

  I am a big believer in the pattern that the creator of an element should also be responsible for its destruction. The life cycle of such a dependency should be managed in a clean manner, clearly depicting where the dependency is created and where it is relieved. Several DI frameworks use a constructor based pattern for injecting dependencies. In such a pattern, the dependency is created in a different scope (DI Framework), and is utilized in a different scope (the service). There is ambiguity over who is responsible for the clean up. If dependencies become members of the Service, it makes the service stateful. If the dependency has a database connection, to take an example, this instance of the service can't be shared across requests.
  
  Also, as a responsible software engineer, I want to utilize resources only if I have to use them. Opening a connection has its costs: the cost to initialize connection objects, establishing communication and handshake with the database, implicit 'transaction', and of course there are processing and memory costs to be considered.

I hope the above paragraphs outline a clear set of expectations. We will now go through a set of convoluted code changes to finally arrive at the DUTDIP, the pattern that I've been told we shouldn't be using.

# v0.5: [Unexpected Dependency Initialization](https://github.com/rdadbhawala/dutdip/compare/v0.4...v0.5)
Achievements
* BusinessService interface was updated to take a bool parameter, which dictates invocation of DAL.
* Some print statements were added to reflect behavior of the code.

Issues
* According to the order of output statements, the dependency is getting initialized before the Service itself. That just doesn't sound right at all.
* Also, the second invocation of the BusinessMethod1 shows that even though the dependency was not invoked, it was certainly initialized.

Plans
* Let us now understand the challenge of a new dependency.

# v0.6: [Changes in Dependency Leads to Breaking Changes](https://github.com/rdadbhawala/dutdip/compare/v0.5...v0.6)
Achievements
* A new dependency is introduced in the same manner as the existing one.
* It is passed as a parameter to the NewBusinessService method.

Issues
* With the signature of the NewBusinessService changing, it leads to breaking changes at every location that it is consumed.
* 'main' function in the sample code had to be fixed to align with the new signature.
* As above, dependencies are initialized before the service, even if they are not utilized.

Plans
* One mechanism to overcome the 2 challenges explained in v0.5 and v0.6 is pass the "New" function itself as a parameter.
* In this sense, the Service has access to the Dependency factory (instead of an instance of a Dependency).
* We will revert back to a single dependency to showcase this solution.

# v0.7 [Factory Not Instance](https://github.com/rdadbhawala/dutdip/compare/v0.5...v0.6)
Achievements
* Passed the "New" function as a parameter.
* Service invokes the function to create an instance of the dependency only when it wants to consume the dependency.
* Dependency is now initialized after the Service, and only if required.

Issues
* We still face the challenge of changes in dependencies. If another dependency is added, it leads to breaking changes in code. (I will not create an example for this issue, as it is the same as described in v0.6).
* There is nothing wrong in passing functions. Functions are first class members of the codebase. However, functions don't lend themselves to easy reuse. Code constructs such as Inheritance and Composition do.

Plans
* To prevent breaking changes in the method signature, we must gather the parameters into a single element. Let us check if we can put factory functions in a structure.
