# Types
This document contains a reference to the types available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [AttrConfig][AttrConfig]
2. [Axes][Axes]
3. [CallbackDesc][CallbackDesc]
4. [ClassDesc][ClassDesc]
5. [Cookie][Cookie]
6. [Cookies][Cookies]
7. [DataModel][DataModel]
8. [DescAction][DescAction]
9. [DescActions][DescActions]
10. [EnumDesc][EnumDesc]
11. [EnumItemDesc][EnumItemDesc]
12. [EventDesc][EventDesc]
13. [Faces][Faces]
14. [FormatSelector][FormatSelector]
15. [FunctionDesc][FunctionDesc]
16. [HTTPHeaders][HTTPHeaders]
17. [HTTPOptions][HTTPOptions]
18. [HTTPRequest][HTTPRequest]
19. [HTTPResponse][HTTPResponse]
20. [Instance][Instance]
21. [Intlike][Intlike]
22. [MemberDesc][MemberDesc]
23. [Numberlike][Numberlike]
24. [ParameterDesc][ParameterDesc]
25. [PropertyDesc][PropertyDesc]
26. [RBXAssetOptions][RBXAssetOptions]
27. [RootDesc][RootDesc]
28. [Stringlike][Stringlike]
29. [TypeDesc][TypeDesc]
30. [UniqueId][UniqueId]

</td></tr></tbody>
</table>

This document only describes the types implemented by rbxmk. The
[Libraries](libraries.md#user-content-roblox) document lists the Roblox types
available in the rbxmk environment. See the [Data type
index](https://developer.roblox.com/en-us/api-reference/data-types) page on the
DevHub for information about Roblox types.

## AttrConfig
[AttrConfig]: #user-content-attrconfig

The **AttrConfig** type configures how an instance encodes and decodes
attributes.

Member                                     | Kind
-------------------------------------------|-----
[AttrConfig.Property][AttrConfig.Property] | property

An AttrConfig can be created with the [AttrConfig.new][AttrConfig.new]
constructor.

### AttrConfig.new
[AttrConfig.new]: #user-content-attrconfignew
<code>AttrConfig.new(property: [string](##)?): [AttrConfig][AttrConfig]</code>

The **new** constructor creates a new AttrConfig. *property* sets the
[Property][AttrConfig.Property] field, defaulting to an empty string.

### AttrConfig.Property
[AttrConfig.Property]: #user-content-attrconfigproperty
<code>AttrConfig.Property: [string](##)</code>

The **Property** property determines which property of an [Instance][Instance]
attributes are applied to. If an empty string, instances will default to
"AttributesSerialize".

### Axes
[Axes]: #user-content-axes

The **Axes** type corresponds to the [Axes][Axes-roblox] Roblox type.

[Axes-roblox]: https://developer.roblox.com/en-us/api-reference/datatype/Axes

#### Axes.new
[Axes.new]: #user-content-axesnew
<code>Axes.new(...any?)</code>

The **Axes.new** constructor returns a new Axes value. Each valid argument sets a
component of the value, depending on the type:

- EnumItem:
	- Enum name is `"Axis"`, and item name is `"X"`, `"Y"`, or `"Z"`.
	- Enum name is `"NormalId"`, and item name is one of the following:
		- `"Right"`, `"Left"`: sets X.
		- `"Top"`, `"Bottom"`: sets Y.
		- `"Back"`, `"Front"`: sets Z.
- number: value is one of the following:
	- `0`: sets X.
	- `1`: sets Y.
	- `2`: sets Z.
- string: value is `"X"`, `"Y"`, or `"Z"`.

Other values will be ignored.

#### Axes.fromComponents
[Axes.fromComponents]: #user-content-axesfromcomponents
<code>Axes.fromComponents(x: bool, y: bool, z: bool)</code>

The **Axes.fromComponents** constructor returns a new Axes value, with each
argument setting the corresponding component.

## CallbackDesc
[CallbackDesc]: #user-content-callbackdesc
<code>type CallbackDesc = {MemberType: [string](##), Name: [string](##), Parameters: {[ParameterDesc][ParameterDesc]}, ReturnType: [TypeDesc][TypeDesc], Security: [string](##), Tags: {[string](##)}}</code>

The **CallbackDesc** type is a table that describes a callback member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Callback".
Name       | [string](##)                     | The name of the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the callback.
ReturnType | [TypeDesc][TypeDesc]             | The type of the value returned by the callback.
Security   | [string](##)                     | The security context required to set the member.
Tags       | {[string](##)}                   | The tags set for the member.

## ClassDesc
[ClassDesc]: #user-content-classdesc
<code>type ClassDesc = {MemoryCategory: [string](##), Name: [string](##), Superclass: [string](##), Tags: {[string](##)}}</code>

The **ClassDesc** type is a table that describes a class. It has the following fields:

Field          | Type           | Description
---------------|----------------|------------
MemoryCategory | [string](##)   | The category of the class.
Name           | [string](##)   | The name of the class.
Superclass     | [string](##)   | The name of the class from which the current class inherits.
Tags           | {[string](##)} | The tags set for the class.

## Cookie
[Cookie]: #user-content-cookie

The **Cookie** type contains information about an HTTP cookie. It has the
following members:

Member                     | Kind
---------------------------|-----
[Cookie.Name][Cookie.Name] | field

For security reasons, the value of the cookie cannot be accessed.

Cookie is immutable. A Cookie can be created with the [Cookie.new][Cookie.new]
constructor. Additionally, Cookies can be fetched from known locations with the
[Cookie.from][Cookie.from] function.

### Cookie.from
[Cookie.from]: #user-content-cookiefrom
<code>Cookie.from(location: string): (cookies: [Cookies][Cookies]?)</code>

The **from** constructor retrieves cookies from a known location. *location* is
case-insensitive.

The following locations are implemented:

Location | Description
---------|------------
`studio` | Returns the cookies used for authentication when logging into Roblox Studio.

Returns nil if no cookies could be retrieved from the location. Throws an error
if an unknown location is given.

### Cookie.new
[Cookie.new]: #user-content-cookienew
<code>Cookie.new(name: [string](##), value: [string](##)): [Cookie][Cookie]</code>

The **new** constructor creates a new cookie object.

### Cookie.Name
[Cookie.Name]: #user-content-cookiename
<code>Cookie.Name: [string](##)</code>

The **Name** field is the name of the cookie.

## Cookies
[Cookies]: #user-content-cookies

The **Cookies** type is a list of [Cookie][Cookie] values.

## DataModel
[DataModel]: #user-content-datamodel

The **DataModel** type is a special case of an [Instance][Instance]. In addition
to the members of Instance, DataModel has the following members:

Member                                              | Kind
----------------------------------------------------|-----
[DataModel.GetService][DataModel.GetService]        | method
[DataModel\[sym.Metadata\]][DataModel.sym.Metadata] | symbol

Unlike a normal Instance, the [ClassName][Instance.ClassName] property of a
DataModel cannot be modified. Properties on a DataModel are usually not
serialized.

A DataModel can be created with the [Instance.new][Instance.new] constructor
with "DataModel" as the *className*.

### DataModel.GetService
[DataModel.GetService]: #user-content-datamodelgetservice
<code>DataModel:GetService(className: [string](##)): [Instance][Instance]</code>

The **GetService** method returns the first child of the DataModel whose
[ClassName][Instance.ClassName] equals *className*. If no such child exists,
then a new instance of *className* is created. The [Name][Instance.Name] of the
instance is set to *className*, [sym.IsService][Instance.sym.IsService] is set
to true, and [Parent][Instance.Parent] is set to the DataModel.

If the DataModel has a descriptor, then GetService will throw an error if the
created class's descriptor does not have the "Service" tag set.

### DataModel[sym.Metadata]
[DataModel.sym.Metadata]: #user-content-datamodelsymmetadata
<code>DataModel\[sym.Metadata\]: {[[string](##)]: [string](##)}</code>

The **Metadata** symbol gets or sets the metadata associated with the DataModel.
This metadata is used by certain formats (e.g. ExplicitAutoJoints).

## DescAction
[DescAction]: #user-content-descaction

The **DescAction** type describes a single action that transforms a descriptor.
It has the following members:

Member                                       | Kind
---------------------------------------------|-----
[DescAction.Type][DescAction.Type]           | property
[DescAction.Element][DescAction.Element]     | property
[DescAction.Primary][DescAction.Primary]     | property
[DescAction.Secondary][DescAction.Secondary] | property
[DescAction.Field][DescAction.Field]         | method
[DescAction.Fields][DescAction.Fields]       | method
[DescAction.SetField][DescAction.SetField]   | method
[DescAction.SetFields][DescAction.SetFields] | method

A DescAction can be created with the [DescAction.new][DescAction.new]
constructor. It is also returned from the [RootDesc.Diff][RootDesc.Diff] method
and the [desc-patch.json](formats.md#user-content-desc-patchjson) format.

### DescAction.new
[DescAction.new]: #user-content-descactionnew
<code>DescAction.new(type: [Enum.DescActionType][Enum.DescActionType], element: [Enum.DescActionElement][Enum.DescActionElement]): [DescAction][DescAction]</code>

The **DescAction.new** constructor returns a new DescAction initialized with a
type and an element.

### DescAction.Type
[DescAction.Type]: #user-content-descactiontype
<code>DescAction.Type: [Enum.DescActionType][Enum.DescActionType]</code>

The **Type** property is the type of transformation performed by the action.

[Enum.DescActionType]: enums.md#user-content-descactiontype

### DescAction.Element
[DescAction.Element]: #user-content-descactionelement
<code>DescAction.Element: [Enum.DescActionElement][Enum.DescActionElement]</code>

The **Element** property is the type of element to which the action applies.

[Enum.DescActionElement]: enums.md#user-content-descactionelement

### DescAction.Primary
[DescAction.Primary]: #user-content-descactionprimary
<code>DescAction.Primary: [string](##)</code>

The **Primary** property is the name of the primary element. For example, the
class name or enum name.

### DescAction.Secondary
[DescAction.Secondary]: #user-content-descactionsecondary
<code>DescAction.Secondary: [string](##)</code>

The **Secondary** property is the name of the secondary element. For example,
the name of a class member or enum item. An empty string indicates that the
action applies to the primary element.

### DescAction.Field
[DescAction.Field]: #user-content-descactionfield
<code>DescAction:Field(name: [string](##)): [any](##)?</code>

The **Field** method returns the value of the *name* field, or nil if the action
has no such field.

### DescAction.Fields
[DescAction.Fields]: #user-content-descactionfields
<code>DescAction:Fields(): {\[[string](##)\]: [any](##)}</code>

The **Fields** method returns a collection of field names mapped to values.

### DescAction.SetField
[DescAction.SetField]: #user-content-descactionsetfield
<code>DescAction:SetField(name: [string](##), value: [any](##))</code>

The **SetField** method sets the value of the *name* field to *value*.

### DescAction.SetFields
[DescAction.SetFields]: #user-content-descactionsetfields
<code>DescAction:SetFields(fields: {\[[string](##)\]: [any](##)})</code>

The **SetFields** method replaces all fields of the action with *fields*.

## DescActions
[DescActions]: #user-content-descactions

The **DescActions** type is a list of [DescAction][DescAction] values.

## EnumDesc
[EnumDesc]: #user-content-enumdesc
<code>type EnumDesc = {Name: [string](##), Tags: {[string](##)}}</code>

The **EnumDesc** type is a table that describes an enum. It has the following
fields:

Field | Type           | Description
------|----------------|------------
Name  | [string](##)   | The name of the enum.
Tags  | {[string](##)} | The tags set for the enum.

## EnumItemDesc
[EnumItemDesc]: #user-content-enumitemdesc
<code>type EnumItemDesc = {Index: [int](##), Name: [string](##), Tags: {[string](##)}, Value: [int](##)}</code>

The **EnumitemDesc** type is a table that describes an item of an enum. It has
the following fields:

Field      | Type           | Description
-----------|----------------|------------
Index      | [int](##)      | Hints the order of the item..
Name       | [string](##)   | The name of the item.
Tags       | {[string](##)} | The tags set for the item.
Value      | [int](##)      | The numeric value of the item.

## EventDesc
[EventDesc]: #user-content-eventdesc
<code>type EventDesc = {MemberType: [string](##), Name: [string](##), Parameters: {[ParameterDesc][ParameterDesc]}, Security: [string](##), Tags: {[string](##)}}</code>

The **EventDesc** type is a table that describes a event member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Event".
Name       | [string](##)                     | The name of the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the event.
Security   | [string](##)                     | The security context required to get the member.
Tags       | {[string](##)}                   | The tags set for the member.

### Faces
[Faces]: #user-content-faces

The **Faces** type corresponds to the [Faces][Faces-roblox] Roblox type.

[Faces-roblox]: https://developer.roblox.com/en-us/api-reference/datatype/Faces

#### Faces.new
[Faces.new]: #user-content-axesnew
<code>Faces.new(...any?)</code>

The **Faces.new** constructor returns a new Faces value. Each valid argument
sets a component of the value, depending on the type:

- EnumItem:
	- Enum name is `"Axis"`, and item name is one of the following:
		- `"X"`: sets Right and Left.
		- `"Y"`: sets Top and Bottom.
		- `"Z"`: sets Back and Front.
	- Enum name is `"NormalId"`, and item name is `"Right"`, `"Top"`, `"Back"`,
	  `"Left"`, `"Bottom"`, or `"Front"`.
- number: value is one of the following:
	- `0`: sets Right.
	- `1`: sets Top.
	- `2`: sets Back.
	- `3`: sets Left.
	- `4`: sets Bottom.
	- `5`: sets Front.
- string: value is `"Right"`, `"Top"`, `"Back"`, `"Left"`, `"Bottom"`, or
  `"Front"`.

Other values will be ignored.

#### Faces.fromComponents
[Faces.fromComponents]: #user-content-axesfromcomponents
<code>Faces.fromComponents(right: bool, top: bool, back: bool, left: bool, bottom: bool, front: bool)</code>

The **Faces.fromComponents** constructor returns a new Faces value, with each
argument setting the corresponding component.

## FormatSelector
[FormatSelector]: #user-content-formatselector
<code>type FormatSelector = string \| {Format: string, ...}</code>

The **FormatSelector** type selects a [format](formats.md), and optionally
configures the format.

If a table, then the Format field indicates the name of the format to use, and
remaining fields are options that configure the format, which depend on the
format specified. All such fields are optional.

If a string, it is the name of the format to use, and specifies no options.

## FunctionDesc
[FunctionDesc]: #user-content-functiondesc

The **FunctionDesc** type is a table that describes a function member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Function".
Name       | [string](##)                     | The name of the member.
ReturnType | [TypeDesc][TypeDesc]             | The type of the value returned by the function.
Security   | [string](##)                     | The security context required to set the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the function.
Tags       | {[string](##)}                   | The tags set for the member.

## HTTPHeaders
[HTTPHeaders]: #user-content-httpheaders
<code>type HTTPHeaders = {\[[string](##)\]: [string](##)\|{[string](##)}}</code>

The **HTTPHeaders** type is a table that specifies the headers of an HTTP
request or response. Each entry consists of a header name mapped to a string
value. If a header requires multiple values, the name may be mapped to an array
of values instead.

For response headers, a header is always mapped to an array, and each array will
have at least one value.

## HTTPOptions
[HTTPOptions]: #user-content-httpoptions
<code>type HTTPOptions = {URL: [string](##), Method: [string](##)?, RequestFormat: [FormatSelector][FormatSelector], ResponseFormat: [FormatSelector][FormatSelector], Headers: [HTTPHeaders][HTTPHeaders]?, Cookies: [Cookies][Cookies]?, Body: [any](##)?}</code>

The **HTTPOptions** type is a table that specifies how an HTTP request is made.
It has the following fields:

Field          | Type                              | Description
---------------|-----------------------------------|------------
URL            | [string](##)                      | The URL to make to request to.
Method         | [string](##)?                     | The HTTP method. Defaults to GET.
RequestFormat  | [FormatSelector][FormatSelector]? | The format used to encode the request body.
ResponseFormat | [FormatSelector][FormatSelector]? | The format used to decode the response body.
Headers        | [HTTPHeaders][HTTPHeaders]?       | The HTTP headers to include with the request.
Cookies        | [Cookies][Cookies]?               | Cookies to append to the Cookie header.
Body           | [any](##)?                        | The body of the request, to be encoded by the specified format.

If RequestFormat is unspecified, then no request body is sent. If ResponseFormat
is unspecified, then no response body is returned.

Use of the Cookies field ensures that cookies sent with the request are
well-formed, and is preferred over setting the Cookie header directly.

## HTTPRequest
[HTTPRequest]: #user-content-httprequest
<code>type HTTPRequest</code>

The **HTTPRequest** type represents a pending HTTP request. It has the following
members:

Member                                     | Kind
-------------------------------------------|-----
[HTTPRequest.Resolve][HTTPRequest.Resolve] | method
[HTTPRequest.Cancel][HTTPRequest.Cancel]   | method

### HTTPRequest.Resolve
[HTTPRequest.Resolve]: #user-content-httprequestresolve
<code>HTTPRequest:Resolve(): (resp: [HTTPResponse][HTTPResponse])</code>

The **Resolve** method blocks until the request has finished, and returns the
response. Throws an error if a problem occurred while resolving the request.

### HTTPRequest.Cancel
[HTTPRequest.Cancel]: #user-content-httprequestcancel
<code>HTTPRequest:Cancel()</code>

The **Cancel** method cancels the pending request.

## HTTPResponse
[HTTPResponse]: #user-content-httpresponse
<code>type HTTPResponse = {Success: [bool](##), StatusCode: [int](##), StatusMessage: [string](##), Headers: [HTTPHeaders][HTTPHeaders], Cookies: [Cookies][Cookies], Body: [any](##)?}</code>

The **HTTPResponse** type is a table that contains the response of a request. It
has the following fields:

Field         | Type                       | Description
--------------|----------------------------|------------
Success       | [bool](##)                 | Whether the request succeeded. True if StatusCode between 200 and 299.
StatusCode    | [int](##)                  | The HTTP status code of the response.
StatusMessage | [string](##)               | A readable message corresponding to the StatusCode.
Headers       | [HTTPHeaders][HTTPHeaders] | A set of response headers.
Cookies       | [Cookies][Cookies]         | Cookies parsed from the Set-Cookie header.
Body          | [any](##)?                 | The decoded body of the response.

## Instance
[Instance]: #user-content-instance

The **Instance** type provides a similar API to that of Roblox. In addition to
getting and setting properties, instances have the following members defined:

Member                                                                   | Kind
-------------------------------------------------------------------------|-----
[Instance.ClassName][Instance.ClassName]                                 | property
[Instance.Name][Instance.Name]                                           | property
[Instance.Parent][Instance.Parent]                                       | property
[Instance.ClearAllChildren][Instance.ClearAllChildren]                   | method
[Instance.Clone][Instance.Clone]                                         | method
[Instance.Descend][Instance.Descend]                                     | method
[Instance.Destroy][Instance.Destroy]                                     | method
[Instance.FindFirstAncestor][Instance.FindFirstAncestor]                 | method
[Instance.FindFirstAncestorOfClass][Instance.FindFirstAncestorOfClass]   | method
[Instance.FindFirstAncestorWhichIsA][Instance.FindFirstAncestorWhichIsA] | method
[Instance.FindFirstChild][Instance.FindFirstChild]                       | method
[Instance.FindFirstChildOfClass][Instance.FindFirstChildOfClass]         | method
[Instance.FindFirstChildWhichIsA][Instance.FindFirstChildWhichIsA]       | method
[Instance.GetAttribute][Instance.GetAttribute]                           | method
[Instance.GetAttributes][Instance.GetAttributes]                         | method
[Instance.GetChildren][Instance.GetChildren]                             | method
[Instance.GetDescendants][Instance.GetDescendants]                       | method
[Instance.GetFullName][Instance.GetFullName]                             | method
[Instance.IsA][Instance.IsA]                                             | method
[Instance.IsAncestorOf][Instance.IsAncestorOf]                           | method
[Instance.IsDescendantOf][Instance.IsDescendantOf]                       | method
[Instance.SetAttribute][Instance.SetAttribute]                           | method
[Instance.SetAttributes][Instance.SetAttributes]                         | method
[Instance\[sym.AttrConfig\]][Instance.sym.AttrConfig]                    | symbol
[Instance\[sym.Desc\]][Instance.sym.Desc]                                | symbol
[Instance\[sym.IsService\]][Instance.sym.IsService]                      | symbol
[Instance\[sym.Properties\]][Instance.sym.Properties]                    | symbol
[Instance\[sym.RawAttrConfig\]][Instance.sym.RawAttrConfig]              | symbol
[Instance\[sym.RawDesc\]][Instance.sym.RawDesc]                          | symbol
[Instance\[sym.Reference\]][Instance.sym.Reference]                      | symbol

See the [Instances section](README.md#user-content-instances) for details on the
implementation of Instances.

An Instance can be created with the [Instance.new][Instance.new] constructor.

### Instance.new
[Instance.new]: #user-content-instancenew
<code>Instance.new(className: [string](##), parent: [Instance][Instance]?, desc: [RootDesc][RootDesc]?): [Instance][Instance] \| [DataModel][DataModel]</code>

The **Instance.new** constructor returns a new Instance of the given class.
*className* sets the [ClassName][Instance.ClassName] property of the instance.
If *parent* is specified, it sets the [Parent][Instance.Parent] property.

If *desc* is specified, then it sets the [sym.Desc][Instance.sym.Desc] member.
Additionally, new will throw an error if the class does not exist. If no
descriptor is specified, then any class name will be accepted.

If *className* is "DataModel", then a [DataModel][DataModel] value is returned.
In this case, new will throw an error if *parent* is not nil.

### Instance.ClassName
[Instance.ClassName]: #user-content-instanceclassname
<code>Instance.ClassName: [string](##)</code>

The **ClassName** property gets or sets the class of the instance.

Unlike in Roblox, ClassName can be modified.

### Instance.Name
[Instance.Name]: #user-content-instancename
<code>Instance.Name: [string](##)</code>

The **Name** property gets or sets a name identifying the instance.

### Instance.Parent
[Instance.Parent]: #user-content-instanceparent
<code>Instance.Parent: [Instance][Instance]?</code>

The **Parent** property gets or sets the parent of the instance, which may be
nil.

### Instance.ClearAllChildren
[Instance.ClearAllChildren]: #user-content-instanceclearallchildren
<code>Instance:ClearAllChildren()</code>

The **ClearAllChildren** method sets the [Parent][Instance.Parent] of each child
of the instance to nil.

Unlike in Roblox, ClearAllChildren does not affect descendants.

### Instance.Clone
[Instance.Clone]: #user-content-instanceclone
<code>Instance:Clone(): [Instance][Instance]</code>

The **Clone** method returns a copy of the instance.

Unlike in Roblox, Clone does not ignore an instance if its Archivable property
is set to false.

### Instance.Descend
[Instance.Descend]: #user-content-instancedescend
<code>Instance:Descend(names: ...[string](##)): [Instance][Instance]?</code>

The **Descend** method returns a descendant of the instance by recursively
searching for each name in succession according to
[FindFirstChild][Instance.FindFirstChild]. Returns nil if a child could not be
found. Throws an error if no arguments are given.

Descend provides a safe alternative to indexing the children of an instance,
which is not implemented by rbxmk.

```lua
local face = game:Descend("Workspace", "Noob", "Head", "face")
```

### Instance.Destroy
[Instance.Destroy]: #user-content-instancedestroy
<code>Instance:Destroy()</code>

The **Destroy** method sets the [Parent][Instance.Parent] of the instance to
nil.

Unlike in Roblox, the Parent of the instance remains unlocked. Destroy also does
not affect descendants.

### Instance.FindFirstAncestor
[Instance.FindFirstAncestor]: #user-content-instancefindfirstancestor
<code>Instance:FindFirstAncestor(name: [string](##)): [Instance][Instance]?</code>

The **FindFirstAncestor** method returns the first ancestor whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found.

### Instance.FindFirstAncestorOfClass
[Instance.FindFirstAncestorOfClass]: #user-content-instancefindfirstancestorofclass
<code>Instance:FindFirstAncestorOfClass(className: [string](##)): [Instance][Instance]?</code>

The **FindFirstAncestorOfClass** method returns the first ancestor of the
instance whose [ClassName][Instance.ClassName] equals *className*, or nil if no
such instance was found.

### Instance.FindFirstAncestorWhichIsA
[Instance.FindFirstAncestorWhichIsA]: #user-content-instancefindfirstancestorwhichisa
<code>Instance:FindFirstAncestorWhichIsA(className: [string](##)): [Instance][Instance]?</code>

The **FindFirstAncestorWhichIsA** method returns the first ancestor of the
instance whose [ClassName][Instance.ClassName] inherits *className* according to
the instance's descriptor, or nil if no such instance was found. If the instance
has no descriptor, then the ClassName is compared directly.

### Instance.FindFirstChild
[Instance.FindFirstChild]: #user-content-instancefindfirstchild
<code>Instance:FindFirstChild(name: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

The **FindFirstChild** method returns the first child of the instance whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found. If
*recurse* is true, then descendants are also searched, top-down.

### Instance.FindFirstChildOfClass
[Instance.FindFirstChildOfClass]: #user-content-instancefindfirstchildofclass
<code>Instance:FindFirstChildOfClass(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

The **FindFirstChildOfClass** method returns the first child of the instance
whose [ClassName][Instance.ClassName] equals *className*, or nil if no such
instance was found. If *recurse* is true, then descendants are also searched,
top-down.

### Instance.FindFirstChildWhichIsA
[Instance.FindFirstChildWhichIsA]: #user-content-instancefindfirstchildwhichisa
<code>Instance:FindFirstChildWhichIsA(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

The **FindFirstChildWhichIsA** method returns the first child of the instance
whose [ClassName][Instance.ClassName] inherits *className*, or nil if no such
instance was found. If the instance has no descriptor, then the ClassName is
compared directly. If *recurse* is true, then descendants are also searched,
top-down.

### Instance.GetAttribute
[Instance.GetAttribute]: #user-content-instancegetattribute
<code>Instance:GetAttribute(attribute: string): Variant?</code>

The **GetAttribute** method returns the value of *attribute*, or nil if the
attribute is not found.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.GetAttributes
[Instance.GetAttributes]: #user-content-instancegetattributes
<code>Instance:GetAttributes(): Dictionary</code>

The **GetAttributes** method returns a dictionary of attribute names mapped to
values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.GetChildren
[Instance.GetChildren]: #user-content-instancegetchildren
<code>Instance:GetChildren(): Objects</code>

The **GetChildren** method returns a list of children of the instance.

### Instance.GetDescendants
[Instance.GetDescendants]: #user-content-instancegetdescendants
<code>Instance:GetDescendants(): [Objects](##)</code>

The **GetDescendants** method returns a list of descendants of the instance.

### Instance.GetFullName
[Instance.GetFullName]: #user-content-instancegetfullname
<code>Instance:GetFullName(): [string](##)</code>

The **GetFullName** method returns the concatenation of the
[Name][Instance.Name] of each ancestor of the instance and the instance itself,
separated by `.` characters. If an ancestor is a [DataModel][DataModel], it is
not included.

### Instance.IsA
[Instance.IsA]: #user-content-instanceisa
<code>Instance:IsA(className: [string](##)): [bool](##)</code>

The **IsA** method returns whether the [ClassName][Instance.ClassName] inherits
from *className*, according to the instance's descriptor. If the instance has no
descriptor, then IsA returns whether ClassName equals *className*.

### Instance.IsAncestorOf
[Instance.IsAncestorOf]: #user-content-instanceisancestorof
<code>Instance:IsAncestorOf(descendant: [Instance][Instance]): [bool](##)</code>

The **IsAncestorOf** method returns whether the instance of an ancestor of
*descendant*.

### Instance.IsDescendantOf
[Instance.IsDescendantOf]: #user-content-instanceisdescendantof
<code>Instance:IsDescendantOf(ancestor: [Instance][Instance]): [bool](##)</code>

The **IsDescendantOf** method returns whether the instance of a descendant of
*ancestor*.

### Instance.SetAttribute
[Instance.SetAttribute]: #user-content-instancesetattribute
<code>Instance:SetAttribute(attribute: string, value: Variant?)</code>

The **SetAttribute** method sets *attribute* to *value*. If *value* is nil, then
the attribute is removed.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. This
function decodes the serialized attributes, sets the given value, then
re-encodes the attributes. An error is thrown if the data could not be decoded
or encoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.SetAttributes
[Instance.SetAttributes]: #user-content-instancesetattributes
<code>Instance:SetAttributes(attributes: Dictionary)</code>

The **SetAttributes** method replaces all attributes with the content of
*attributes*, which contains attribute names mapped to values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to encode to. An error is thrown if the data could not be
encoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance[sym.AttrConfig]
[Instance.sym.AttrConfig]: #user-content-instancesymattrconfig
<code>Instance\[sym.AttrConfig\]: [AttrConfig][AttrConfig] \| [nil](##)</code>

The **AttrConfig** symbol is the [AttrConfig][AttrConfig] being used by the
instance. AttrConfig is inherited, the behavior of which is described in the
[Value inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.Desc]
[Instance.sym.Desc]: #user-content-instancesymdesc
<code>Instance\[sym.Desc\]: [RootDesc][RootDesc] \| [nil](##)</code>

The **Desc** symbol is the descriptor being used by the instance. Desc is
inherited, the behavior of which is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.IsService]
[Instance.sym.IsService]: #user-content-instancesymisservice
<code>Instance\[sym.IsService\]: [bool](##)</code>

The **IsService** symbol indicates whether the instance is a service, such as
Workspace or Lighting. This is used by some formats to determine how to encode
and decode the instance.

### Instance[sym.Properties]
[Instance.sym.Properties]: #user-content-instancesymproperties
<code>Instance\[sym.Properties\]: {\[[string](##)\]: [any](##)}</code>

The **Properties** symbol gets or sets all properties of the instance. Each
entry in the table is a property name mapped to the value of the property.

When getting, properties that would produce an error are ignored.

When setting, properties in the instance that are not in the table are removed.
If any property could not be set, then an error is thrown, and no properties are
set or removed.

### Instance[sym.RawAttrConfig]
[Instance.sym.RawAttrConfig]: #user-content-instancesymrawattrconfig
<code>Instance\[sym.RawAttrConfig\]: [AttrConfig][AttrConfig] \| [bool](##) \| [nil](##)</code>

The **RawAttrConfig** symbol is the raw member corresponding to to
[sym.AttrConfig][Instance.sym.AttrConfig]. It is similar to AttrConfig, except
that it considers only the direct value of the current instance. The exact
behavior of RawAttrConfig is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.RawDesc]
[Instance.sym.RawDesc]: #user-content-instancesymrawdesc
<code>Instance\[sym.RawDesc\]: [RootDesc][RootDesc] \| [bool](##) \| [nil](##)</code>

The **RawDesc** symbol is the raw member corresponding to to
[sym.Desc][Instance.sym.Desc]. It is similar to Desc, except that it considers
only the direct value of the current instance. The exact behavior of RawDesc is
described in the [Value inheritance](README.md#user-content-value-inheritance)
section.

### Instance[sym.Reference]
[Instance.sym.Reference]: #user-content-instancesymreference
<code>Instance\[sym.Reference\]: [string](##)</code>

The **Reference** symbol is a string used to refer to the instance from within a
[DataModel][DataModel]. Certain formats use this to encode a reference to an
instance. For example, the RBXMX format will generate random UUIDs for its
references (e.g. "RBX8B658F72923F487FAE2F7437482EF16D").

A reference should not be expected to persist when being encoded or decoded.

## Intlike
[Intlike]: #user-content-intlike

The **Intlike** type is any type that can be converted directly to an integer.
The following types are int-like:

- double
- float
- int
- int64
- token

## MemberDesc
[MemberDesc]: #user-content-memberdesc
<code>type MemberDesc = [PropertyDesc][PropertyDesc] \| [FunctionDesc][FunctionDesc] \| [EventDesc][EventDesc] \| [CallbackDesc][CallbackDesc]</code>

The **MemberDesc** is one of any of the class member descriptor types.

## Numberlike
[Numberlike]: #user-content-numberlike

The **Numberlike** type is any type that can be converted directly to a
floating-point number. The following types are number-like:

- double
- float
- int
- int64
- token

## ParameterDesc
[ParameterDesc]: #user-content-parameterdesc
<code>type ParameterDesc = {Type: [TypeDesc][TypeDesc], Name: [string](##), Default: [string](##)?}</code>

The **ParameterDesc** type describes a parameter of a function, event, or
callback member. It has the following members:

Field   | Type                 | Description
--------|----------------------|------------
Type    | [TypeDesc][TypeDesc] | The type of the parameter.
Name    | [string](##)         | The name of the parameter.
Default | [string](##)?        | Describes the default value of the parameter. If nil, then the parameter has no default value.

## PropertyDesc
[PropertyDesc]: #user-content-propertydesc
<code>type PropertyDesc = {CanLoad: [string](##), CanSave: [string](##), MemberType: [string](##), Name: [string](##), ReadSecurity: [string](##), Tags: {[string](##)}, ValueType: [TypeDesc][TypeDesc], WriteSecurity: [string](##)}</code>

The **PropertyDesc** type is a table that describes a property member of a
class. It has the following fields:

Field         | Type                 | Description
--------------|----------------------|------------
CanLoad       | [string](##)         | Whether the property is deserialized.
CanSave       | [string](##)         | Whether the property is serialized.
MemberType    | [string](##)         | Indicates the type of member. Always "Property".
Name          | [string](##)         | The name of the member.
ReadSecurity  | [string](##)         | The security context required to get the member.
Tags          | {[string](##)}       | The tags set for the member.
ValueType     | [TypeDesc][TypeDesc] | The type of the value of the property.
WriteSecurity | [string](##)         | The security context required to set the member.

## RBXAssetOptions
[RBXAssetOptions]: #user-content-rbxassetoptions
<code>type RBXAssetOptions = {AssetID: [int64](##), Cookies: [Cookies][Cookies]?, Format: [FormatSelector][FormatSelector], Body: [any](##)?}</code>

The **RBXAssetOptions** type is a table that specifies the options of a request
to an asset on the Roblox website. It has the following fields:

Field          | Type                             | Description
---------------|----------------------------------|------------
AssetID        | [int64](##)                      | The ID of the asset to request.
Cookies        | [Cookies][Cookies]?              | Optional cookies to send with requests, usually used for authentication.
Format         | [FormatSelector][FormatSelector] | The format used to encode or decode an asset.
Body           | [any](##)?                       | The body of an asset, to be encoded by the specified format.

## RootDesc
[RootDesc]: #user-content-rootdesc

The **RootDesc** type describes an entire API. It has the following members:

Member                                       | Kind
---------------------------------------------|-----
[RootDesc.Class][RootDesc.Class]             | method
[RootDesc.Classes][RootDesc.Classes]         | method
[RootDesc.ClassTag][RootDesc.ClassTag]       | method
[RootDesc.Copy][RootDesc.Copy]               | method
[RootDesc.Diff][RootDesc.Diff]               | method
[RootDesc.Enum][RootDesc.Enum]               | method
[RootDesc.EnumItem][RootDesc.EnumItem]       | method
[RootDesc.EnumItems][RootDesc.EnumItems]     | method
[RootDesc.EnumItemTag][RootDesc.EnumItemTag] | method
[RootDesc.Enums][RootDesc.Enums]             | method
[RootDesc.EnumTag][RootDesc.EnumTag]         | method
[RootDesc.EnumTypes][RootDesc.EnumTypes]     | method
[RootDesc.Member][RootDesc.Member]           | method
[RootDesc.Members][RootDesc.Members]         | method
[RootDesc.MemberTag][RootDesc.MemberTag]     | method
[RootDesc.Patch][RootDesc.Patch]             | method
[RootDesc.SetClass][RootDesc.SetClass]       | method
[RootDesc.SetEnum][RootDesc.SetEnum]         | method
[RootDesc.SetEnumItem][RootDesc.SetEnumItem] | method
[RootDesc.SetMember][RootDesc.SetMember]     | method

A RootDesc can be created with the [RootDesc.new][RootDesc.new] constructor.

### RootDesc.new
[RootDesc.new]: #user-content-rootdescnew
<code>RootDesc.new(): [RootDesc][RootDesc]</code>

The **new** constructor creates a new RootDesc.

### RootDesc.Class
[RootDesc.Class]: #user-content-rootdescclass
<code>RootDesc:Class(class: [string](##)): [ClassDesc][ClassDesc]?</code>

The **Class** method returns the class of the API corresponding to the name
*class*, or nil if the class does not exist.

### RootDesc.Classes
[RootDesc.Classes]: #user-content-rootdescclasses
<code>RootDesc:Classes(): {[string](##)}</code>

The **Classes** method returns a list of names of all the classes of the API.

### RootDesc.ClassTag
[RootDesc.ClassTag]: #user-content-rootdescclasstag
<code>RootDesc:ClassTag(class: [string](##), tag: [string](##)): [bool](##)?</code>

Returns whether *tag* is set for the class corresponding to the name *class*.
Returns nil if the class does not exist.

### RootDesc.Copy
[RootDesc.Copy]: #user-content-rootdesccopy
<code>RootDesc:Copy(): [RootDesc][RootDesc]</code>

The **Copy** method returns a deep copy of the RootDesc.

### RootDesc.Diff
[RootDesc.Diff]: #user-content-rootdescdiff
<code>RootDesc:Diff(next: [RootDesc][RootDesc]?): (diff: [DescActions][DescActions])</code>

The **Diff** method compares the root descriptor with another and returns the
differences between them. A nil value for *next* is treated the same as an empty
descriptor. The result is a list of actions that describe how to transform the
descriptor into *next*.

### RootDesc.Enum
[RootDesc.Enum]: #user-content-rootdescenum

The **Enum** method returns the enum of the API corresponding to the name
*enum*, or nil if the enum does not exist.

### RootDesc.EnumItem
[RootDesc.EnumItem]: #user-content-rootdescenumitem

The **EnumItem** method returns the enum item of the API corresponding to the
enum name *enum* and item name *item*, or nil if the enum or item does not
exist.

### RootDesc.EnumItems
[RootDesc.EnumItems]: #user-content-rootdescenumitems

The **EnumItems** method returns a list of names of all the items of the enum
corresponding to the name *enum*. Returns nil if the enum does not exist.

### RootDesc.EnumItemTag
[RootDesc.EnumItemTag]: #user-content-rootdescenumitemtag

Returns whether *tag* is set for the enum item corresponding to the name *item*
of the enum named *enum*. Returns nil if the enum or item does not exist.

### RootDesc.Enums
[RootDesc.Enums]: #user-content-rootdescenums

The **Enums** method returns a list of names of all the enums of the API.

### RootDesc.EnumTag
[RootDesc.EnumTag]: #user-content-rootdescenumtag

Returns whether *tag* is set for the enum corresponding to the name *enum*.
Returns nil if the enum does not exist.

### RootDesc.EnumTypes
[RootDesc.EnumTypes]: #user-content-rootdescenumtypes
<code>RootDesc:EnumTypes(): [Enums](##)</code>

The **EnumTypes** method returns a set of enum values generated from the current
state of the RootDesc. These enums are associated with the RootDesc, and may be
used by certain properties, so it is important to generate them before operating
on such properties. Additionally, EnumTypes should be called after modifying
enum and enum item descriptors, to regenerate the enum values.

The API of the resulting enums matches that of Roblox's Enums type. A common
pattern is to assign the result of EnumTypes to the "Enum" variable so that it
matches Roblox's API:

```lua
Enum = rootDesc:EnumTypes()
print(Enum.NormalId.Front)
```

### RootDesc.Member
[RootDesc.Member]: #user-content-rootdescmember

The **Member** method returns the class member of the API corresponding to the
class name *enum* and member name *item*, or nil if the class or member does not
exist.

### RootDesc.Members
[RootDesc.Members]: #user-content-rootdescmembers

The **Members** method returns a list of names of all the members of the class
corresponding to the name *class*. Returns nil if the class does not exist.

### RootDesc.MemberTag
[RootDesc.MemberTag]: #user-content-rootdescmembertag

Returns whether *tag* is set for the class member corresponding to the name
*member* of the class named *class*. Returns nil if the class or member does not
exist.

### RootDesc.Patch
[RootDesc.Patch]: #user-content-rootdescpatch
<code>RootDesc:Patch(actions: [DescActions][DescActions])</code>

The **Patch** method transforms the root descriptor according to a list of
actions. Each action in the list is applied in order. Actions that are
incompatible are ignored.

### RootDesc.SetClass
[RootDesc.SetClass]: #user-content-rootdescsetclass
<code>RootDesc:SetClass(class: [string](##), desc: [ClassDesc][ClassDesc]?): [bool](##)</code>

The **SetClass** method sets the fields of the class corresponding to the name
*class*. If the class already exists, then only non-nil fields are set. Fields
with the incorrect type are ignored. If *desc* is nil, then the class is
removed.

Returns false if *desc* is nil and no class exists. Returns true otherwise.

### RootDesc.SetEnum
[RootDesc.SetEnum]: #user-content-rootdescsetenum
<code>RootDesc:SetEnum(enum: [string](##), desc: [EnumDesc][EnumDesc]?): [bool](##)</code>

The **SetEnum** method sets the fields of the enum corresponding to the name
*enum*. If the enum already exists, then only non-nil fields are set. Fields
with the incorrect type are ignored. If *desc* is nil, then the enum is removed.

Returns false if *desc* is nil and no enum exists. Returns true otherwise.

### RootDesc.SetEnumItem
[RootDesc.SetEnumItem]: #user-content-rootdescsetenumitem
<code>RootDesc:SetEnumItem(enum: [string](##),item: [string](##), desc: [EnumItemDesc][EnumItemDesc]?): [bool](##)?</code>

The **SetEnumItem** method sets the fields of the enum item corresponding to the
name *item* of the enum named *enum*. If the item already exists, then only
non-nil fields are set. Fields with the incorrect type are ignored. If *desc* is
nil, then the enum is removed.

Returns nil if the enum does not exist. Returns false if *desc* is nil and no
item exists. Returns true otherwise.

### RootDesc.SetMember
[RootDesc.SetMember]: #user-content-rootdescsetmember
<code>RootDesc:SetMember(class: [string](##), member: [string](##), desc: [MemberDesc][MemberDesc]?): [bool](##)?</code>

The **SetMember** method sets the fields of the member corresponding to the name
*member* of the class named *class*. If the member already exists, then only
non-nil fields are set. Fields with the incorrect type are ignored. If *desc* is
nil, then the member is removed.

Returns nil if the class does not exist. Returns false if *desc* is nil and no
member exists. Returns true otherwise.

## Stringlike
[Stringlike]: #user-content-stringlike

The **Stringlike** type is any type that can be converted directly to a string.
The following types are string-like:

- BinaryString
- Content
- ProtectedString
- SharedString
- string

## TypeDesc
[TypeDesc]: #user-content-typedesc
<code>type TypeDesc = {Category: [string](##), Name: [string](##)}</code>

The **TypeDesc** type is a table that describes a value type. It has the
following fields:

Field      | Type         | Description
-----------|--------------|------------
Category   | [string](##) | The category of the type.
Name       | [string](##) | The name of the type.

Certain categories are treated specially:

- **Class**: Name is the name of a class. A value of the type is expected to be
  an Instance of the class.
- **Enum**: Name is the name of an enum. A value of the type is expected to be
  an enum item of the enum.

## UniqueId
[UniqueId]: #user-content-uniqueid

The **UniqueId** type represents the value of a unique identifier. It has the
following members:

Member                             | Kind
-----------------------------------|-----
[UniqueId.Index][UniqueId.Index]   | property
[UniqueId.Random][UniqueId.Random] | property
[UniqueId.Time][UniqueId.Time]     | property

A UniqueId can be created with the [UniqueId.new][UniqueId.new] constructor.

### UniqueId.new
[UniqueId.new]: #user-content-uniqueidnew
<code>UniqueId.new(random: [int64](##)?, index: [int64](##)?, index: [int64](##)?): [UniqueId][UniqueId]</code>

The **new** constructor returns a new UniqueId composed of the components from
the given arguments. If no arguments are specified, then the value is generated
from an internal source. The method to generate the value is similar to Roblox's
implementation.

### UniqueId.Index
[UniqueId.Index]: #user-content-uniqueidindex
<code>UniqueId.Index: [int64](##)</code>

The **Index** property represents the sequential portion of the unique
identifier. This value is generated such that it is almost certain to be unique,
but is also predictable.

### UniqueId.Random
[UniqueId.Random]: #user-content-uniqueidrandom
<code>UniqueId.Random: [int64](##)</code>

The **Random** property represents the random portion of the unique identifier.
This value is generated from a pseudo-random source.

### UniqueId.Time
[UniqueId.Time]: #user-content-uniqueidtime
<code>UniqueId.Time: [int64](##)</code>

The **Time** property represents the time portion of the unique identifier. This
value generated based on the time.
