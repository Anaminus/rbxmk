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
9. [EnumDesc][EnumDesc]
10. [EnumItemDesc][EnumItemDesc]
11. [EventDesc][EventDesc]
12. [Faces][Faces]
13. [FormatSelector][FormatSelector]
14. [FunctionDesc][FunctionDesc]
15. [HTTPHeaders][HTTPHeaders]
16. [HTTPOptions][HTTPOptions]
17. [HTTPRequest][HTTPRequest]
18. [HTTPResponse][HTTPResponse]
19. [Instance][Instance]
20. [Intlike][Intlike]
21. [Numberlike][Numberlike]
22. [ParameterDesc][ParameterDesc]
23. [PropertyDesc][PropertyDesc]
24. [RBXAssetOptions][RBXAssetOptions]
25. [RootDesc][RootDesc]
26. [Stringlike][Stringlike]
27. [TypeDesc][TypeDesc]

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

The **CallbackDesc** type describes a callback member of a class. It has the
following members:

Member                                                   | Kind
---------------------------------------------------------|-----
[CallbackDesc.Name][CallbackDesc.Name]                   | property
[CallbackDesc.ReturnType][CallbackDesc.ReturnType]       | property
[CallbackDesc.Security][CallbackDesc.Security]           | property
[CallbackDesc.Parameters][CallbackDesc.Parameters]       | method
[CallbackDesc.SetParameters][CallbackDesc.SetParameters] | method
[CallbackDesc.Tag][CallbackDesc.Tag]                     | method
[CallbackDesc.Tags][CallbackDesc.Tags]                   | method
[CallbackDesc.SetTag][CallbackDesc.SetTag]               | method
[CallbackDesc.UnsetTag][CallbackDesc.UnsetTag]           | method

A CallbackDesc can be created with the [CallbackDesc.new][CallbackDesc.new]
constructor.

### CallbackDesc.new
[CallbackDesc.new]: #user-content-callbackdescnew
<code>CallbackDesc.new(): [CallbackDesc][CallbackDesc]</code>

The **new** constructor creates a new CallbackDesc.

### CallbackDesc.Name
[CallbackDesc.Name]: #user-content-callbackdescname
<code>CallbackDesc.Name: [string](##)</code>

The **Name** property is the name of the member.

### CallbackDesc.ReturnType
[CallbackDesc.ReturnType]: #user-content-callbackdescreturntype
<code>CallbackDesc.ReturnType: [TypeDesc][TypeDesc]</code>

The **ReturnType** property is the type returned by the callback.

### CallbackDesc.Security
[CallbackDesc.Security]: #user-content-callbackdescsecurity
<code>CallbackDesc.Security: [string](##)</code>

The **Security** property indicates the security content required to set the
member.

### CallbackDesc.Parameters
[CallbackDesc.Parameters]: #user-content-callbackdescparameters
<code>CallbackDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

The **Parameters** method returns a list of parameters of the callback.

### CallbackDesc.SetParameters
[CallbackDesc.SetParameters]: #user-content-callbackdescsetparameters
<code>CallbackDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

The **SetParameters** method sets the parameters of the callback.

### CallbackDesc.Tag
[CallbackDesc.Tag]: #user-content-callbackdesctag
<code>CallbackDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### CallbackDesc.Tags
[CallbackDesc.Tags]: #user-content-callbackdesctags
<code>CallbackDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### CallbackDesc.SetTag
[CallbackDesc.SetTag]: #user-content-callbackdescsettag
<code>CallbackDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### CallbackDesc.UnsetTag
[CallbackDesc.UnsetTag]: #user-content-callbackdescunsettag
<code>CallbackDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

## ClassDesc
[ClassDesc]: #user-content-classdesc

The **ClassDesc** type describes a class. It has the following members:

Member                                               | Kind
-----------------------------------------------------|-----
[ClassDesc.Name][ClassDesc.Name]                     | property
[ClassDesc.Superclass][ClassDesc.Superclass]         | property
[ClassDesc.MemoryCategory][ClassDesc.MemoryCategory] | property
[ClassDesc.Member][ClassDesc.Member]                 | method
[ClassDesc.Members][ClassDesc.Members]               | method
[ClassDesc.AddMember][ClassDesc.AddMember]           | method
[ClassDesc.RemoveMember][ClassDesc.RemoveMember]     | method
[ClassDesc.Tag][ClassDesc.Tag]                       | method
[ClassDesc.Tags][ClassDesc.Tags]                     | method
[ClassDesc.SetTag][ClassDesc.SetTag]                 | method
[ClassDesc.UnsetTag][ClassDesc.UnsetTag]             | method

A ClassDesc can be created with the [ClassDesc.new][ClassDesc.new] constructor.

### ClassDesc.new
[ClassDesc.new]: #user-content-classdescnew
<code>ClassDesc.new(): [ClassDesc][ClassDesc]</code>

The **new** constructor creates a new ClassDesc.

### ClassDesc.Name
[ClassDesc.Name]: #user-content-classdescname
<code>ClassDesc.Name: [string](##)</code>

The **Name** property is the name of the class.

### ClassDesc.Superclass
[ClassDesc.Superclass]: #user-content-classdescsuperclass
<code>ClassDesc.Superclass: [string](##)</code>

The **Superclass** property is the name of the class from which the current
class inherits.

### ClassDesc.MemoryCategory
[ClassDesc.MemoryCategory]: #user-content-classdescmemorycategory
<code>ClassDesc.MemoryCategory: [string](##)</code>

The **MemoryCategory** property describes the category of the class.

### ClassDesc.Member
[ClassDesc.Member]: #user-content-classdescmember
<code>ClassDesc:Member(name: [string](##)): [MemberDesc](##)</code>

The **Member** method returns a member of the class corresponding to the given
name, or nil of no such member exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.Members
[ClassDesc.Members]: #user-content-classdescmembers
<code>ClassDesc:Members(): {[MemberDesc](##)}</code>

The **Members** method returns a list of all the members of the class.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.AddMember
[ClassDesc.AddMember]: #user-content-classdescaddmember
<code>ClassDesc:AddMember(member: [MemberDesc](##)): [bool](##)</code>

The **AddMember** method adds a new member to the ClassDesc, returning whether
the member was added successfully. The member will fail to be added if a member
of the same name already exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.RemoveMember
[ClassDesc.RemoveMember]: #user-content-classdescremovemember
<code>ClassDesc:RemoveMember(name: [string](##)): [bool](##)</code>

The **RemoveMember** method removes a member from the ClassDesc, returning
whether the member was removed successfully. False will be returned if a member
of the given name does not exist.

### ClassDesc.Tag
[ClassDesc.Tag]: #user-content-classdesctag
<code>ClassDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### ClassDesc.Tags
[ClassDesc.Tags]: #user-content-classdesctags
<code>ClassDesc:Tags(): {[string](##)}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### ClassDesc.SetTag
[ClassDesc.SetTag]: #user-content-classdescsettag
<code>ClassDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### ClassDesc.UnsetTag
[ClassDesc.UnsetTag]: #user-content-classdescunsettag
<code>ClassDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

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

A DataModel can be created with the [DataModel.new][DataModel.new] constructor.

### DataModel.new
[DataModel.new]: #user-content-datamodelnew
<code>DataModel.new(desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The **DataModel.new** constructor returns a new Instance of the DataModel class.
If *desc* is specified, then it sets the [sym.Desc][Instance.sym.Desc] member.

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
constructor. It is also returned from the
[rbxmk.diffDesc](libraries.md#user-content-rbxmkdiffdesc) function and the
[desc-patch.json](formats.md#user-content-desc-patchjson) format.

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

## EnumDesc
[EnumDesc]: #user-content-enumdesc

The **EnumDesc** type describes an enum. It has the following members:

Member                                     | Kind
-------------------------------------------|-----
[EnumDesc.Name][EnumDesc.Name]             | property
[EnumDesc.Item][EnumDesc.Item]             | method
[EnumDesc.Items][EnumDesc.Items]           | method
[EnumDesc.AddItem][EnumDesc.AddItem]       | method
[EnumDesc.RemoveItem][EnumDesc.RemoveItem] | method
[EnumDesc.Tag][EnumDesc.Tag]               | method
[EnumDesc.Tags][EnumDesc.Tags]             | method
[EnumDesc.SetTag][EnumDesc.SetTag]         | method
[EnumDesc.UnsetTag][EnumDesc.UnsetTag]     | method

An EnumDesc can be created with the [EnumDesc.new][EnumDesc.new] constructor.

### EnumDesc.new
[EnumDesc.new]: #user-content-enumdescnew
<code>EnumDesc.new(): [EnumDesc][EnumDesc]</code>

The **new** constructor creates a new EnumDesc.

### EnumDesc.Name
[EnumDesc.Name]: #user-content-enumdescname
<code>EnumDesc.Name: [string](##)</code>

The **Name** property is the name of the enum.

### EnumDesc.Item
[EnumDesc.Item]: #user-content-enumdescitem
<code>EnumDesc:Item(name: [string](##)): [EnumItemDesc][EnumItemDesc]</code>

The **Item** method returns an item of the enum corresponding to given name, or
nil of no such item exists.

### EnumDesc.Items
[EnumDesc.Items]: #user-content-enumdescitems
<code>EnumDesc:Items(): {[EnumItemDesc][EnumItemDesc]}</code>

The **Items** method returns a list of all the items of the enum.

### EnumDesc.AddItem
[EnumDesc.AddItem]: #user-content-enumdescadditem
<code>EnumDesc:AddItem(item: [EnumItemDesc][EnumItemDesc]): [bool](##)</code>

The **AddItem** method adds a new item to the EnumDesc, returning whether the
item was added successfully. The item will fail to be added if an item of the
same name already exists.

### EnumDesc.RemoveItem
[EnumDesc.RemoveItem]: #user-content-enumdescremoveitem
<code>EnumDesc:RemoveItem(name: [string](##)): [bool](##)</code>

The **RemoveItem** method removes an item from the EnumDesc, returning whether
the item was removed successfully. False will be returned if an item of the
given name does not exist.

### EnumDesc.Tag
[EnumDesc.Tag]: #user-content-enumdesctag
<code>EnumDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### EnumDesc.Tags
[EnumDesc.Tags]: #user-content-enumdesctags
<code>EnumDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### EnumDesc.SetTag
[EnumDesc.SetTag]: #user-content-enumdescsettag
<code>EnumDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### EnumDesc.UnsetTag
[EnumDesc.UnsetTag]: #user-content-enumdescunsettag
<code>EnumDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

## EnumItemDesc
[EnumItemDesc]: #user-content-enumitemdesc

The **EnumItemDesc** type describes an enum item. It has the following members:

Member                                         | Kind
-----------------------------------------------|-----
[EnumItemDesc.Name][EnumItemDesc.Name]         | property
[EnumItemDesc.Value][EnumItemDesc.Value]       | property
[EnumItemDesc.Index][EnumItemDesc.Index]       | property
[EnumItemDesc.Tag][EnumItemDesc.Tag]           | method
[EnumItemDesc.Tags][EnumItemDesc.Tags]         | method
[EnumItemDesc.SetTag][EnumItemDesc.SetTag]     | method
[EnumItemDesc.UnsetTag][EnumItemDesc.UnsetTag] | method

An EnumItemDesc can be created with the [EnumItemDesc.new][EnumItemDesc.new]
constructor.

### EnumItemDesc.new
[EnumItemDesc.new]: #user-content-enumitemdescnew
<code>EnumItemDesc.new(): [EnumItemDesc][EnumItemDesc]</code>

The **new** constructor creates a new EnumItemDesc.

### EnumItemDesc.Name
[EnumItemDesc.Name]: #user-content-enumitemdescname
<code>EnumItemDesc.Name: [string](##)</code>

The **Name** property is the name of the enum item.

### EnumItemDesc.Value
[EnumItemDesc.Value]: #user-content-enumitemdescvalue
<code>EnumItemDesc.Value: [int](##)</code>

The **Value** property is the numeric value of the enum item.

### EnumItemDesc.Index
[EnumItemDesc.Index]: #user-content-enumitemdescindex
<code>EnumItemDesc.Index: [int](##)</code>

The **Index** property is an integer that hints the order of the enum item.

### EnumItemDesc.Tag
[EnumItemDesc.Tag]: #user-content-enumitemdesctag
<code>EnumItemDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### EnumItemDesc.Tags
[EnumItemDesc.Tags]: #user-content-enumitemdesctags
<code>EnumItemDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### EnumItemDesc.SetTag
[EnumItemDesc.SetTag]: #user-content-enumitemdescsettag
<code>EnumItemDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### EnumItemDesc.UnsetTag
[EnumItemDesc.UnsetTag]: #user-content-enumitemdescunsettag
<code>EnumItemDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

## EventDesc
[EventDesc]: #user-content-eventdesc

The **EventDesc** type describes an event member of a class. It has the
following members:

Member                                             | Kind
---------------------------------------------------|-----
[EventDesc.Name][EventDesc.Name]                   | property
[EventDesc.Security][EventDesc.Security]           | property
[EventDesc.Parameters][EventDesc.Parameters]       | method
[EventDesc.SetParameters][EventDesc.SetParameters] | method
[EventDesc.Tag][EventDesc.Tag]                     | method
[EventDesc.Tags][EventDesc.Tags]                   | method
[EventDesc.SetTag][EventDesc.SetTag]               | method
[EventDesc.UnsetTag][EventDesc.UnsetTag]           | method

An EventDesc can be created with the [EventDesc.new][EventDesc.new] constructor.

### EventDesc.new
[EventDesc.new]: #user-content-eventdescnew
<code>EventDesc.new(): [EventDesc][EventDesc]</code>

The **new** constructor creates a new EventDesc.

### EventDesc.Name
[EventDesc.Name]: #user-content-eventdescname
<code>EventDesc.Name: [string](##)</code>

The **Name** property is the name of the member.

### EventDesc.Security
[EventDesc.Security]: #user-content-eventdescsecurity
<code>EventDesc.Security: [string](##)</code>

The **Security** property indicates the security content required to index the
member.

### EventDesc.Parameters
[EventDesc.Parameters]: #user-content-eventdescparameters
<code>EventDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

The **Parameters** method returns a list of parameters of the event.

### EventDesc.SetParameters
[EventDesc.SetParameters]: #user-content-eventdescsetparameters
<code>EventDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

The **SetParameters** method sets the parameters of the event.

### EventDesc.Tag
[EventDesc.Tag]: #user-content-eventdesctag
<code>EventDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### EventDesc.Tags
[EventDesc.Tags]: #user-content-eventdesctags
<code>EventDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### EventDesc.SetTag
[EventDesc.SetTag]: #user-content-eventdescsettag
<code>EventDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### EventDesc.UnsetTag
[EventDesc.UnsetTag]: #user-content-eventdescunsettag
<code>EventDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

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

The **FunctionDesc** type describes a function member of a class. It has the
following members:

Member                                                   | Kind
---------------------------------------------------------|-----
[FunctionDesc.Name][FunctionDesc.Name]                   | property
[FunctionDesc.ReturnType][FunctionDesc.ReturnType]       | property
[FunctionDesc.Security][FunctionDesc.Security]           | property
[FunctionDesc.Parameters][FunctionDesc.Parameters]       | method
[FunctionDesc.SetParameters][FunctionDesc.SetParameters] | method
[FunctionDesc.Tag][FunctionDesc.Tag]                     | method
[FunctionDesc.Tags][FunctionDesc.Tags]                   | method
[FunctionDesc.SetTag][FunctionDesc.SetTag]               | method
[FunctionDesc.UnsetTag][FunctionDesc.UnsetTag]           | method

A FunctionDesc can be created with the [FunctionDesc.new][FunctionDesc.new]
constructor.

### FunctionDesc.new
[FunctionDesc.new]: #user-content-callbackdescnew
<code>FunctionDesc.new(): [FunctionDesc][FunctionDesc]</code>

The **new** constructor creates a new FunctionDesc.

### FunctionDesc.Name
[FunctionDesc.Name]: #user-content-functiondescname
<code>FunctionDesc.Name: [string](##)</code>

The **Name** property is the name of the member.

### FunctionDesc.ReturnType
[FunctionDesc.ReturnType]: #user-content-functiondescreturntype
<code>FunctionDesc.ReturnType: [TypeDesc][TypeDesc]</code>

The **ReturnType** property is the type returned by the function.

### FunctionDesc.Security
[FunctionDesc.Security]: #user-content-functiondescsecurity
<code>FunctionDesc.Security: [string](##)</code>

The **Security** property indicates the security content required to index the
member.

### FunctionDesc.Parameters
[FunctionDesc.Parameters]: #user-content-functiondescparameters
<code>FunctionDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

The **Parameters** method returns a list of parameters of the function.

### FunctionDesc.SetParameters
[FunctionDesc.SetParameters]: #user-content-functiondescsetparameters
<code>FunctionDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

The **SetParameters** method sets the parameters of the function.

### FunctionDesc.Tag
[FunctionDesc.Tag]: #user-content-functiondesctag
<code>FunctionDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### FunctionDesc.Tags
[FunctionDesc.Tags]: #user-content-functiondesctags
<code>FunctionDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### FunctionDesc.SetTag
[FunctionDesc.SetTag]: #user-content-functiondescsettag
<code>FunctionDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### FunctionDesc.UnsetTag
[FunctionDesc.UnsetTag]: #user-content-functiondescunsettag
<code>FunctionDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

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
<code>Instance.new(className: [string](##), parent: [Instance][Instance]?, desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The **Instance.new** constructor returns a new Instance of the given class.
*className* sets the [ClassName][Instance.ClassName] property of the instance.
If *parent* is specified, it sets the [Parent][Instance.Parent] property.

If *desc* is specified, then it sets the [sym.Desc][Instance.sym.Desc] member.
Additionally, new will throw an error if the class does not exist. If no
descriptor is specified, then any class name will be accepted.

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

The **ParameterDesc** type describes a parameter of a function, event, or
callback member. It has the following members:

Member                                         | Kind
-----------------------------------------------|-----
[ParameterDesc.Type][ParameterDesc.Type]       | field
[ParameterDesc.Name][ParameterDesc.Name]       | field
[ParameterDesc.Default][ParameterDesc.Default] | field

ParameterDesc is immutable. A new value with different fields can be created
with the [ParameterDesc.new][ParameterDesc.new] constructor.

### ParameterDesc.new
[ParameterDesc.new]: #user-content-parameterdescnew
<code>ParameterDesc.new(type: [TypeDesc][TypeDesc]?, name: [string](##)?, default: [string](##)?): [ParameterDesc][ParameterDesc]</code>

The **new** constructor creates a new ParameterDesc. *type* sets the
[Type][ParameterDesc.Type] property, if specified. *name* sets the
[Name][ParameterDesc.Name] property, if specified. *default* sets the
[Default][ParameterDesc.Default] property, if specified.

### ParameterDesc.Type
[ParameterDesc.Type]: #user-content-parameterdesctype
<code>ParameterDesc.Type: [TypeDesc][TypeDesc]</code>

The **Type** field is the type of the parameter.

### ParameterDesc.Name
[ParameterDesc.Name]: #user-content-parameterdescname
<code>ParameterDesc.Name: [string](##)</code>

The **Name** field is a name describing the parameter.

### ParameterDesc.Default
[ParameterDesc.Default]: #user-content-parameterdescdefault
<code>ParameterDesc.Default: [string](##)?</code>

The **Default** field is a string describing the default value of the parameter.
May also be nil, indicating that the parameter has no default value.

## PropertyDesc
[PropertyDesc]: #user-content-propertydesc

The **PropertyDesc** type describes a property member of a class. It has the
following members:

Member                                                   | Kind
---------------------------------------------------------|-----
[PropertyDesc.Name][PropertyDesc.Name]                   | property
[PropertyDesc.ValueType][PropertyDesc.ValueType]         | property
[PropertyDesc.ReadSecurity][PropertyDesc.ReadSecurity]   | property
[PropertyDesc.WriteSecurity][PropertyDesc.WriteSecurity] | property
[PropertyDesc.CanLoad][PropertyDesc.CanLoad]             | property
[PropertyDesc.CanSave][PropertyDesc.CanSave]             | property
[PropertyDesc.Tag][PropertyDesc.Tag]                     | method
[PropertyDesc.Tags][PropertyDesc.Tags]                   | method
[PropertyDesc.SetTag][PropertyDesc.SetTag]               | method
[PropertyDesc.UnsetTag][PropertyDesc.UnsetTag]           | method

A PropertyDesc can be created with the [PropertyDesc.new][PropertyDesc.new]
constructor.

### PropertyDesc.new
[PropertyDesc.new]: #user-content-propertydescnew
<code>PropertyDesc.new(): [PropertyDesc][PropertyDesc]</code>

The **new** constructor creates a new PropertyDesc.

### PropertyDesc.Name
[PropertyDesc.Name]: #user-content-propertydescname
<code>PropertyDesc.Name: [string](##)</code>

The **Name** property is the name of the member.

### PropertyDesc.ValueType
[PropertyDesc.ValueType]: #user-content-propertydescvaluetype
<code>PropertyDesc.ValueType: [TypeDesc][TypeDesc]</code>

The **ValueType** property is the value type of the property.

### PropertyDesc.ReadSecurity
[PropertyDesc.ReadSecurity]: #user-content-propertydescreadsecurity
<code>PropertyDesc.ReadSecurity: [string](##)</code>

The **ReadSecurity** property indicates the security context required to get the
property.

### PropertyDesc.WriteSecurity
[PropertyDesc.WriteSecurity]: #user-content-propertydescwritesecurity
<code>PropertyDesc.WriteSecurity: [string](##)</code>

The **WriteSecurity** property indicates the security context required to set
the property.

### PropertyDesc.CanLoad
[PropertyDesc.CanLoad]: #user-content-propertydesccanload
<code>PropertyDesc.CanLoad: [bool](##)</code>

The **CanLoad** property indicates whether the property is deserialized when
decoding from a file.

### PropertyDesc.CanSave
[PropertyDesc.CanSave]: #user-content-propertydesccansave
<code>PropertyDesc.CanSave: [bool](##)</code>

The **CanSave** property indicates whether the property is serialized when
encoding to a file.

### PropertyDesc.Tag
[PropertyDesc.Tag]: #user-content-propertydesctag
<code>PropertyDesc:Tag(name: [string](##)): [bool](##)</code>

The **Tag** method returns whether a tag of the given name is set on the
descriptor.

### PropertyDesc.Tags
[PropertyDesc.Tags]: #user-content-propertydesctags
<code>PropertyDesc:Tags(): {string}</code>

The **Tags** method returns a list of tags that are set on the descriptor.

### PropertyDesc.SetTag
[PropertyDesc.SetTag]: #user-content-propertydescsettag
<code>PropertyDesc:SetTag(tags: ...[string](##))</code>

The **SetTag** method sets the given tags on the descriptor.

### PropertyDesc.UnsetTag
[PropertyDesc.UnsetTag]: #user-content-propertydescunsettag
<code>PropertyDesc:UnsetTag(tags: ...[string](##))</code>

The **UnsetTag** method unsets the given tags on the descriptor.

## RBXAssetOptions
[RBXAssetOptions]: #user-content-rbxassetoptions
<code>type RBXAssetOptions = {AssetID: [int64](##), Cookies: [Cookies][Cookies]?, Format: [FormatSelector][FormatSelector], Body: [any](##)?}</code>

The **RBXAssetOptions** type is a table that specifies the options of a request
to an asset on the Roblox website. It has the following fields:

Field          | Type                              | Description
---------------|-----------------------------------|------------
AssetID        | [int64](##)                       | The ID of the asset to request.
Cookies        | [Cookies][Cookies]?               | Optional cookies to send with requests, usually used for authentication.
Format         | [FormatSelector][FormatSelector]  | The format used to encode or decode an asset.
Body           | [any](##)?                        | The body of an asset, to be encoded by the specified format.

## RootDesc
[RootDesc]: #user-content-rootdesc

The **RootDesc** type describes an entire API. It has the following members:

Member                                       | Kind
---------------------------------------------|-----
[RootDesc.Class][RootDesc.Class]             | method
[RootDesc.Classes][RootDesc.Classes]         | method
[RootDesc.AddClass][RootDesc.AddClass]       | method
[RootDesc.RemoveClass][RootDesc.RemoveClass] | method
[RootDesc.Enum][RootDesc.Enum]               | method
[RootDesc.Enums][RootDesc.Enums]             | method
[RootDesc.AddEnum][RootDesc.AddEnum]         | method
[RootDesc.RemoveEnum][RootDesc.RemoveEnum]   | method
[RootDesc.EnumTypes][RootDesc.EnumTypes]     | method

A RootDesc can be created with the [RootDesc.new][RootDesc.new] constructor.

### RootDesc.new
[RootDesc.new]: #user-content-rootdescnew
<code>RootDesc.new(): [RootDesc][RootDesc]</code>

The **new** constructor creates a new RootDesc.

### RootDesc.Class
[RootDesc.Class]: #user-content-rootdescclass
<code>RootDesc:Class(name: [string](##)): [ClassDesc][ClassDesc]</code>

The **Class** method returns the class of the API corresponding to the given
name, or nil if no such class exists.

### RootDesc.Classes
[RootDesc.Classes]: #user-content-rootdescclasses
<code>RootDesc:Classes(): {[ClassDesc][ClassDesc]}</code>

The **Classes** method returns a list of all the classes of the API.

### RootDesc.AddClass
[RootDesc.AddClass]: #user-content-rootdescaddclass
<code>RootDesc:AddClass(class: [ClassDesc][ClassDesc]): [bool](##)</code>

The **AddClass** method adds a new class to the RootDesc, returning whether the
class was added successfully. The class will fail to be added if a class of the
same name already exists.

### RootDesc.RemoveClass
[RootDesc.RemoveClass]: #user-content-rootdescremoveclass
<code>RootDesc:RemoveClass(name: [string](##)): [bool](##)</code>

The **RemoveClass** method removes a class from the RootDesc, returning whether
the class was removed successfully. False will be returned if a class of the
given name does not exist.

### RootDesc.Enum
[RootDesc.Enum]: #user-content-rootdescenum
<code>RootDesc:Enum(name: [string](##)): [EnumDesc][EnumDesc]</code>

The **Enum** method returns an enum of the API corresponding to the given name,
or nil if no such enum exists.

### RootDesc.Enums
[RootDesc.Enums]: #user-content-rootdescenums
<code>RootDesc:Enums(): {[EnumDesc][EnumDesc]}</code>

The **Enums** method returns a list of all the enums of the API.

### RootDesc.AddEnum
[RootDesc.AddEnum]: #user-content-rootdescaddenum
<code>RootDesc:AddEnum(enum: [EnumDesc][EnumDesc]): [bool](##)</code>

The **AddEnum** method adds a new enum to the RootDesc, returning whether the
enum was added successfully. The enum will fail to be added if an enum of the
same name already exists.

### RootDesc.RemoveEnum
[RootDesc.RemoveEnum]: #user-content-rootdescremoveenum
<code>RootDesc:RemoveEnum(name: [string](##)): [bool](##)</code>

The **RemoveEnum** method removes an enum from the RootDesc, returning whether
the enum was removed successfully. False will be returned if an enum of the
given name does not exist.

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

The **TypeDesc** type describes a value type. It has the following members:

Member                                 | Kind
---------------------------------------|-----
[TypeDesc.Category][TypeDesc.Category] | field
[TypeDesc.Name][TypeDesc.Name]         | field

TypeDesc is immutable. A new value with different fields can be created with the
[TypeDesc.new][TypeDesc.new] constructor.

### TypeDesc.new
[TypeDesc.new]: #user-content-callbackdescnew
<code>TypeDesc.new(category: [string](##), name: [string](##)): [TypeDesc][TypeDesc]</code>

The **new** constructor creates a new TypeDesc. *category* sets the
[Category][TypeDesc.Category] property, if specified. *name* sets the
[Name][TypeDesc.Name] property, if specified.

### TypeDesc.Category
[TypeDesc.Category]: #user-content-typedesccategory
<code>TypeDesc.Category: [string](##)</code>

The **Category** field is the category of the type. Certain categories are
treated specially:

- **Class**: Name is the name of a class. A value of the type is expected to be
  an Instance of the class.
- **Enum**: Name is the name of an enum. A value of the type is expected to be
  an enum item of the enum.

### TypeDesc.Name
[TypeDesc.Name]: #user-content-typedescname
<code>TypeDesc.Name: [string](##)</code>

The **Name** field is the name of the type.
