## Builder Design Pattern Motivation
- Some objects are simple and can be created in a single constructor call
- Other objects may require a lot of ceremony to create
- Ex. Having a factory function with 10 arguments is not productive.
- Instead, we can opt for piecewise(piece by piece) construction.
- The Builder pattern provides an API for constructing an object step-by-step.
- We use the design pattern when piecewise object construction is complicated, we provide an API for doing it succinctly.
---