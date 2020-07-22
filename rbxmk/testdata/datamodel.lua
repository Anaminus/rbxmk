local game = DataModel.new()
PASS(function() game.ExplicitAutoJoints = true         end, "DataModel must set properties")
PASS(function() return game.ExplicitAutoJoints == true end, "DataModel must get properties")
FAIL(function() game.ClassName = "BoolValue"           end, "cannot set ClassName of DataModel")

local workspace = game:GetService("Workspace")
PASS(function() return workspace.ClassName == "Workspace"        end, "GetService must set ClassName to given value")
PASS(function() return workspace.Name == "Workspace"             end, "GetService must set Name to given value")
PASS(function() return workspace.Parent == game                  end, "GetService must set Parent to DataModel")
PASS(function() return game:GetService("Workspace") == workspace end, "GetService must return singleton")

FAIL(function() Instance.new("BoolValue"):GetService("Workspace") end, "non-DataModel instance must not have GetService")
