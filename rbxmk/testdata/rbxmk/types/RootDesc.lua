local desc = rbxmk.newDesc("RootDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "RootDesc"                   , "type of descriptor is RootDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("RootDesc")            , "descriptor is not equal to another descriptor of the same type")

-- Classes
local classA = rbxmk.newDesc("ClassDesc")
classA.Name = "ClassA"
local classB = rbxmk.newDesc("ClassDesc")
classB.Name = "ClassB"
local classC = rbxmk.newDesc("ClassDesc")
classC.Name = "ClassC"
local classD = rbxmk.newDesc("ClassDesc")
classD.Name = "ClassD"

T.Pass(function() desc:Class("") end      , "can call Class method with string")
T.Fail(function() desc:Class(42) end      , "cannot call Class method with non-string")
T.Pass(desc:AddClass(classA) == true      , "can call AddClass method with ClassDesc")
T.Pass(desc:AddClass(classB) == true      , "can call AddClass method with second ClassDesc")
T.Pass(desc:AddClass(classC) == true      , "can call AddClass method with third ClassDesc")
T.Pass(desc:AddClass(classD) == true      , "can call AddClass method with fourth ClassDesc")
T.Pass(desc:AddClass(classA) == false     , "AddClass returns false for existing class")
T.Fail(function() desc:AddClass(42) end   , "cannot call AddClass method with non-class descriptor")
T.Pass(desc:Class("Nonextant") == nil     , "class returns nil for nonextant class")
T.Pass(desc:Class("ClassA") == classA     , "class returns class for existing class")
T.Pass(function() desc:RemoveClass("") end, "can call RemoveClass method with string")
T.Fail(function() desc:RemoveClass(42) end, "cannot call RemoveClass method with non-string")
T.Pass(desc:RemoveClass("ClassA") == true , "RemoveClass returns true for existing class")
T.Pass(desc:RemoveClass("ClassA") == false, "RemoveClass returns false for nonextant class")
T.Pass(desc:Class("ClassA") == nil        , "removal of class persists")
T.Pass(type(desc:Classes()) == "table"    , "Classes method returns a table")
T.Pass(#desc:Classes() == 3               , "has three classes")
T.Pass(desc:Classes()[1] == classB        , "first class is Callback")
T.Pass(desc:Classes()[2] == classC        , "second class is Event")
T.Pass(desc:Classes()[3] == classD        , "third class is Function")

-- Enums
local enumA = rbxmk.newDesc("EnumDesc")
enumA.Name = "EnumA"
local enumB = rbxmk.newDesc("EnumDesc")
enumB.Name = "EnumB"
local enumC = rbxmk.newDesc("EnumDesc")
enumC.Name = "EnumC"
local enumD = rbxmk.newDesc("EnumDesc")
enumD.Name = "EnumD"

T.Pass(function() desc:Enum("") end      , "can call Enum method with string")
T.Fail(function() desc:Enum(42) end      , "cannot call Enum method with non-string")
T.Pass(desc:AddEnum(enumA) == true       , "can call AddEnum method with EnumDesc")
T.Pass(desc:AddEnum(enumB) == true       , "can call AddEnum method with second EnumDesc")
T.Pass(desc:AddEnum(enumC) == true       , "can call AddEnum method with third EnumDesc")
T.Pass(desc:AddEnum(enumD) == true       , "can call AddEnum method with fourth EnumDesc")
T.Pass(desc:AddEnum(enumA) == false      , "AddEnum returns false for existing enum")
T.Fail(function() desc:AddEnum(42) end   , "cannot call AddEnum method with non-enum descriptor")
T.Pass(desc:Enum("Nonextant") == nil     , "enum returns nil for nonextant enum")
T.Pass(desc:Enum("EnumA") == enumA       , "enum returns enum for existing enum")
T.Pass(function() desc:RemoveEnum("") end, "can call RemoveEnum method with string")
T.Fail(function() desc:RemoveEnum(42) end, "cannot call RemoveEnum method with non-string")
T.Pass(desc:RemoveEnum("EnumA") == true  , "RemoveEnum returns true for existing enum")
T.Pass(desc:RemoveEnum("EnumA") == false , "RemoveEnum returns false for nonextant enum")
T.Pass(desc:Enum("EnumA") == nil         , "removal of enum persists")
T.Pass(type(desc:Enums()) == "table"     , "Enums method returns a table")
T.Pass(#desc:Enums() == 3                , "has three enums")
T.Pass(desc:Enums()[1] == enumB          , "first enum is Callback")
T.Pass(desc:Enums()[2] == enumC          , "second enum is Event")
T.Pass(desc:Enums()[3] == enumD          , "third enum is Function")

-- EnumTypes
T.Pass(typeof(desc:EnumTypes()) == "Enums", "EnumTypes method returns Enums")
local Enum = desc:EnumTypes()
T.Pass(#Enum:GetEnums() == 3 and
	Enum.EnumB and Enum.EnumC and Enum.EnumD, "EnumTypes reflects defined enums")
T.Pass(Enum ~= desc:EnumTypes(), "EnumTypes regenerates enums.")
desc:RemoveEnum("EnumB")
desc:RemoveEnum("EnumC")
desc:RemoveEnum("EnumD")
local Enum = desc:EnumTypes()
T.Pass(#Enum:GetEnums() == 0, "EnumTypes reflects no defined enums.")
