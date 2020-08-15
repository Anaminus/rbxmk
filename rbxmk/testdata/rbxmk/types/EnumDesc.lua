local desc = rbxmk.newDesc("EnumDesc")

-- Metamethod tests
T.Pass("type of descriptor is EnumDesc",
	typeof(desc) == "EnumDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("EnumDesc"))

-- Items
local itemA = rbxmk.newDesc("EnumItemDesc")
itemA.Name = "ItemA"
itemA.Value = 1
local itemB = rbxmk.newDesc("EnumItemDesc")
itemB.Name = "ItemB"
itemB.Value = 2
local itemC = rbxmk.newDesc("EnumItemDesc")
itemC.Name = "ItemC"
itemC.Value = 3
local itemD = rbxmk.newDesc("EnumItemDesc")
itemD.Name = "ItemD"
itemD.Value = 4

T.Pass("can call Item method with string",
	function() desc:Item("") end)
T.Fail("cannot call Item method with non-string",
	function() desc:Item(42) end)
T.Pass("can call AddItem method with EnumItemDesc",
	desc:AddItem(itemA) == true)
T.Pass("can call AddItem method with second EnumItemDesc",
	desc:AddItem(itemB) == true)
T.Pass("can call AddItem method with third EnumItemDesc",
	desc:AddItem(itemC) == true)
T.Pass("can call AddItem method with fourth EnumItemDesc",
	desc:AddItem(itemD) == true)
T.Pass("AddItem returns false for existing item",
	desc:AddItem(itemA) == false)
T.Fail("cannot call AddItem method with non-item descriptor",
	function() desc:AddItem(42) end)
T.Pass("Item returns nil for nonextant item",
	desc:Item("Nonextant") == nil)
T.Pass("Item returns item for existing item",
	desc:Item("ItemA") == itemA)
T.Pass("can call RemoveItem method with string",
	function() desc:RemoveItem("") end)
T.Fail("cannot call RemoveItem method with non-string",
	function() desc:RemoveItem(42) end)
T.Pass("RemoveItem returns true for existing item",
	desc:RemoveItem("ItemA") == true)
T.Pass("RemoveItem returns false for nonextant item",
	desc:RemoveItem("ItemA") == false)
T.Pass("removal of item persists",
	desc:Item("ItemA") == nil)
T.Pass("Items method returns a table",
	type(desc:Items()) == "table")
T.Pass("has three items",
	#desc:Items() == 3)
T.Pass("first item is ItemB",
	desc:Items()[1] == itemB)
T.Pass("second item is ItemC",
	desc:Items()[2] == itemC)
T.Pass("third item is ItemD",
	desc:Items()[3] == itemD)
