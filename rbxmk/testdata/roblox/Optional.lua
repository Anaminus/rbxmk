local desc = fs.read(path.expand("$sd/../dump.desc.json"))
desc:Patch(fs.read(path.expand("$sd/../dump.desc-patch.json")))

-- No desc
local model = Instance.new("Model")
T.Pass(model.WorldPivotData == nil, "get uninitialized property")
T.Pass(function() model.WorldPivotData = Optional.none("CFrame") end, "set to none")
T.Pass(typeof(model.WorldPivotData) == "Optional", "getting none returns Optional")
T.Pass(model.WorldPivotData.Value == nil, "value of none is nil")
T.Pass(function() model.WorldPivotData = nil end, "set to nil")
T.Pass(function() model.WorldPivotData = Optional.some(CFrame.new(1,2,3)) end, "set to some")
T.Pass(typeof(model.WorldPivotData) == "Optional", "getting some returns Optional")
T.Pass(typeof(model.WorldPivotData.Value) == "CFrame", "value of some is CFrame")
T.Pass(model.WorldPivotData.Value == CFrame.new(1,2,3), "CFrame is expected value")

-- With desc
rbxmk.globalDesc = desc
local model = Instance.new("Model")
T.Fail(function() return model.WorldPivotData == nil end, "with desc, get uninitialized property")
T.Pass(function() model.WorldPivotData = Optional.none("CFrame") end, "with desc, set to none CFrame")
T.Pass(model.WorldPivotData == nil, "with desc, getting none CFrame returns nil")
rbxmk.globalDesc = nil
T.Pass(typeof(model.WorldPivotData) == "Optional", "check that none CFrame sets none CFrame")
rbxmk.globalDesc = desc

T.Pass(function() model.WorldPivotData = nil end, "with desc, set to nil")
T.Pass(model.WorldPivotData == nil, "with desc, getting nil returns nil")
rbxmk.globalDesc = nil
T.Pass(typeof(model.WorldPivotData) == "Optional", "check that nil sets none CFrame")
rbxmk.globalDesc = desc

T.Fail(function() model.WorldPivotData = Optional.none("string") end, "with desc, set to none string")

T.Pass(function() model.WorldPivotData = Optional.some(CFrame.new(1,2,3)) end, "with desc, set to some CFrame")
T.Pass(typeof(model.WorldPivotData) == "CFrame", "with desc, value of some CFrame is CFrame")
T.Pass(model.WorldPivotData == CFrame.new(1,2,3), "with desc, CFrame is expected value")

rbxmk.globalDesc = nil
T.Pass(typeof(model.WorldPivotData) == "Optional", "check that CFrame sets some CFrame")
T.Pass(model.WorldPivotData.Value == CFrame.new(1,2,3), "check that some CFrame value is correct")
rbxmk.globalDesc = desc

T.Fail(function() model.WorldPivotData = Optional.some("Foobar") end, "with desc, set to some string")
T.Pass(typeof(model.WorldPivotData) == "CFrame", "did not set some string")
