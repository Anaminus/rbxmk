T.Pass(typeof(RootDesc.new()) == "RootDesc", "new returns RootDesc")

local desc = RootDesc.new()

-- Metamethod tests
T.Pass(typeof(desc) == "RootDesc"                   , "type of descriptor is RootDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= RootDesc.new()                       , "descriptor is not equal to another descriptor of the same type")

-- Classes
local classA = {Name = "ClassA"}
local classB = {Name = "ClassB"}
local classC = {Name = "ClassC",Superclass = "",MemoryCategory = "",Tags={"Foo","Bar"}}
local classD = {Name = "ClassD"}
T.Equal("Classes returns empty table", desc:Classes(), {})
T.Fail(function() desc:Class() end, "Class with no values")
T.Pass(function() desc:Class("") end, "Class(1) with string")
T.Fail(function() desc:Class(42) end, "Class(1) with non-string")
T.Fail(function() desc:SetClass() end, "SetClass with no values")
T.Fail(function() desc:SetClass(42) end, "SetClass(1) with non-string")
T.Pass(function() desc:SetClass("") end, "SetClass(1) with string")
T.Pass(function() desc:SetClass(classA.Name, nil) end, "SetClass(2) with nil")
T.Fail(function() desc:SetClass(classA.Name, 42) end, "SetClass(2) with non-string")
T.Pass(function() return desc:SetClass(classA.Name, classA) == true end, "SetClass(2) with table returns true")
T.Pass(function() return desc:SetClass(classA.Name, classA) == true end, "SetClass(2) with table again returns true")
T.Pass(function() return desc:SetClass(classA.Name, nil) == true end, "SetClass(2) with nil returns true")
T.Pass(function() return desc:SetClass(classA.Name, nil) == false end, "SetClass(2) with nil again returns false")
desc:SetClass(classA.Name, classA)
T.Pass(function() return desc:Class("") == nil end, "Class: non-extant class returns nil")
T.Pass(function() return type(desc:Class(classA.Name)) == "table" end, "Class: extant class returns table")
T.Pass(function() return desc:Class(classA.Name).Name == classA.Name end, "Class: class name matches argument")
T.Pass(function() desc:SetClass("Foo", {Name="Bar"}) end, "set class Foo with name Bar")
T.Pass(function() return desc:Class("Foo").Name == "Foo" end, "name argument overrides Name field")
T.Pass(function() desc:SetClass("Foo", nil) end, "set class Foo to nil")
T.Pass(function() return desc:Class("Foo") == nil end, "class Foo is removed")
desc:SetClass(classB.Name, classB)
desc:SetClass(classC.Name, classC)
desc:SetClass(classD.Name, classD)
T.Equal("class with tags", desc:Class("ClassC"), classC)
T.Equal("Classes returns class names", desc:Classes(), {"ClassA","ClassB","ClassC","ClassD"})
T.Fail(function() desc:ClassTag() end, "ClassTag with no values")
T.Fail(function() desc:ClassTag("") end, "ClassTag with one string")
T.Pass(function() desc:ClassTag("","") end, "ClassTag with two strings")
T.Fail(function() desc:ClassTag(42,"") end, "ClassTag(1) with non-string")
T.Fail(function() desc:ClassTag("",42) end, "ClassTag(2) with non-string")
T.Pass(function() return desc:ClassTag("Nonextant","Foo") == nil end, "nonextant class tag returns nil")
T.Pass(function() return desc:ClassTag("ClassA","Foo") == false end, "unset class tag returns false")
T.Pass(function() return desc:ClassTag("ClassC","Foo") == true end, "set class tag returns true")

-- Members
local memberA = {MemberType="Property",Name = "MemberA"}
local memberB = {MemberType="Function",Name = "MemberB"}
local memberC = {MemberType="Event",Name = "MemberC",Tags={"Foo","Bar"},Security="",ThreadSafety="",Parameters={}}
local memberD = {MemberType="Callback",Name = "MemberD"}
T.Fail(function() desc:Members() end, "Members with no values")
T.Fail(function() desc:Members(42) end, "Members with non-string")
T.Pass(function() desc:Members("") end, "Members with string")
T.Equal("Members returns empty table", desc:Members("ClassA"), {})
T.Fail(function() desc:Member() end, "Member with no values")
T.Fail(function() desc:Member("") end, "Member with one string")
T.Pass(function() desc:Member("","") end, "Member with two strings")
T.Fail(function() desc:Member(42,"") end, "Member(1) with non-string")
T.Fail(function() desc:Member("",42) end, "Member(2) with non-string")
T.Fail(function() desc:SetMember() end, "SetMember with no values")
T.Fail(function() desc:SetMember("") end, "SetMember with one string")
T.Pass(function() desc:SetMember("","") end, "SetMember with two strings")
T.Fail(function() desc:SetMember(42,"") end, "SetMember(1) with non-string")
T.Fail(function() desc:SetMember("",42) end, "SetMember(2) with non-string")
T.Pass(function() return desc:Member("Nonextant", "Nonextant") == nil end, "Member(no, no) returns nil")
T.Pass(function() return desc:Member("ClassA", "Nonextant") == nil end, "Member(yes, no) returns nil")
T.Pass(function() return desc:SetMember("ClassA", "MemberA", memberA) == true end, "SetMember(yes, yes, yes) returns true")
T.Pass(function() return desc:SetMember("ClassA", "MemberA", memberA) == true end, "SetMember(yes, yes, yes) again returns true")
T.Pass(function() return desc:SetMember("ClassA", "MemberA", nil) == true end, "SetMember(yes, yes, no) returns true")
T.Pass(function() return desc:SetMember("ClassA", "MemberA", nil) == false end, "SetMember(yes, yes, no) returns again false")
T.Pass(function() return desc:SetMember("Nonextant", "MemberA", nil) == nil end, "SetMember(no, no, no) returns nil")
T.Pass(function() return desc:SetMember("Nonextant", "MemberA", memberA) == nil end, "SetMember(no, no, yes) returns nil")
T.Fail(function() desc:SetMember("ClassA", "MemberE", {Name = "MemberE"}) end, "SetMember table must have MemberType")
T.Fail(function() desc:SetMember("ClassA", "MemberE", {MemberType="Invalid",Name = "MemberE"}) end, "SetMember table must have valid MemberType")
T.Pass(function() desc:SetMember("ClassA", memberA.Name, nil) end, "SetMember(2) with nil")
T.Fail(function() desc:SetMember("ClassA", memberA.Name, 42) end, "SetMember(2) with non-string")
T.Pass(function() desc:SetMember("ClassA", memberA.Name, memberA) end, "SetMember(2) with table")
T.Pass(function() return desc:Member("ClassA","") == nil end, "Member: non-extant member returns nil")
T.Pass(function() return type(desc:Member("ClassA",memberA.Name)) == "table" end, "Member: extant member returns table")
T.Pass(function() return desc:Member("ClassA",memberA.Name).Name == memberA.Name end, "Member: member name matches argument")
T.Pass(function() desc:SetMember("ClassA","Foo", {MemberType="Property",Name="Bar"}) end, "set member Foo with name Bar")
T.Pass(function() return desc:Member("ClassA","Foo").Name == "Foo" end, "name argument overrides Name field")
T.Pass(function() desc:SetMember("ClassA","Foo", nil) end, "set member Foo to nil")
T.Pass(function() return desc:Member("ClassA","Foo") == nil end, "member Foo is removed")
desc:SetMember("ClassA",memberB.Name, memberB)
desc:SetMember("ClassA",memberC.Name, memberC)
desc:SetMember("ClassA",memberD.Name, memberD)
T.Equal("member with tags", desc:Member("ClassA","MemberC"), memberC)
T.Equal("Members returns member names", desc:Members("ClassA"), {"MemberA","MemberB","MemberC","MemberD"})
T.Fail(function() desc:MemberTag("ClassA") end, "MemberTag with no values")
T.Fail(function() desc:MemberTag("ClassA","") end, "MemberTag with one string")
T.Pass(function() desc:MemberTag("ClassA","","") end, "MemberTag with two strings")
T.Fail(function() desc:MemberTag("ClassA",42,"") end, "MemberTag(1) with non-string")
T.Fail(function() desc:MemberTag("ClassA","",42) end, "MemberTag(2) with non-string")
T.Pass(function() return desc:MemberTag("ClassA","Nonextant","Foo") == nil end, "nonextant member tag returns nil")
T.Pass(function() return desc:MemberTag("ClassA","MemberA","Foo") == false end, "unset member tag returns false")
T.Pass(function() return desc:MemberTag("ClassA","MemberC","Foo") == true end, "set member tag returns true")

-- Enums
local enumA = {Name = "EnumA"}
local enumB = {Name = "EnumB"}
local enumC = {Name = "EnumC",Tags={"Foo","Bar"}}
local enumD = {Name = "EnumD"}
T.Equal("Enums returns empty table", desc:Enums(), {})
T.Fail(function() desc:Enum() end, "Enum with no values")
T.Pass(function() desc:Enum("") end, "Enum(1) with string")
T.Fail(function() desc:Enum(42) end, "Enum(1) with non-string")
T.Fail(function() desc:SetEnum() end, "SetEnum with no values")
T.Fail(function() desc:SetEnum(42) end, "SetEnum(1) with non-string")
T.Pass(function() desc:SetEnum("") end, "SetEnum(1) with string")
T.Pass(function() desc:SetEnum(enumA.Name, nil) end, "SetEnum(2) with nil")
T.Fail(function() desc:SetEnum(enumA.Name, 42) end, "SetEnum(2) with non-string")
T.Pass(function() return desc:SetEnum(enumA.Name, enumA) == true end, "SetEnum(2) with table returns true")
T.Pass(function() return desc:SetEnum(enumA.Name, enumA) == true end, "SetEnum(2) with table again returns true")
T.Pass(function() return desc:SetEnum(enumA.Name, nil) == true end, "SetEnum(2) with nil returns true")
T.Pass(function() return desc:SetEnum(enumA.Name, nil) == false end, "SetEnum(2) with nil again returns false")
desc:SetEnum(enumA.Name, enumA)
T.Pass(function() return desc:Enum("") == nil end, "Enum: non-extant enum returns nil")
T.Pass(function() return type(desc:Enum(enumA.Name)) == "table" end, "Enum: extant enum returns table")
T.Pass(function() return desc:Enum(enumA.Name).Name == enumA.Name end, "Enum: enum name matches argument")
T.Pass(function() desc:SetEnum("Foo", {Name="Bar"}) end, "set enum Foo with name Bar")
T.Pass(function() return desc:Enum("Foo").Name == "Foo" end, "name argument overrides Name field")
T.Pass(function() desc:SetEnum("Foo", nil) end, "set enum Foo to nil")
T.Pass(function() return desc:Enum("Foo") == nil end, "enum Foo is removed")
desc:SetEnum(enumB.Name, enumB)
desc:SetEnum(enumC.Name, enumC)
desc:SetEnum(enumD.Name, enumD)
T.Equal("enum with tags", desc:Enum("EnumC"), enumC)
T.Equal("Enums returns enum names", desc:Enums(), {"EnumA","EnumB","EnumC","EnumD"})
T.Fail(function() desc:EnumTag() end, "EnumTag with no values")
T.Fail(function() desc:EnumTag("") end, "EnumTag with one string")
T.Pass(function() desc:EnumTag("","") end, "EnumTag with two strings")
T.Fail(function() desc:EnumTag(42,"") end, "EnumTag(1) with non-string")
T.Fail(function() desc:EnumTag("",42) end, "EnumTag(2) with non-string")
T.Pass(function() return desc:EnumTag("Nonextant","Foo") == nil end, "nonextant enum tag returns nil")
T.Pass(function() return desc:EnumTag("EnumA","Foo") == false end, "unset enum tag returns false")
T.Pass(function() return desc:EnumTag("EnumC","Foo") == true end, "set enum tag returns true")

-- EnumItems
local enumitemA = {Name = "EnumItemA"}
local enumitemB = {Name = "EnumItemB"}
local enumitemC = {Name = "EnumItemC",Tags={"Foo","Bar"},Value=0,Index=0}
local enumitemD = {Name = "EnumItemD"}
T.Fail(function() desc:EnumItems() end, "EnumItems with no values")
T.Fail(function() desc:EnumItems(42) end, "EnumItems with non-string")
T.Pass(function() desc:EnumItems("") end, "EnumItems with string")
T.Equal("EnumItems returns empty table", desc:EnumItems("EnumA"), {})
T.Fail(function() desc:EnumItem() end, "EnumItem with no values")
T.Fail(function() desc:EnumItem("") end, "EnumItem with one string")
T.Pass(function() desc:EnumItem("","") end, "EnumItem with two strings")
T.Fail(function() desc:EnumItem(42,"") end, "EnumItem(1) with non-string")
T.Fail(function() desc:EnumItem("",42) end, "EnumItem(2) with non-string")
T.Fail(function() desc:SetEnumItem() end, "SetEnumItem with no values")
T.Fail(function() desc:SetEnumItem("") end, "SetEnumItem with one string")
T.Pass(function() desc:SetEnumItem("","") end, "SetEnumItem with two strings")
T.Fail(function() desc:SetEnumItem(42,"") end, "SetEnumItem(1) with non-string")
T.Fail(function() desc:SetEnumItem("",42) end, "SetEnumItem(2) with non-string")
T.Pass(function() return desc:EnumItem("Nonextant", "Nonextant") == nil end, "EnumItem(no, no) returns nil")
T.Pass(function() return desc:EnumItem("EnumA", "Nonextant") == nil end, "EnumItem(yes, no) returns nil")
T.Pass(function() return desc:SetEnumItem("EnumA", "EnumItemA", enumitemA) == true end, "SetEnumItem(yes, yes, yes) returns true")
T.Pass(function() return desc:SetEnumItem("EnumA", "EnumItemA", enumitemA) == true end, "SetEnumItem(yes, yes, yes) again returns true")
T.Pass(function() return desc:SetEnumItem("EnumA", "EnumItemA", nil) == true end, "SetEnumItem(yes, yes, no) returns true")
T.Pass(function() return desc:SetEnumItem("EnumA", "EnumItemA", nil) == false end, "SetEnumItem(yes, yes, no) returns again false")
T.Pass(function() return desc:SetEnumItem("Nonextant", "EnumItemA", nil) == nil end, "SetEnumItem(no, no, no) returns nil")
T.Pass(function() return desc:SetEnumItem("Nonextant", "EnumItemA", enumitemA) == nil end, "SetEnumItem(no, no, yes) returns nil")
T.Pass(function() desc:SetEnumItem("EnumA", enumitemA.Name, nil) end, "SetEnumItem(2) with nil")
T.Fail(function() desc:SetEnumItem("EnumA", enumitemA.Name, 42) end, "SetEnumItem(2) with non-string")
T.Pass(function() desc:SetEnumItem("EnumA", enumitemA.Name, enumitemA) end, "SetEnumItem(2) with table")
T.Pass(function() return desc:EnumItem("EnumA","") == nil end, "EnumItem: non-extant enumitem returns nil")
T.Pass(function() return type(desc:EnumItem("EnumA",enumitemA.Name)) == "table" end, "EnumItem: extant enumitem returns table")
T.Pass(function() return desc:EnumItem("EnumA",enumitemA.Name).Name == enumitemA.Name end, "EnumItem: enumitem name matches argument")
T.Pass(function() desc:SetEnumItem("EnumA","Foo", {EnumItemType="Property",Name="Bar"}) end, "set enumitem Foo with name Bar")
T.Pass(function() return desc:EnumItem("EnumA","Foo").Name == "Foo" end, "name argument overrides Name field")
T.Pass(function() desc:SetEnumItem("EnumA","Foo", nil) end, "set enumitem Foo to nil")
T.Pass(function() return desc:EnumItem("EnumA","Foo") == nil end, "enumitem Foo is removed")
desc:SetEnumItem("EnumA",enumitemB.Name, enumitemB)
desc:SetEnumItem("EnumA",enumitemC.Name, enumitemC)
desc:SetEnumItem("EnumA",enumitemD.Name, enumitemD)
T.Equal("enumitem with tags", desc:EnumItem("EnumA","EnumItemC"), enumitemC)
T.Equal("EnumItems returns enumitem names", desc:EnumItems("EnumA"), {"EnumItemA","EnumItemB","EnumItemC","EnumItemD"})
T.Fail(function() desc:EnumItemTag("EnumA") end, "EnumItemTag with no values")
T.Fail(function() desc:EnumItemTag("EnumA","") end, "EnumItemTag with one string")
T.Pass(function() desc:EnumItemTag("EnumA","","") end, "EnumItemTag with two strings")
T.Fail(function() desc:EnumItemTag("EnumA",42,"") end, "EnumItemTag(1) with non-string")
T.Fail(function() desc:EnumItemTag("EnumA","",42) end, "EnumItemTag(2) with non-string")
T.Pass(function() return desc:EnumItemTag("EnumA","Nonextant","Foo") == nil end, "nonextant enumitem tag returns nil")
T.Pass(function() return desc:EnumItemTag("EnumA","EnumItemA","Foo") == false end, "unset enumitem tag returns false")
T.Pass(function() return desc:EnumItemTag("EnumA","EnumItemC","Foo") == true end, "set enumitem tag returns true")

-- EnumTypes
T.Pass(typeof(desc:EnumTypes()) == "Enums", "EnumTypes method returns Enums")
local Enum = desc:EnumTypes()
T.Pass(#Enum:GetEnums() == 4 and
	Enum.EnumA and Enum.EnumB and Enum.EnumC and Enum.EnumD, "EnumTypes reflects defined enums")
T.Pass(Enum ~= desc:EnumTypes(), "EnumTypes regenerates enums.")
desc:SetEnum("EnumA", nil)
desc:SetEnum("EnumB", nil)
desc:SetEnum("EnumC", nil)
desc:SetEnum("EnumD", nil)
local Enum = desc:EnumTypes()
T.Pass(#Enum:GetEnums() == 0, "EnumTypes reflects no defined enums.")

-- Diff
local prev = RootDesc.new()
local next = RootDesc.new()
T.Fail(function() return prev:Diff(42) end        , "Diff with non-desc")
T.Pass(function() return prev:Diff(next) end      , "Diff with desc")
T.Pass(function() return prev:Diff(nil) end       , "Diff with nil")
T.Pass(function() return #prev:Diff(next) == 0 end, "Diff with no differences returns an empty table")
T.Pass(function() return #prev:Diff(nil) == 0 end , "Diff with nil returns an empty table")

T.Fail(function() prev:Patch() end          , "Patch with no values")
T.Fail(function() prev:Patch(42) end        , "Patch with non-patch")
T.Pass(function() prev:Patch({}) end        , "Patch with patch")

-- TODO: verify correctness of returned actions.
