## Factories Design Pattern Motivation
- Object creation logic becomes too convoluted
- Struct has too many fields, need to initialize all correctly
- Wholesale object creation(non-piecewise, unlike the builder pattern) can be outsourced to:
  - A separate function `Factory Function, a.k.a Constructor`
  - That may exist in a separate struct(Factory)
- Create object in one invocation of a factory function.
- `Factory` - a component responsible solely for the wholesale(not piecewise) creation of objects.
- Ways of controlling how an object is constructed.
---