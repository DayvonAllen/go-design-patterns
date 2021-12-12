## SOLID Design Principles
- `Single Responsibility Principle`(S) - A type should have one main responsible and as a result it should have one reason to change.
  - `Separation Of Concerns` - different problems of the system has to reside in different constructs.
- `Open-Closed Principal`(O) - Types should be open for extension but closed for modification.
- `Liskov Substitution Principle`(L) - If you have an API that takes a base class and works correctly with it, it should also work correctly with the derived class.
- `Interface Segregation Principal`(I) - You shouldn't put too much into an interface, it sometimes makes sense to break up interfaces into several smaller interfaces
- `Dependency Inversion Principle`(D) - It states that high level modules(HLM) shouldn't depend on low level modules(LLM)
  - It also states that both HLM and LLM should depend on abstraction
---

## Anti-Patterns
- `God Object` - Place everything in a single package
---