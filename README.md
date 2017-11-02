# dutdip
Don't Use This Dependency Injection Pattern (D-U-T-D-I-P)

Read STORY.md to understand how dutdip came about.

## What constitutes a DI system

## Extra expectations from a DI system

# Releases

## v1.1.2 [Refer only what you need](https://github.com/rdadbhawala/dutdip/compare/v1.1.1...v1.1.2)
(continued from STORY.md)
Challenges
* The biggest irritation is the passing of SuperFactory all around in the code. Also, if any service is a Singleton, it may never get a refreshed factory.

Plans
* Let's define a mechanism for accessing it globally, to make code simpler and lighter.
