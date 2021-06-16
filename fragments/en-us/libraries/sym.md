# Summary
Symbols for accessing instance metadata.

# Description
The **sym** library contains **Symbol** values. A symbol is a unique identifier
that can be used to access certain metadata fields of an [Instance][Instance].

An instance can be indexed with a symbol to get a metadata value in the same way
it can be indexed with a string to get a property value:

```lua
local instance = Instance.new("Workspace")
instance[sym.IsService] = true
print(instance[sym.IsService]) --> true
```

The following symbols are defined:

$FIELDS

# Fields
## AttrConfig
### Summary
Gets the inherited [AttrConfig][AttrConfig] of an instance.

### Description
The **AttrConfig** symbol gets the inherited [AttrConfig][AttrConfig] of an
instance.

## Desc
### Summary
Gets the inherited [descriptor][RootDesc] of an instance.

### Description
The **Desc** symbol gets the inherited [descriptor][RootDesc] of an instance.

## IsService
### Summary
Determines whether an instance is a service.

### Description
The **IsService** symbol determines whether an instance is a service.

## Metadata
### Summary
Gets the metadata of a [DataModel][DataModel].

### Description
The **Metadata** symbol gets the metadata of a [DataModel][DataModel].

## Properties
### Summary
Gets the properties of an instance.

### Description
The **Properties** symbol gets the properties of an instance.

## RawAttrConfig
### Summary
Accesses the direct [AttrConfig][AttrConfig] of an instance.

### Description
The **RawAttrConfig** symbol accesses the direct [AttrConfig][AttrConfig] of an
instance.

## RawDesc
### Summary
Accesses the direct [descriptor][RootDesc] of an instance.

### Description
The **RawDesc** symbol accesses the direct [descriptor][RootDesc] of an
instance.

## Reference
### Summary
Determines the value used to identify the instance.

### Description
The **Reference** symbol determines the value used to identify the instance.
