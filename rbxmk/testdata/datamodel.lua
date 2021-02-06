-- Test DataModel type.
local game = DataModel.new()
T.Pass(function() game[sym.Metadata] = {ExplicitAutoJoints = "true"} end                          , "DataModel can set metadata")
T.Pass(game[sym.Metadata].ExplicitAutoJoints == "true"                                            , "DataModel can get metadata")
T.Pass(function() game[sym.Metadata] = {}; return game[sym.Metadata].ExplicitAutoJoints == nil end, "DataModel can unset metadata")
T.Fail(function() game[sym.Metadata] = {ExplicitAutoJoints = true} end                            , "DataModel metadata value must be string")

T.Fail(function() game.ClassName = "BoolValue" end, "cannot set ClassName of DataModel")

local workspace = game:GetService("Workspace")
T.Pass(workspace.ClassName == "Workspace"       , "GetService sets ClassName to given value")
T.Pass(workspace.Name == "Workspace"            , "GetService sets Name to given value")
T.Pass(workspace.Parent == game                 , "GetService sets Parent to DataModel")
T.Pass(game:GetService("Workspace") == workspace, "GetService returns singleton")

T.Fail(function() Instance.new("BoolValue"):GetService("Workspace") end, "non-DataModel instance does not have GetService")
