# Summary
The physical properties of an object.

# Description
The **PhysicalProperties** type represents the physical properties of an object.

# Constructors
## new
### Components
#### Summary
Returns properties with density, friction, and elasticity.

#### Description
The **new** constructor returns a new PhysicalProperties value. *density* sets
the Density field, *friction* sets the Friction field, and *elasticity* sets the
Elasticity field.

### Weights
#### Summary
#### Description
The **new** constructor returns a new PhysicalProperties value. *density* sets
the Density field, *friction* sets the Friction field, *elasticity* sets the
Elasticity field, *frictionWeight* sets the FrictionWeight field, and
*elasticityWeight* sets the ElasticityWeight field.

# Properties
## Density
### Summary
The mass per unit volume.

### Description
The **Density** field returns the mass for unit volume of the object.

## Elasticity
### Summary
How much energy is retained after a collision.

### Description
The **Elasticity** field returns how much energy is retained after a collision.
A value of 1 means the same energy is retained.

## ElasticityWeight
### Summary
The ratio baised towards the object for elasticity calculations.

### Description
The **ElasticityWeight** field returns how much the ratio of an elasticity
calculation is baised towards the object during a collision.

## Friction
### Summary
The force that opposes lateral motion.

### Description
The **Friction** field returns how much force is used to oppose lateral motion
between two contacting surfaces.

## FrictionWeight
### Summary
The ratio baised towards the object for friction calculations.

### Description
The **FrictionWeight** field returns how much the ratio of a friction
calculation is baised towards the object during a collision.

