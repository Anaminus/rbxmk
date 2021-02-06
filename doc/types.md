# Types
This document contains a reference to the types available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [AttrConfig][AttrConfig]
2. [CallbackDesc][CallbackDesc]
3. [ClassDesc][ClassDesc]
4. [Cookie][Cookie]
5. [Cookies][Cookies]
6. [DataModel][DataModel]
7. [DescAction][DescAction]
8. [EnumDesc][EnumDesc]
9. [EnumItemDesc][EnumItemDesc]
10. [EventDesc][EventDesc]
11. [FormatSelector][FormatSelector]
12. [FunctionDesc][FunctionDesc]
13. [HTTPHeaders][HTTPHeaders]
14. [HTTPOptions][HTTPOptions]
15. [HTTPRequest][HTTPRequest]
16. [HTTPResponse][HTTPResponse]
17. [Instance][Instance]
18. [Intlike][Intlike]
19. [Numberlike][Numberlike]
20. [ParameterDesc][ParameterDesc]
21. [PropertyDesc][PropertyDesc]
22. [RBXAssetOptions][RBXAssetOptions]
23. [RootDesc][RootDesc]
24. [Stringlike][Stringlike]
25. [TypeDesc][TypeDesc]

</td></tr></tbody>
</table>

This document only describes the types implemented by rbxmk. The
[Libraries](libraries.md#user-content-roblox) document lists the Roblox types
available in the rbxmk environment. See the [Data type
index](https://developer.roblox.com/en-us/api-reference/data-types) page on the
DevHub for information about Roblox types.

## AttrConfig
[AttrConfig]: #user-content-attrconfig

AttrConfig configures how an instance encodes and decodes attributes.

Member                          | Kind
--------------------------------|-----
[Property][AttrConfig.Property] | field

An AttrConfig can be created with the [AttrConfig.new][AttrConfig.new]
constructor.

### AttrConfig.new
[AttrConfig.new]: #user-content-attrconfignew
<code>AttrConfig.new(property: [string](##)?): [AttrConfig][AttrConfig]</code>

The `AttrConfig.new` constructor returns a new AttrConfig. *property* sets the
[Property][AttrConfig.Property] field, defaulting to an empty string.

### AttrConfig.Property
[AttrConfig.Property]: #user-content-attrconfigproperty
<code>AttrConfig.Property: [string](##)</code>

Property determines which property of an instance attributes are applied to. If
an empty string, instances will default to "AttributesSerialize".

## CallbackDesc
[CallbackDesc]: #user-content-callbackdesc

CallbackDesc describes a callback member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][CallbackDesc.Name]                   | field
[ReturnType][CallbackDesc.ReturnType]       | field
[Security][CallbackDesc.Security]           | field
[Parameters][CallbackDesc.Parameters]       | method
[SetParameters][CallbackDesc.SetParameters] | method
[Tag][CallbackDesc.Tag]                     | method
[Tags][CallbackDesc.Tags]                   | method
[SetTag][CallbackDesc.SetTag]               | method
[UnsetTag][CallbackDesc.UnsetTag]           | method

A CallbackDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### CallbackDesc.Name
[CallbackDesc.Name]: #user-content-callbackdescname
<code>CallbackDesc.Name: [string](##)</code>

Name is the name of the member.

### CallbackDesc.ReturnType
[CallbackDesc.ReturnType]: #user-content-callbackdescreturntype
<code>CallbackDesc.ReturnType: [TypeDesc][TypeDesc]</code>

ReturnType is the type returned by the callback.

### CallbackDesc.Security
[CallbackDesc.Security]: #user-content-callbackdescsecurity
<code>CallbackDesc.Security: [string](##)</code>

Security indicates the security content required to set the member.

### CallbackDesc.Parameters
[CallbackDesc.Parameters]: #user-content-callbackdescparameters
<code>CallbackDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the callback.

### CallbackDesc.SetParameters
[CallbackDesc.SetParameters]: #user-content-callbackdescsetparameters
<code>CallbackDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the callback.

### CallbackDesc.Tag
[CallbackDesc.Tag]: #user-content-callbackdesctag
<code>CallbackDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### CallbackDesc.Tags
[CallbackDesc.Tags]: #user-content-callbackdesctags
<code>CallbackDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### CallbackDesc.SetTag
[CallbackDesc.SetTag]: #user-content-callbackdescsettag
<code>CallbackDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### CallbackDesc.UnsetTag
[CallbackDesc.UnsetTag]: #user-content-callbackdescunsettag
<code>CallbackDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## ClassDesc
[ClassDesc]: #user-content-classdesc

ClassDesc describes a class. It has the following members:

Member                                     | Kind
-------------------------------------------|-----
[Name][ClassDesc.Name]                     | field
[Superclass][ClassDesc.Superclass]         | field
[MemoryCategory][ClassDesc.MemoryCategory] | field
[Member][ClassDesc.Member]                 | method
[Members][ClassDesc.Members]               | method
[AddMember][ClassDesc.AddMember]           | method
[RemoveMember][ClassDesc.RemoveMember]     | method
[Tag][ClassDesc.Tag]                       | method
[Tags][ClassDesc.Tags]                     | method
[SetTag][ClassDesc.SetTag]                 | method
[UnsetTag][ClassDesc.UnsetTag]             | method

A ClassDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### ClassDesc.Name
[ClassDesc.Name]: #user-content-classdescname
<code>ClassDesc.Name: [string](##)</code>

Name is the name of the class.

### ClassDesc.Superclass
[ClassDesc.Superclass]: #user-content-classdescsuperclass
<code>ClassDesc.Superclass: [string](##)</code>

Superclass is the name of the class from which the current class inherits.

### ClassDesc.MemoryCategory
[ClassDesc.MemoryCategory]: #user-content-classdescmemorycategory
<code>ClassDesc.MemoryCategory: [string](##)</code>

MemoryCategory describes the category of the class.

### ClassDesc.Member
[ClassDesc.Member]: #user-content-classdescmember
<code>ClassDesc:Member(name: [string](##)): [MemberDesc](##)</code>

Member returns a member of the class corresponding to the given name, or nil of
no such member exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.Members
[ClassDesc.Members]: #user-content-classdescmembers
<code>ClassDesc:Members(): {[MemberDesc](##)}</code>

Members returns a list of all the members of the class.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.AddMember
[ClassDesc.AddMember]: #user-content-classdescaddmember
<code>ClassDesc:AddMember(member: [MemberDesc](##)): [bool](##)</code>

AddMember adds a new member to the ClassDesc, returning whether the member was
added successfully. The member will fail to be added if a member of the same
name already exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

### ClassDesc.RemoveMember
[ClassDesc.RemoveMember]: #user-content-classdescremovemember
<code>ClassDesc:RemoveMember(name: [string](##)): [bool](##)</code>

RemoveMember removes a member from the ClassDesc, returning whether the member
was removed successfully. False will be returned if a member of the given name
does not exist.

### ClassDesc.Tag
[ClassDesc.Tag]: #user-content-classdesctag
<code>ClassDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### ClassDesc.Tags
[ClassDesc.Tags]: #user-content-classdesctags
<code>ClassDesc:Tags(): {[string](##)}</code>

Tags returns a list of tags that are set on the descriptor.

### ClassDesc.SetTag
[ClassDesc.SetTag]: #user-content-classdescsettag
<code>ClassDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### ClassDesc.UnsetTag
[ClassDesc.UnsetTag]: #user-content-classdescunsettag
<code>ClassDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## Cookie
[Cookie]: #user-content-cookie

A **Cookie** contains information about an HTTP cookie. It has the following
members:

Member              | Kind
--------------------|-----
[Name][Cookie.Name] | field

For security reasons, the value of the cookie cannot be accessed.

Cookie is immutable. A Cookie can be created with the
[rbxmk.newCookie](libraries.md#user-content-rbxmknewcookie) constructor.
Additionally, Cookies can be fetched from known locations with the
[rbxmk.cookiesFrom](libraries.md#user-content-rbxmkcookiesfrom) function.

### Cookie.Name
[Cookie.Name]: #user-content-cookiename
<code>Cookie.Name: [string](##)</code>

Name is the name of the cookie.

## Cookies
[Cookies]: #user-content-cookies

A **Cookies** is a list of [Cookie][Cookie] values.

## DataModel
[DataModel]: #user-content-datamodel

A **DataModel** is a special case of an [Instance][Instance]. Unlike a normal
Instance, the [ClassName][Instance.ClassName] property of a DataModel cannot be
modified, and the instance has a [GetService][DataModel.GetService] method.
Additionally, other properties are not serialized, and instead determine
metadata used by certain formats (e.g. ExplicitAutoJoints).

A DataModel can be created with the [DataModel.new][DataModel.new]
constructor.

### DataModel.new
[DataModel.new]: #user-content-datamodelnew
<code>DataModel.new(desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The `DataModel.new` constructor returns a new Instance of the DataModel class.
If *desc* is specified, then it sets the [`sym.Desc`][Instance.sym.Desc] member.

### DataModel.GetService
[DataModel.GetService]: #user-content-datamodelgetservice
<code>DataModel:GetService(className: [string](##)): [Instance][Instance]</code>

GetService returns the first child of the DataModel whose
[ClassName][Instance.ClassName] equals *className*. If no such child exists,
then a new instance of *className* is created. The [Name][Instance.Name] of the
instance is set to *className*, [`sym.IsService`][Instance.sym.IsService] is set
to true, and [Parent][Instance.Parent] is set to the DataModel.

If the DataModel has a descriptor, then GetService will throw an error if the
created class's descriptor does not have the "Service" tag set.

## DescAction
[DescAction]: #user-content-descaction

A **DescAction** describes a single action that transforms a descriptor.

Currently, DescAction has no members. However, converting a DescAction to a
string will display the content of the action in a human-readable format.

A DescAction can be created with the
[rbxmk.diffDesc](libraries.md#user-content-rbxmkdiffdesc) function. Additionally,
it is returned from the
[`desc-patch.json`](formats.md#user-content-desc-patchjson) format.

## EnumDesc
[EnumDesc]: #user-content-enumdesc

EnumDesc describes an enum. It has the following members:

Member                            | Kind
----------------------------------|-----
[Name][EnumDesc.Name]             | field
[Item][EnumDesc.Item]             | method
[Items][EnumDesc.Items]           | method
[AddItem][EnumDesc.AddItem]       | method
[RemoveItem][EnumDesc.RemoveItem] | method
[Tag][EnumDesc.Tag]               | method
[Tags][EnumDesc.Tags]             | method
[SetTag][EnumDesc.SetTag]         | method
[UnsetTag][EnumDesc.UnsetTag]     | method

An EnumDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### EnumDesc.Name
[EnumDesc.Name]: #user-content-enumdescname
<code>EnumDesc.Name: [string](##)</code>

Name is the name of the enum.

### EnumDesc.Item
[EnumDesc.Item]: #user-content-enumdescitem
<code>EnumDesc:Item(name: [string](##)): [EnumItemDesc][EnumItemDesc]</code>

Item returns an item of the enum corresponding to given name, or nil of no such
item exists.

### EnumDesc.Items
[EnumDesc.Items]: #user-content-enumdescitems
<code>EnumDesc:Items(): {[EnumItemDesc][EnumItemDesc]}</code>

Items returns a list of all the items of the enum.

### EnumDesc.AddItem
[EnumDesc.AddItem]: #user-content-enumdescadditem
<code>EnumDesc:AddItem(item: [EnumItemDesc][EnumItemDesc]): [bool](##)</code>

AddItem adds a new item to the EnumDesc, returning whether the item was added
successfully. The item will fail to be added if an item of the same name already
exists.

### EnumDesc.RemoveItem
[EnumDesc.RemoveItem]: #user-content-enumdescremoveitem
<code>EnumDesc:RemoveItem(name: [string](##)): [bool](##)</code>

RemoveItem removes an item from the EnumDesc, returning whether the item was
removed successfully. False will be returned if an item of the given name does
not exist.

### EnumDesc.Tag
[EnumDesc.Tag]: #user-content-enumdesctag
<code>EnumDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### EnumDesc.Tags
[EnumDesc.Tags]: #user-content-enumdesctags
<code>EnumDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### EnumDesc.SetTag
[EnumDesc.SetTag]: #user-content-enumdescsettag
<code>EnumDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### EnumDesc.UnsetTag
[EnumDesc.UnsetTag]: #user-content-enumdescunsettag
<code>EnumDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## EnumItemDesc
[EnumItemDesc]: #user-content-enumitemdesc

EnumDesc describes an enum item. It has the following members:

Member                            | Kind
----------------------------------|-----
[Name][EnumItemDesc.Name]         | field
[Value][EnumItemDesc.Value]       | field
[Index][EnumItemDesc.Index]       | field
[Tag][EnumItemDesc.Tag]           | method
[Tags][EnumItemDesc.Tags]         | method
[SetTag][EnumItemDesc.SetTag]     | method
[UnsetTag][EnumItemDesc.UnsetTag] | method

An EnumItemDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### EnumItemDesc.Name
[EnumItemDesc.Name]: #user-content-enumitemdescname
<code>EnumItemDesc.Name: [string](##)</code>

Name is the name of the enum item.

### EnumItemDesc.Value
[EnumItemDesc.Value]: #user-content-enumitemdescvalue
<code>EnumItemDesc.Value: [int](##)</code>

Value is the numeric value of the enum item.

### EnumItemDesc.Index
[EnumItemDesc.Index]: #user-content-enumitemdescindex
<code>EnumItemDesc.Index: [int](##)</code>

Index is an integer that hints the order of the enum item.

### EnumItemDesc.Tag
[EnumItemDesc.Tag]: #user-content-enumitemdesctag
<code>EnumItemDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### EnumItemDesc.Tags
[EnumItemDesc.Tags]: #user-content-enumitemdesctags
<code>EnumItemDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### EnumItemDesc.SetTag
[EnumItemDesc.SetTag]: #user-content-enumitemdescsettag
<code>EnumItemDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### EnumItemDesc.UnsetTag
[EnumItemDesc.UnsetTag]: #user-content-enumitemdescunsettag
<code>EnumItemDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## EventDesc
[EventDesc]: #user-content-eventdesc

EventDesc describes an event member of a class. It has the following members:

Member                                   | Kind
-----------------------------------------|-----
[Name][EventDesc.Name]                   | field
[Security][EventDesc.Security]           | field
[Parameters][EventDesc.Parameters]       | method
[SetParameters][EventDesc.SetParameters] | method
[Tag][EventDesc.Tag]                     | method
[Tags][EventDesc.Tags]                   | method
[SetTag][EventDesc.SetTag]               | method
[UnsetTag][EventDesc.UnsetTag]           | method

An EventDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### EventDesc.Name
[EventDesc.Name]: #user-content-eventdescname
<code>EventDesc.Name: [string](##)</code>

Name is the name of the member.

### EventDesc.Security
[EventDesc.Security]: #user-content-eventdescsecurity
<code>EventDesc.Security: [string](##)</code>

Security indicates the security content required to index the member.

### EventDesc.Parameters
[EventDesc.Parameters]: #user-content-eventdescparameters
<code>EventDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the event.

### EventDesc.SetParameters
[EventDesc.SetParameters]: #user-content-eventdescsetparameters
<code>EventDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the event.

### EventDesc.Tag
[EventDesc.Tag]: #user-content-eventdesctag
<code>EventDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### EventDesc.Tags
[EventDesc.Tags]: #user-content-eventdesctags
<code>EventDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### EventDesc.SetTag
[EventDesc.SetTag]: #user-content-eventdescsettag
<code>EventDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### EventDesc.UnsetTag
[EventDesc.UnsetTag]: #user-content-eventdescunsettag
<code>EventDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## FormatSelector
[FormatSelector]: #user-content-formatselector
<code>type FormatSelector = string \| {Format: string, ...}</code>

The FormatSelector type selects a [format](formats.md), and optionally
configures the format.

If a table, then the Format field indicates the name of the format to use, and
remaining fields are options that configure the format, which depend on the
format specified. All such fields are optional.

If a string, it is the name of the format to use, and specifies no options.

## FunctionDesc
[FunctionDesc]: #user-content-functiondesc

FunctionDesc describes a function member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][FunctionDesc.Name]                   | field
[ReturnType][FunctionDesc.ReturnType]       | field
[Security][FunctionDesc.Security]           | field
[Parameters][FunctionDesc.Parameters]       | method
[SetParameters][FunctionDesc.SetParameters] | method
[Tag][FunctionDesc.Tag]                     | method
[Tags][FunctionDesc.Tags]                   | method
[SetTag][FunctionDesc.SetTag]               | method
[UnsetTag][FunctionDesc.UnsetTag]           | method

A FunctionDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### FunctionDesc.Name
[FunctionDesc.Name]: #user-content-functiondescname
<code>FunctionDesc.Name: [string](##)</code>

Name is the name of the member.

### FunctionDesc.ReturnType
[FunctionDesc.ReturnType]: #user-content-functiondescreturntype
<code>FunctionDesc.ReturnType: [TypeDesc][TypeDesc]</code>

ReturnType is the type returned by the function.

### FunctionDesc.Security
[FunctionDesc.Security]: #user-content-functiondescsecurity
<code>FunctionDesc.Security: [string](##)</code>

Security indicates the security content required to index the member.

### FunctionDesc.Parameters
[FunctionDesc.Parameters]: #user-content-functiondescparameters
<code>FunctionDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the function.

### FunctionDesc.SetParameters
[FunctionDesc.SetParameters]: #user-content-functiondescsetparameters
<code>FunctionDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the function.

### FunctionDesc.Tag
[FunctionDesc.Tag]: #user-content-functiondesctag
<code>FunctionDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### FunctionDesc.Tags
[FunctionDesc.Tags]: #user-content-functiondesctags
<code>FunctionDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### FunctionDesc.SetTag
[FunctionDesc.SetTag]: #user-content-functiondescsettag
<code>FunctionDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### FunctionDesc.UnsetTag
[FunctionDesc.UnsetTag]: #user-content-functiondescunsettag
<code>FunctionDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## HTTPHeaders
[HTTPHeaders]: #user-content-httpheaders
<code>type HTTPHeaders = {\[[string](##)\]: [string](##)\|{[string](##)}}</code>

HTTPHeaders is a table that specifies the headers of an HTTP request or
response. Each entry consists of a header name mapped to a string value. If a
header requires multiple values, the name may be mapped to an array of values
instead.

For response headers, a header is always mapped to an array, and each array will
have at least one value.

## HTTPOptions
[HTTPOptions]: #user-content-httpoptions
<code>type HTTPOptions = {URL: [string](##), Method: [string](##)?, RequestFormat: [FormatSelector][FormatSelector], ResponseFormat: [FormatSelector][FormatSelector], Headers: [HTTPHeaders][HTTPHeaders]?, Cookies: [Cookies][Cookies]?, Body: [any](##)?}</code>

An HTTPOptions is a table that specifies how an HTTP request is made.

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

An HTTPRequest represents a pending HTTP request. It has the following members:

Member                         | Kind
-------------------------------|-----
[Resolve][HTTPRequest.Resolve] | method
[Cancel][HTTPRequest.Cancel]   | method

### HTTPRequest.Resolve
[HTTPRequest.Resolve]: #user-content-httprequestresolve
<code>HTTPRequest:Resolve(): (resp: [HTTPResponse][HTTPResponse])</code>

Resolve blocks until the request has finished, and returns the response. Throws
an error if a problem occurred while resolving the request.

### HTTPRequest.Cancel
[HTTPRequest.Cancel]: #user-content-httprequestcancel
<code>HTTPRequest:Cancel()</code>

Cancel cancels the pending request.

## HTTPResponse
[HTTPResponse]: #user-content-httpresponse
<code>type HTTPResponse = {Success: [bool](##), StatusCode: [int](##), StatusMessage: [string](##), Headers: [HTTPHeaders][HTTPHeaders], Cookies: [Cookies][Cookies], Body: [any](##)?}</code>

An HTTPResponse is a table that contains the response of a request.

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
getting and setting properties as described previously, instances have the
following members defined:

Member                                                          | Kind
----------------------------------------------------------------|-----
[ClassName][Instance.ClassName]                                 | property
[Name][Instance.Name]                                           | property
[Parent][Instance.Parent]                                       | property
[ClearAllChildren][Instance.ClearAllChildren]                   | method
[Clone][Instance.Clone]                                         | method
[Destroy][Instance.Destroy]                                     | method
[FindFirstAncestor][Instance.FindFirstAncestor]                 | method
[FindFirstAncestorOfClass][Instance.FindFirstAncestorOfClass]   | method
[FindFirstAncestorWhichIsA][Instance.FindFirstAncestorWhichIsA] | method
[FindFirstChild][Instance.FindFirstChild]                       | method
[FindFirstChildOfClass][Instance.FindFirstChildOfClass]         | method
[FindFirstChildWhichIsA][Instance.FindFirstChildWhichIsA]       | method
[GetAttribute][Instance.GetAttribute]                           | method
[GetAttributes][Instance.GetAttributes]                         | method
[GetChildren][Instance.GetChildren]                             | method
[GetDescendants][Instance.GetDescendants]                       | method
[GetFullName][Instance.GetFullName]                             | method
[IsA][Instance.IsA]                                             | method
[IsAncestorOf][Instance.IsAncestorOf]                           | method
[IsDescendantOf][Instance.IsDescendantOf]                       | method
[SetAttribute][Instance.SetAttribute]                           | method
[SetAttributes][Instance.SetAttributes]                         | method
[sym.AttrConfig][Instance.sym.AttrConfig]                       | symbol
[sym.Desc][Instance.sym.Desc]                                   | symbol
[sym.IsService][Instance.sym.IsService]                         | symbol
[sym.RawAttrConfig][Instance.sym.RawAttrConfig]                 | symbol
[sym.RawDesc][Instance.sym.RawDesc]                             | symbol
[sym.Reference][Instance.sym.Reference]                         | symbol

See the [Instances section](README.md#user-content-instances) for details on the
implementation of Instances.

An Instance can be created with the [Instance.new][Instance.new] constructor.

### Instance.new
[Instance.new]: #user-content-instancenew
<code>Instance.new(className: [string](##), parent: [Instance][Instance]?, desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The `Instance.new` constructor returns a new Instance of the given class.
*className* sets the [ClassName][Instance.ClassName] property of the instance.
If *parent* is specified, it sets the [Parent][Instance.Parent] property.

If *desc* is specified, then it sets the [`sym.Desc`][Instance.sym.Desc] member.
Additionally, `Instance.new` will throw an error if the class does not exist. If
no descriptor is specified, then any class name will be accepted.

### Instance.ClassName
[Instance.ClassName]: #user-content-instanceclassname
<code>Instance.ClassName: [string](##)</code>

ClassName gets or sets the class of the instance.

Unlike in Roblox, ClassName can be modified.

### Instance.Name
[Instance.Name]: #user-content-instancename
<code>Instance.Name: [string](##)</code>

Name gets or sets a name identifying the instance.

### Instance.Parent
[Instance.Parent]: #user-content-instanceparent
<code>Instance.Parent: [Instance][Instance]?</code>

Parent gets or sets the parent of the instance, which may be nil.

### Instance.ClearAllChildren
[Instance.ClearAllChildren]: #user-content-instanceclearallchildren
<code>Instance:ClearAllChildren()</code>

ClearAllChildren sets the [Parent][Instance.Parent] of each child of the
instance to nil.

Unlike in Roblox, ClearAllChildren does not affect descendants.

### Instance.Clone
[Instance.Clone]: #user-content-instanceclone
<code>Instance:Clone(): [Instance][Instance]</code>

Clone returns a copy of the instance.

Unlike in Roblox, Clone does not ignore an instance if its Archivable property
is set to false.

### Instance.Destroy
[Instance.Destroy]: #user-content-instancedestroy
<code>Instance:Destroy()</code>

Destroy sets the [Parent][Instance.Parent] of the instance to nil.

Unlike in Roblox, the Parent of the instance remains unlocked. Destroy also does
not affect descendants.

### Instance.FindFirstAncestor
[Instance.FindFirstAncestor]: #user-content-instancefindfirstancestor
<code>Instance:FindFirstAncestor(name: [string](##)): [Instance][Instance]?</code>

FindFirstAncestor returns the first ancestor whose [Name][Instance.Name] equals
*name*, or nil if no such instance was found.

### Instance.FindFirstAncestorOfClass
[Instance.FindFirstAncestorOfClass]: #user-content-instancefindfirstancestorofclass
<code>Instance:FindFirstAncestorOfClass(className: [string](##)): [Instance][Instance]?</code>

FindFirstAncestorOfClass returns the first ancestor of the instance whose
[ClassName][Instance.ClassName] equals *className*, or nil if no such instance
was found.

### Instance.FindFirstAncestorWhichIsA
[Instance.FindFirstAncestorWhichIsA]: #user-content-instancefindfirstancestorwhichisa
<code>Instance:FindFirstAncestorWhichIsA(className: [string](##)): [Instance][Instance]?</code>

FindFirstAncestorWhichIsA returns the first ancestor of the instance whose
[ClassName][Instance.ClassName] inherits *className* according to the instance's
descriptor, or nil if no such instance was found. If the instance has no
descriptor, then the ClassName is compared directly.

### Instance.FindFirstChild
[Instance.FindFirstChild]: #user-content-instancefindfirstchild
<code>Instance:FindFirstChild(name: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChild returns the first child of the instance whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found. If
*recurse* is true, then descendants are also searched, top-down.

### Instance.FindFirstChildOfClass
[Instance.FindFirstChildOfClass]: #user-content-instancefindfirstchildofclass
<code>Instance:FindFirstChildOfClass(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChildOfClass returns the first child of the instance whose
[ClassName][Instance.ClassName] equals *className*, or nil if no such instance
was found. If *recurse* is true, then descendants are also searched, top-down.

### Instance.FindFirstChildWhichIsA
[Instance.FindFirstChildWhichIsA]: #user-content-instancefindfirstchildwhichisa
<code>Instance:FindFirstChildWhichIsA(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChildWhichIsA returns the first child of the instance whose
[ClassName][Instance.ClassName] inherits *className*, or nil if no such instance
was found. If the instance has no descriptor, then the ClassName is compared
directly. If *recurse* is true, then descendants are also searched, top-down.

### Instance.GetAttribute
[Instance.GetAttribute]: #user-content-instancegetattribute
<code>Instance:GetAttribute(attribute: string): Variant?</code>

GetAttribute returns the value of *attribute*, or nil if the attribute is not
found.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [`rbxattr` format](formats.md#user-content-rbxattr) for a list of
possible attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.GetAttributes
[Instance.GetAttributes]: #user-content-instancegetattributes
<code>Instance:GetAttributes(): Dictionary</code>

GetAttributes returns a dictionary of attribute names mapped to values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [`rbxattr` format](formats.md#user-content-rbxattr) for a list of
possible attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.GetChildren
[Instance.GetChildren]: #user-content-instancegetchildren
<code>Instance:GetChildren(): Objects</code>

GetChildren returns a list of children of the instance.

### Instance.GetDescendants
[Instance.GetDescendants]: #user-content-instancegetdescendants
<code>Instance:GetDescendants(): [Objects](##)</code>

GetDescendants returns a list of descendants of the instance.

### Instance.GetFullName
[Instance.GetFullName]: #user-content-instancegetfullname
<code>Instance:GetFullName(): [string](##)</code>

GetFullName returns the concatenation of the [Name][Instance.Name] of each
ancestor of the instance and the instance itself, separated by `.` characters.
If an ancestor is a [DataModel][DataModel], it is not included.

### Instance.IsA
[Instance.IsA]: #user-content-instanceisa
<code>Instance:IsA(className: [string](##)): [bool](##)</code>

IsA returns whether the [ClassName][Instance.ClassName] inherits from
*className*, according to the instance's descriptor. If the instance has no
descriptor, then IsA returns whether ClassName equals *className*.

### Instance.IsAncestorOf
[Instance.IsAncestorOf]: #user-content-instanceisancestorof
<code>Instance:IsAncestorOf(descendant: [Instance][Instance]): [bool](##)</code>

IsAncestorOf returns whether the instance of an ancestor of *descendant*.

### Instance.IsDescendantOf
[Instance.IsDescendantOf]: #user-content-instanceisdescendantof
<code>Instance:IsDescendantOf(ancestor: [Instance][Instance]): [bool](##)</code>

IsDescendantOf returns whether the instance of a descendant of *ancestor*.

### Instance.SetAttribute
[Instance.SetAttribute]: #user-content-instancesetattribute
<code>Instance:SetAttribute(attribute: string, value: Variant?)</code>

SetAttribute sets *attribute* to *value*. If *value* is nil, then the attribute
is removed.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. This
function decodes the serialized attributes, sets the given value, then
re-encodes the attributes. An error is thrown if the data could not be decoded
or encoded.

See the [`rbxattr` format](formats.md#user-content-rbxattr) for a list of
possible attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance.SetAttributes
[Instance.SetAttributes]: #user-content-instancesetattributes
<code>Instance:SetAttributes(attributes: Dictionary)</code>

SetAttributes replaces all attributes with the content of *attributes*, which
contains attribute names mapped to values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to encode to. An error is thrown if the data could not be
encoded.

See the [`rbxattr` format](formats.md#user-content-rbxattr) for a list of
possible attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

### Instance[sym.AttrConfig]
[Instance.sym.AttrConfig]: #user-content-instancesymattrconfig
<code>Instance\[sym.AttrConfig\]: [AttrConfig][AttrConfig] \| [nil](##)</code>

AttrConfig is the [AttrConfig][AttrConfig] being used by the instance.
AttrConfig is inherited, the behavior of which is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.Desc]
[Instance.sym.Desc]: #user-content-instancesymdesc
<code>Instance\[sym.Desc\]: [RootDesc][RootDesc] \| [nil](##)</code>

Desc is the descriptor being used by the instance. Desc is inherited, the
behavior of which is described in the [Value inheritance][value-inheritance]
section.


### Instance[sym.IsService]
[Instance.sym.IsService]: #user-content-instancesymisservice
<code>Instance\[sym.IsService\]: [bool](##)</code>

IsService indicates whether the instance is a service, such as Workspace or
Lighting. This is used by some formats to determine how to encode and decode the
instance.

### Instance[sym.RawAttrConfig]
[Instance.sym.RawAttrConfig]: #user-content-instancesymrawattrconfig
<code>Instance\[sym.RawAttrConfig\]: [AttrConfig][AttrConfig] \| [bool](##) \| [nil](##)</code>

RawAttrConfig is the raw member corresponding to to
[`sym.AttrConfig`][Instance.sym.AttrConfig]. It is similar to AttrConfig, except
that it considers only the direct value of the current instance. The exact
behavior of RawAttrConfig is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.RawDesc]
[Instance.sym.RawDesc]: #user-content-instancesymrawdesc
<code>Instance\[sym.RawDesc\]: [RootDesc][RootDesc] \| [bool](##) \| [nil](##)</code>

RawDesc is the raw member corresponding to to [`sym.Desc`][Instance.sym.Desc].
It is similar to Desc, except that it considers only the direct value of the
current instance. The exact behavior of RawDesc is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

### Instance[sym.Reference]
[Instance.sym.Reference]: #user-content-instancesymreference
<code>Instance\[sym.Reference\]: [string](##)</code>

Reference is a string used to refer to the instance from within a
[DataModel][DataModel]. Certain formats use this to encode a reference to an
instance. For example, the RBXMX format will generate random UUIDs for its
references (e.g. "RBX8B658F72923F487FAE2F7437482EF16D").

## Intlike
[Intlike]: #user-content-intlike

Intlike is any type that can be converted directly to an integer. The following
types are int-like:

- double
- float
- int
- int64
- token

## Numberlike
[Numberlike]: #user-content-numberlike

Numberlike is any type that can be converted directly to a floating-point
number. The following types are number-like:

- double
- float
- int
- int64
- token

## ParameterDesc
[ParameterDesc]: #user-content-parameterdesc

ParameterDesc describes a parameter of a function, event, or callback member. It
has the following members:

Member                           | Kind
---------------------------------|-----
[Type][ParameterDesc.Type]       | field
[Name][ParameterDesc.Name]       | field
[Default][ParameterDesc.Default] | field

ParameterDesc is immutable. A new value with different fields can be created
with the [rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### ParameterDesc.Type
[ParameterDesc.Type]: #user-content-parameterdesctype
<code>ParameterDesc.Type: [TypeDesc][TypeDesc]</code>

Type is the type of the parameter.

### ParameterDesc.Name
[ParameterDesc.Name]: #user-content-parameterdescname
<code>ParameterDesc.Name: [string](##)</code>

Name is a name describing the parameter.

### ParameterDesc.Default
[ParameterDesc.Default]: #user-content-parameterdescdefault
<code>ParameterDesc.Default: [string](##)?</code>

Default is a string describing the default value of the parameter. May also be
nil, indicating that the parameter has no default value.

## PropertyDesc
[PropertyDesc]: #user-content-propertydesc

PropertyDesc describes a property member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][PropertyDesc.Name]                   | field
[ValueType][PropertyDesc.ValueType]         | field
[ReadSecurity][PropertyDesc.ReadSecurity]   | field
[WriteSecurity][PropertyDesc.WriteSecurity] | field
[CanLoad][PropertyDesc.CanLoad]             | field
[CanSave][PropertyDesc.CanSave]             | field
[Tag][PropertyDesc.Tag]                     | method
[Tags][PropertyDesc.Tags]                   | method
[SetTag][PropertyDesc.SetTag]               | method
[UnsetTag][PropertyDesc.UnsetTag]           | method

A PropertyDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### PropertyDesc.Name
[PropertyDesc.Name]: #user-content-propertydescname
<code>PropertyDesc.Name: [string](##)</code>

Name is the name of the member.

### PropertyDesc.ValueType
[PropertyDesc.ValueType]: #user-content-propertydescvaluetype
<code>PropertyDesc.ValueType: [TypeDesc][TypeDesc]</code>

ValueType is the value type of the property.

### PropertyDesc.ReadSecurity
[PropertyDesc.ReadSecurity]: #user-content-propertydescreadsecurity
<code>PropertyDesc.ReadSecurity: [string](##)</code>

ReadSecurity indicates the security context required to get the property.

### PropertyDesc.WriteSecurity
[PropertyDesc.WriteSecurity]: #user-content-propertydescwritesecurity
<code>PropertyDesc.WriteSecurity: [string](##)</code>

WriteSecurity indicates the security context required to set the property.

### PropertyDesc.CanLoad
[PropertyDesc.CanLoad]: #user-content-propertydesccanload
<code>PropertyDesc.CanLoad: [bool](##)</code>

CanLoad indicates whether the property is deserialized when decoding from a file.

### PropertyDesc.CanSave
[PropertyDesc.CanSave]: #user-content-propertydesccansave
<code>PropertyDesc.CanSave: [bool](##)</code>

CanLoad indicates whether the property is serialized when encoding to a file.

### PropertyDesc.Tag
[PropertyDesc.Tag]: #user-content-propertydesctag
<code>PropertyDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

### PropertyDesc.Tags
[PropertyDesc.Tags]: #user-content-propertydesctags
<code>PropertyDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

### PropertyDesc.SetTag
[PropertyDesc.SetTag]: #user-content-propertydescsettag
<code>PropertyDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

### PropertyDesc.UnsetTag
[PropertyDesc.UnsetTag]: #user-content-propertydescunsettag
<code>PropertyDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## RBXAssetOptions
[RBXAssetOptions]: #user-content-rbxassetoptions
<code>type RBXAssetOptions = {AssetID: [int64](##), Cookies: [Cookies][Cookies]?, Format: [FormatSelector][FormatSelector], Body: [any](##)?}</code>

An RBXAssetOptions is a table that specifies the options of a request to an
asset on the Roblox website.

Field          | Type                              | Description
---------------|-----------------------------------|------------
AssetID        | [int64](##)                       | The ID of the asset to request.
Cookies        | [Cookies][Cookies]?               | Optional cookies to send with requests, usually used for authentication.
Format         | [FormatSelector][FormatSelector]  | The format used to encode or decode an asset.
Body           | [any](##)?                        | The body of an asset, to be encoded by the specified format.

## RootDesc
[RootDesc]: #user-content-rootdesc

RootDesc describes an entire API. It has the following members:

Member                              | Kind
------------------------------------|-----
[Class][RootDesc.Class]             | method
[Classes][RootDesc.Classes]         | method
[AddClass][RootDesc.AddClass]       | method
[RemoveClass][RootDesc.RemoveClass] | method
[Enum][RootDesc.Enum]               | method
[Enums][RootDesc.Enums]             | method
[AddEnum][RootDesc.AddEnum]         | method
[RemoveEnum][RootDesc.RemoveEnum]   | method
[EnumTypes][RootDesc.EnumTypes]     | method

A RootDesc can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### RootDesc.Class
[RootDesc.Class]: #user-content-rootdescclass
<code>RootDesc:Class(name: [string](##)): [ClassDesc][ClassDesc]</code>

Class returns the class of the API corresponding to the given name, or nil if no
such class exists.

### RootDesc.Classes
[RootDesc.Classes]: #user-content-rootdescclasses
<code>RootDesc:Classes(): {[ClassDesc][ClassDesc]}</code>

Classes returns a list of all the classes of the API.

### RootDesc.AddClass
[RootDesc.AddClass]: #user-content-rootdescaddclass
<code>RootDesc:AddClass(class: [ClassDesc][ClassDesc]): [bool](##)</code>

AddClass adds a new class to the RootDesc, returning whether the class was added
successfully. The class will fail to be added if a class of the same name
already exists.

### RootDesc.RemoveClass
[RootDesc.RemoveClass]: #user-content-rootdescremoveclass
<code>RootDesc:RemoveClass(name: [string](##)): [bool](##)</code>

RemoveClass removes a class from the RootDesc, returning whether the class was
removed successfully. False will be returned if a class of the given name does
not exist.

### RootDesc.Enum
[RootDesc.Enum]: #user-content-rootdescenum
<code>RootDesc:Enum(name: [string](##)): [EnumDesc][EnumDesc]</code>

Enum returns an enum of the API corresponding to the given name, or nil if no
such enum exists.

### RootDesc.Enums
[RootDesc.Enums]: #user-content-rootdescenums
<code>RootDesc:Enums(): {[EnumDesc][EnumDesc]}</code>

Enums returns a list of all the enums of the API.

### RootDesc.AddEnum
[RootDesc.AddEnum]: #user-content-rootdescaddenum
<code>RootDesc:AddEnum(enum: [EnumDesc][EnumDesc]): [bool](##)</code>

AddEnum adds a new enum to the RootDesc, returning whether the enum was added
successfully. The enum will fail to be added if an enum of the same name already
exists.

### RootDesc.RemoveEnum
[RootDesc.RemoveEnum]: #user-content-rootdescremoveenum
<code>RootDesc:RemoveEnum(name: [string](##)): [bool](##)</code>

RemoveEnum removes an enum from the RootDesc, returning whether the enum was
removed successfully. False will be returned if an enum of the given name does
not exist.

### RootDesc.EnumTypes
[RootDesc.EnumTypes]: #user-content-rootdescenumtypes
<code>RootDesc:EnumTypes(): [Enums](##)</code>

EnumTypes returns a set of enum values generated from the current state of the
RootDesc. These enums are associated with the RootDesc, and may be used by
certain properties, so it is important to generate them before operating on such
properties. Additionally, EnumTypes should be called after modifying enum and
enum item descriptors, to regenerate the enum values.

The API of the resulting enums matches that of Roblox's Enums type. A common
pattern is to assign the result of EnumTypes to the "Enum" variable so that it
matches Roblox's API:

```lua
Enum = rootDesc:EnumTypes()
print(Enum.NormalId.Front)
```

## Stringlike
[Stringlike]: #user-content-stringlike

Stringlike is any type that can be converted directly to a string. The following
types are string-like:

- BinaryString
- Content
- ProtectedString
- SharedString
- string

## TypeDesc
[TypeDesc]: #user-content-typedesc

TypeDesc describes a value type. It has the following members:

Member                        | Kind
------------------------------|-----
[Category][TypeDesc.Category] | field
[Name][TypeDesc.Name]         | field

TypeDesc is immutable. A new value with different fields can be created with the
[rbxmk.newDesc](libraries.md#user-content-rbxmknewdesc) constructor.

### TypeDesc.Category
[TypeDesc.Category]: #user-content-typedesccategory
<code>TypeDesc.Category: [string](##)</code>

Category is the category of the type. Certain categories are treated specially:

- **Class**: Name is the name of a class. A value of the type is expected to be
  an Instance of the class.
- **Enum**: Name is the name of an enum. A value of the type is expected to be
  an enum item of the enum.

### TypeDesc.Name
[TypeDesc.Name]: #user-content-typedescname
<code>TypeDesc.Name: [string](##)</code>

Name is the name of the type.
