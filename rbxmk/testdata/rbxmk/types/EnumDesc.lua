local desc = rbxmk.newDesc("EnumDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "EnumDesc"                   , "type of descriptor is EnumDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("EnumDesc")            , "descriptor is not equal to another descriptor of the same type")

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

T.Pass(function() desc:Item("") end      , "can call Item method with string")
T.Fail(function() desc:Item(42) end      , "cannot call Item method with non-string")
T.Pass(desc:AddItem(itemA) == true       , "can call AddItem method with EnumItemDesc")
T.Pass(desc:AddItem(itemB) == true       , "can call AddItem method with second EnumItemDesc")
T.Pass(desc:AddItem(itemC) == true       , "can call AddItem method with third EnumItemDesc")
T.Pass(desc:AddItem(itemD) == true       , "can call AddItem method with fourth EnumItemDesc")
T.Pass(desc:AddItem(itemA) == false      , "AddItem returns false for existing item")
T.Fail(function() desc:AddItem(42) end   , "cannot call AddItem method with non-item descriptor")
T.Pass(desc:Item("Nonextant") == nil     , "Item returns nil for nonextant item")
T.Pass(desc:Item("ItemA") == itemA       , "Item returns item for existing item")
T.Pass(function() desc:RemoveItem("") end, "can call RemoveItem method with string")
T.Fail(function() desc:RemoveItem(42) end, "cannot call RemoveItem method with non-string")
T.Pass(desc:RemoveItem("ItemA") == true  , "RemoveItem returns true for existing item")
T.Pass(desc:RemoveItem("ItemA") == false , "RemoveItem returns false for nonextant item")
T.Pass(desc:Item("ItemA") == nil         , "Removal of item persists")
T.Pass(type(desc:Items()) == "table"     , "Items method returns a table")
T.Pass(#desc:Items() == 3                , "has three items")
T.Pass(desc:Items()[1] == itemB          , "first item is ItemB")
T.Pass(desc:Items()[2] == itemC          , "second item is ItemC")
T.Pass(desc:Items()[3] == itemD          , "third item is ItemD")
