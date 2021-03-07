local desc0 = fs.read(os.expand("$sd/enums.desc.json"))
local desc1 = fs.read(os.expand("$sd/enums.desc.json"))

local Enums0 = desc0:EnumTypes()
local Enums1 = desc1:EnumTypes()
local Enum0 = Enums0.NormalId
local Enum1 = Enums1.NormalId
local EnumItem0 = Enums0.NormalId.Front
local EnumItem1 = Enums1.NormalId.Front

-- Assignment
T.Fail(function() Enums0.NormalId = Enums1.NormalId end, "assign to Enums")
T.Fail(function() Enum0.Front = Enum1.Front end, "assign to Enum")
T.Fail(function() EnumItem0.Value = 42 end, "assign to EnumItem")

-- type
T.Pass(type(Enums0) == "userdata", "Enums is userdata")
T.Pass(type(Enum0) == "userdata", "Enum is userdata")
T.Pass(type(EnumItem0) == "userdata", "EnumItem is userdata")

-- __type
T.Pass(typeof(Enums0) == "Enums", "typeof Enums")
T.Pass(typeof(Enum0) == "Enum", "typeof Enum")
T.Pass(typeof(EnumItem0) == "EnumItem", "typeof EnumItem")

-- __eq
T.Pass(Enums0 == Enums0, "Enums 0 == 0")
T.Pass(Enums1 ~= Enums0, "Enums 1 ~= 0")
T.Pass(Enums0 ~= Enums1, "Enums 0 ~= 1")
T.Pass(Enums1 == Enums1, "Enums 1 == 1")

T.Pass(Enum0 == Enum0, "Enum 0 == 0")
T.Pass(Enum1 ~= Enum0, "Enum 1 ~= 0")
T.Pass(Enum0 ~= Enum1, "Enum 0 ~= 1")
T.Pass(Enum1 == Enum1, "Enum 1 == 1")

T.Pass(EnumItem0 == EnumItem0, "EnumItem 0 == 0")
T.Pass(EnumItem1 ~= EnumItem0, "EnumItem 1 ~= 0")
T.Pass(EnumItem0 ~= EnumItem1, "EnumItem 0 ~= 1")
T.Pass(EnumItem1 == EnumItem1, "EnumItem 1 == 1")

-- __tostring
T.Pass(tostring(Enums0) == "Enums", "tostring Enums")
T.Pass(tostring(Enum0) == "NormalId", "tostring Enum")
T.Pass(tostring(EnumItem0) == "Enum.NormalId.Front", "tostring EnumItem")

-- GetEnums
local enums = Enums0:GetEnums()
T.Pass(type(enums) == "table", "GetEnums returns table")
T.Pass(#enums == 14, "#GetEnums == 14")
-- Enums are sorted.
T.Pass(tostring(enums[13]) == "ZZZY", "GetEnums[13] == ZZZY")
T.Pass(tostring(enums[14]) == "ZZZZ", "GetEnums[14] == ZZZZ")

-- GetEnumItems
local items = Enums0.ZZZZ:GetEnumItems()
T.Pass(type(items) == "table", "GetEnumItems returns table")
T.Pass(#items == 4, "#GetEnumItems == 4")
-- items are not sorted.
T.Pass(items[1].Name == "DDDD" and items[1].Value == 4, "GetEnumItems[1] == DDDD")
T.Pass(items[2].Name == "CCCC" and items[2].Value == 3, "GetEnumItems[2] == CCCC")
T.Pass(items[3].Name == "BBBB" and items[3].Value == 2, "GetEnumItems[3] == BBBB")
T.Pass(items[4].Name == "AAAA" and items[4].Value == 1, "GetEnumItems[4] == AAAA")

