local desc = fs.read(os.expand("$sd/../../dump.desc.json"))

local b = Instance.new("BoolValue")
local v = Instance.new("ObjectValue")
T.Pass(function() v.Value = b end, "set Instance")
T.Pass(function() v.Value = nil end, "set nil")

local b = Instance.new("BoolValue", nil, desc)
local v = Instance.new("ObjectValue", nil, desc)
T.Pass(function() v.Value = b end, "with desc, set Instance")
T.Pass(function() v.Value = nil end, "with desc, set nil")

local b = Instance.new("BoolValue")
local v = Instance.new("ObjectValue")
v.Value = nil
T.Pass(function() return v.Value == nil end, "get nil")
v.Value = b
T.Pass(function() return v.Value == b end, "get value")
v.Value = nil
T.Pass(function() return v.Value == nil end, "get nil again")

local b = Instance.new("BoolValue", nil, desc)
local v = Instance.new("ObjectValue", nil, desc)
v.Value = nil
T.Pass(function() return v.Value == nil end, "with desc, get nil")
v.Value = b
T.Pass(function() return v.Value == b end, "with desc, get value")
v.Value = nil
T.Pass(function() return v.Value == nil end, "with desc, get nil again")
