# dutdip
Don't Use This Dependency Injection Pattern (D-U-T-D-I-P)

Read [STORY.md](STORY.md) to understand how dutdip came about.

## What constitutes a DI system
Quoting [Wikipedia](https://en.wikipedia.org/wiki/Dependency_injection): "In software engineering, dependency injection is a technique whereby one object supplies the dependencies of another object. A dependency is an object that can be used (a service). An injection is the passing of a dependency to a dependent object (a client) that would use it. The service is made part of the client's state. Passing the service to the client, rather than allowing a client to build or find the service, is the fundamental requirement of the pattern."

In other words, we are trying to move the responsibility of knowing and instantiating a dependency out of the service into some "container". This allows us to do beautiful things like replacing dependencies with mocks for Unit Testing. We are therefore, primarily, looking at:
* Dependency (which provides some service)
* Service (which wants to consume a dependency)
* Injector (who knows services and their dependencies)

Of course, a Service can be a Dependency, and a Dependency is a Service. Different implementations use different mechanisms for doing this voodoo. Some systems rely on configuration files, some on a precise piece of code at the beginning of the main, etc.

## Extra expectations from a DI system
Based on my experience with some of the projects, there's one additional constraint that I like to set on a DI system: **Initialize a Dependency only when required**. In most of the implementations, dependencies get initialized "outside" the service that wants to consume them; or they may get initialized before the service is constructed; or they get initialized, but based on the input, may not get utilized at all.
* **Hierarchy and Lifecycle of Objects**

  I am a big believer in the hierarchy of instances. I'm usually disturbed by overlapping lifecycles of objects and look at them as potentials "bugs" in the system. If instances are created and destroyed in a definite hierarchical scope, they are less likely to cause "issues". One of the challenges is that the dependencies are instantiated before they are injected. Especially in the constructor based mechanism, it therefore implies that dependencies get initialized before the service. There is greater control on the construction and destruction of the instances when a definite hierarchical scope is enforced.

* **Unutilized Dependencies**

  A Service may be implementing a complex workflow. Based on the input, some sections of the workflow are executed and some are skipped. This is entirely the business logic. This implies, that according to the workflow, some dependencies may not get utilized at all. Is it ok that such dependencies get initialized and injected? It can lead to utilization of some resources, such as a database connection. Of course, some "trivial" amount of CPU, Memory, Disc or Network may get utilized. As Engineers, it is our responsibility to optimize the resources that we consume, and so I'm inclined towards the idea that _dependencies should be initialized only when they are to be utilized_.

## Final Comments on DUTDIP
* Global Reference vs passing factories around?
  * I find the global reference practical. It reduces code, and doesn't require services to store the factories as a member.
  * Every Unit Test can setup the right mocks and assign to the global reference prior to running the test.
* Factory structure vs Package Factory Functions?
  * Factory functions are good for simpler scenarios. My projects involved several dependencies: database, http-client, message queue, invocations to other micro services, etc. and so I felt I needed something that is more centralized than package based factory function.
  * By using a Factory structure behind an interface, it is possible to try out multiple alternatives of the factory for the same service.
* It is about Loose Coupling
  * The core issue is not where an instance is created. The core issue is: who invokes this function, and what benefits can be claimed by creating the loose coupling.

# Releases

## v1.1.2 [Refer only what you need](https://github.com/rdadbhawala/dutdip/compare/v1.1.1...v1.1.2)
(continued from STORY.md)
Challenges
* The biggest irritation is the passing of SuperFactory all around in the code. Also, if any service is a Singleton, it may never get a refreshed factory.

Plans
* Let's define a mechanism for accessing it globally, to make code simpler and lighter.

## v1.2 [Don't Pass it Around](https://github.com/rdadbhawala/dutdip/compare/v1.1.2...v1.2)
Achievements
* Created a global reference for the Factories.
* Setup the factory at the beginning of main.

Challenges
* The code is now tightly coupled to a global reference to the Factories.

Plans
* None :)

