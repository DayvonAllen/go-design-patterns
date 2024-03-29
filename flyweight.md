## Flyweight Design Pattern Motivation
- Space Optimization
- Avoid redundancy when storing data
- Store common data externally
- Specify an index or a pointer into the external data store
- `Flyweight Pattern` - A space optimization technique that lets us use less memory by storing externally the data associated with similar objects.
- `memento` vs `flyweight`:
    - Both patterns provide a 'token' clients can hold on to
    - Memento is used only to be fed back into the system:
        - N public/mutable state
        - no methods
    - A flyweight is similar to an ordinary reference to an object
        - Can mutate state
        - can provide additional functionality(fields/methods)
---