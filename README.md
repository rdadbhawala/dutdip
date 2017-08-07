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

# v0.1
Achievements
* Created a [BusinessService](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/service/businessService.go) that depends on [DataAccessLayer](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/dependency/dal.go)
* Invoked from [Main](https://github.com/rdadbhawala/dutdip/blob/9362e5f67c33288da543acd9f96469b7aef5db62/main/main.go)

Issues
* Tight Coupling: Main to BusinessService & DataAccessLayer, and BusinessService to DataAccessLayer

Plans
* Create interfaces in model package for loose coupling
* Create "New" functions to instantiate Service and Dependency in respective package.

# v0.2
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
