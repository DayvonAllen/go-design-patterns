## Serialization
- Good for deep copying
- `serialization` constructs are smart so, instead of serializing pointers, it will go to the object at the memory address and serialize that
- `serializers` know how to unwrap a structure and serialize all of its members.
- We can use the `gob` package to serialize structs into binary
- When you `deserialize` an object, you construct a new object with all the same values.
---