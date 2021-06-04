local desc = fs.read(path.expand("$sd/../dump.desc.json"))

-- As property
local pp = PhysicalProperties.new(1,2,3,4,5)

local part = Instance.new("Part")
T.Pass(function() part.CustomPhysicalProperties = pp end, "set value")
T.Pass(function() part.CustomPhysicalProperties = nil end, "set to nil")

local part = Instance.new("Part", nil, desc)
T.Pass(function() part.CustomPhysicalProperties = pp end, "with desc, set value")
T.Pass(function() part.CustomPhysicalProperties = nil end, "with desc, set to nil")

local part = Instance.new("Part")
T.Pass(function() return part.CustomPhysicalProperties == nil end, "get nil")
part.CustomPhysicalProperties = pp
T.Pass(function() return part.CustomPhysicalProperties == pp end, "get value")
part.CustomPhysicalProperties = nil
T.Pass(function() return part.CustomPhysicalProperties == nil end, "get nil again")

local part = Instance.new("Part", nil, desc)
part.CustomPhysicalProperties = nil
T.Pass(function() return part.CustomPhysicalProperties == nil end, "with desc, get nil")
part.CustomPhysicalProperties = pp
T.Pass(function() return part.CustomPhysicalProperties == pp end, "with desc, get value")
part.CustomPhysicalProperties = nil
T.Pass(function() return part.CustomPhysicalProperties == nil end, "with desc, get nil again")
