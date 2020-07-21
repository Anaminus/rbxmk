local game = DataModel.new()
assert(pcall(function() game.ExplicitAutoJoints = true end), "DataModel must set properties")
assert(game.ExplicitAutoJoints == true, "DataModel must get properties")
assert(not pcall(function() game.ClassName = "BoolValue" end), "cannot set ClassName of DataModel")

local workspace = game:GetService("Workspace")
assert(workspace.ClassName == "Workspace", "GetService must set ClassName to given value")
assert(workspace.Name == "Workspace", "GetService must set Name to given value")
assert(workspace.Parent == game, "GetService must set Parent to DataModel")
assert(game:GetService("Workspace") == workspace, "GetService must return singleton")

assert(not pcall(function() Instance.new("BoolValue"):GetService("Workspace") end), "non-DataModel instance must not have GetService")
