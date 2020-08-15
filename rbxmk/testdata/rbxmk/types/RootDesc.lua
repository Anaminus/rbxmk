local desc = rbxmk.newDesc("RootDesc")

-- Metamethod tests
T.Pass("type of descriptor is RootDesc",
	typeof(desc) == "RootDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("RootDesc"))

-- Classes
local classA = rbxmk.newDesc("ClassDesc")
classA.Name = "ClassA"
local classB = rbxmk.newDesc("ClassDesc")
classB.Name = "ClassB"
local classC = rbxmk.newDesc("ClassDesc")
classC.Name = "ClassC"
local classD = rbxmk.newDesc("ClassDesc")
classD.Name = "ClassD"

T.Pass("can call Class method with string",
	function() desc:Class("") end)
T.Fail("cannot call Class method with non-string",
	function() desc:Class(42) end)
T.Pass("can call AddClass method with ClassDesc",
	desc:AddClass(classA) == true)
T.Pass("can call AddClass method with second ClassDesc",
	desc:AddClass(classB) == true)
T.Pass("can call AddClass method with third ClassDesc",
	desc:AddClass(classC) == true)
T.Pass("can call AddClass method with fourth ClassDesc",
	desc:AddClass(classD) == true)
T.Pass("AddClass returns false for existing class",
	desc:AddClass(classA) == false)
T.Fail("cannot call AddClass method with non-class descriptor",
	function() desc:AddClass(42) end)
T.Pass("Class returns nil for nonextant class",
	desc:Class("Nonextant") == nil)
T.Pass("Class returns class for existing class",
	desc:Class("ClassA") == classA)
T.Pass("can call RemoveClass method with string",
	function() desc:RemoveClass("") end)
T.Fail("cannot call RemoveClass method with non-string",
	function() desc:RemoveClass(42) end)
T.Pass("RemoveClass returns true for existing class",
	desc:RemoveClass("ClassA") == true)
T.Pass("RemoveClass returns false for nonextant class",
	desc:RemoveClass("ClassA") == false)
T.Pass("removal of class persists",
	desc:Class("ClassA") == nil)
T.Pass("Classes method returns a table",
	type(desc:Classes()) == "table")
T.Pass("has three classes",
	#desc:Classes() == 3)
T.Pass("first class is Callback",
	desc:Classes()[1] == classB)
T.Pass("second class is Event",
	desc:Classes()[2] == classC)
T.Pass("third class is Function",
	desc:Classes()[3] == classD)

-- Enums
local enumA = rbxmk.newDesc("EnumDesc")
enumA.Name = "EnumA"
local enumB = rbxmk.newDesc("EnumDesc")
enumB.Name = "EnumB"
local enumC = rbxmk.newDesc("EnumDesc")
enumC.Name = "EnumC"
local enumD = rbxmk.newDesc("EnumDesc")
enumD.Name = "EnumD"

T.Pass("can call Enum method with string",
	function() desc:Enum("") end)
T.Fail("cannot call Enum method with non-string",
	function() desc:Enum(42) end)
T.Pass("can call AddEnum method with EnumDesc",
	desc:AddEnum(enumA) == true)
T.Pass("can call AddEnum method with second EnumDesc",
	desc:AddEnum(enumB) == true)
T.Pass("can call AddEnum method with third EnumDesc",
	desc:AddEnum(enumC) == true)
T.Pass("can call AddEnum method with fourth EnumDesc",
	desc:AddEnum(enumD) == true)
T.Pass("AddEnum returns false for existing enum",
	desc:AddEnum(enumA) == false)
T.Fail("cannot call AddEnum method with non-enum descriptor",
	function() desc:AddEnum(42) end)
T.Pass("Enum returns nil for nonextant enum",
	desc:Enum("Nonextant") == nil)
T.Pass("Enum returns enum for existing enum",
	desc:Enum("EnumA") == enumA)
T.Pass("can call RemoveEnum method with string",
	function() desc:RemoveEnum("") end)
T.Fail("cannot call RemoveEnum method with non-string",
	function() desc:RemoveEnum(42) end)
T.Pass("RemoveEnum returns true for existing enum",
	desc:RemoveEnum("EnumA") == true)
T.Pass("RemoveEnum returns false for nonextant enum",
	desc:RemoveEnum("EnumA") == false)
T.Pass("removal of enum persists",
	desc:Enum("EnumA") == nil)
T.Pass("Enums method returns a table",
	type(desc:Enums()) == "table")
T.Pass("has three enums",
	#desc:Enums() == 3)
T.Pass("first enum is Callback",
	desc:Enums()[1] == enumB)
T.Pass("second enum is Event",
	desc:Enums()[2] == enumC)
T.Pass("third enum is Function",
	desc:Enums()[3] == enumD)

-- EnumTypes
T.Pass("EnumTypes method returns Enums",
	typeof(desc:EnumTypes()) == "Enums")
local Enum = desc:EnumTypes()
T.Pass("EnumTypes reflects defined enums",
	#Enum:GetEnums() == 3 and
	Enum.EnumB and Enum.EnumC and Enum.EnumD)
T.Pass("EnumTypes regenerates enums",
	Enum ~= desc:EnumTypes())
desc:RemoveEnum("EnumB")
desc:RemoveEnum("EnumC")
desc:RemoveEnum("EnumD")
local Enum = desc:EnumTypes()
T.Pass("EnumTypes reflects no defined enums",
	#Enum:GetEnums() == 0)
