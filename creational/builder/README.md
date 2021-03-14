# Builder

Allow separating the construction of an object (complex or not) from its representation
so that the same creation process can create different representations.

- **Product**: main object built. Represents the object under construction.
- **Builder**: interface with the methods that the constructors should implement.
- **Concrete builder**: implements the interface builder to deliver the concrete object.
- **Director**: build the object using the interface
