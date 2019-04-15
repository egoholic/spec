# Spec

## Goal
There are two goals we want to achieve with Spec library:

- We want to write less code and do more business with people.
- We want to be sure that our system's pieces fit well together.

The method to fullfill both needs is to generate typical, well-known pieces of code and write only too specific code so-called "business logic".

## Purpose
Spec allows:

- specify endpoint, its input and output,
- verify endpoints according to specification through tests code generation,
- generate endpoint's form objects,
- generate endpoint's data serializers
- generate endpoint's data deserializers
- define custom types by applying predicates on standard types



## Ubiquitous Language

**Endpoint** - is a basic deliverable unit, quant of our busines software architecture which could deal with one or many interactions.

**Interaction** - is the way how one endpoint could touch another.

**Raw Signature** - text representing an operationable (parsed) signature.

**Signature** - parsead, operationable signature that describes a data type. Signatures are composable and most operations on them support mathematical quality of closedness (products of operation on signtures are signatures too).

**Domain**  - a type scoped with one or many predicates. For example `int` is a type while `int{isOdd}` is a domain. Domains are equal to types. We use term domain here as a detail of implementation. By default spec supports only Golang standard data types and all the custom types could be described in spec as domains.

**Alias** - alias is a short name for signature. Signatures and their aliaces are interchangeable. Aliases are signatures too.

**Primitive Signature** -

**Endpoint Signature** -

**Endpoint Composition Signature** -

**Form Object** -

**Endpoint Serializer** -

**Endpoint Deserializer** -



## How To's

### How to describe endpoint?

### How to generate and use form object?

### How to generate and use endpoint serializer?

### How to generate and use endpoint deserializer?

### How to create and use domain?

### How to create and use alias?

### How to work with subtypes?

### How to play with type algebra?



## Participation



## What is Stoa?