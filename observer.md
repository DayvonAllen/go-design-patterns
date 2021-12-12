## Observer Design Pattern Motivation
- `observer pattern` - An observer is an object that wishes to be informed about events happening in the system. The entity generating the events is an observable.
- Observer is an intrusive approach
- Must provide a way for clients to subscribe
- Event data sent from observable to all observers
- Data represented as `interface{}`
- Unsubscription is possible.
---